package droid

import (
	"log"
	"time"

	"github.com/mocheer/droid/gadb"
	"github.com/mocheer/pluto/pkg/ts/clock"
)

type Vm struct {
	device      gadb.Device
	options     DroidRunOptions
	cancelFuncs []func()
	startFunc   func()
	stopFunc    func()
	isRuning    bool
}

type DroidRunOptions struct {
	Name             string // 名称
	Serial           string // 设备序列号，为空时使用第一个设备
	AppName          string // 应用包名
	ActivityName     string // 应用启动的Activity名称，可为空
	KeepAlive        bool
	RestartInterval  time.Duration // 启动应用，不能太低，启动应用需要时间
	StopInterval     time.Duration //
	SnapshotInterval time.Duration // 间隔截图，可用于分析和检测应用状态
	SnapshotCallback func(data []byte)
}

// NewDroidRunOptions 创建默认的DroidRunOptions
var NewDroidRunOptions = func(name string, appName string, activityName string) DroidRunOptions {
	return DroidRunOptions{
		Name:            name,
		AppName:         appName,
		ActivityName:    activityName,
		RestartInterval: 6 * time.Second,
		StopInterval:    15 * time.Minute,
		KeepAlive:       true,
	}
}

func (m *Vm) Run() {
	t := time.Now().Unix()
	device := m.device
	//
	m.cancelFuncs = []func(){}
	m.startFunc = func() {
		_, err := device.StartWithActivity(m.options.AppName, m.options.ActivityName)
		if err != nil {
			log.Println("启动失败", err)
			return
		}
		m.isRuning = true
		// log.Println("启动信息", result)
	}
	// 应用守护，防止因为崩溃等各种原因导致的应用关闭
	if m.options.KeepAlive {
		if m.options.RestartInterval == 0 {
			m.options.RestartInterval = 5000 * time.Millisecond
		}
		cancel := clock.SetInterval(func() {
			log.Println("启动应用", m.options.AppName)
			m.startFunc()
		}, m.options.RestartInterval, true)
		m.cancelFuncs = append(m.cancelFuncs, cancel)
	}
	// 有些应用为了防止请求阻塞和卡顿问题需要间隔重启
	m.stopFunc = func() {
		device.StopAndClear(m.options.AppName)
		m.isRuning = false
		t2 := time.Now().Unix()
		log.Println("应用已关闭，运行时间为：", m.options.AppName, t2-t)
	}
	if m.options.StopInterval > 0 {
		cancel := clock.SetInterval(func() {
			log.Println("关闭应用", m.options.AppName)
			m.stopFunc()
		}, m.options.StopInterval, false)
		m.cancelFuncs = append(m.cancelFuncs, cancel)
	}

	//
	if m.options.SnapshotInterval > 0 {
		cancel := clock.SetInterval(func() {
			isStarted, err := device.IsAppStarted(m.options.AppName)
			if err != nil {
				log.Println("获取设备是否启动失败", err)
			}
			if isStarted {

				data, err := device.Screencap()
				if err != nil {
					log.Println("截图错误", err)
					return
				}
				// 这个函数可能会做耗时的操作，如果SnapshotInterval间隔很短就很容易阻塞
				go m.options.SnapshotCallback(data)
				// filename := clock.Now().Fmt(clock.FmtCompactFullDate)
				// device.ScreencapSaveMuMuShared(filename + ".jpg")
			}
		}, m.options.SnapshotInterval, false)
		m.cancelFuncs = append(m.cancelFuncs, cancel)
	}
}

func (m *Vm) Stop() {
	for _, cancel := range m.cancelFuncs {
		cancel()
	}
	m.stopFunc() //关闭应用
	m.cancelFuncs = []func(){}
}
