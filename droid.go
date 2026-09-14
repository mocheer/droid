package droid

import (
	"log"
	"slices"

	"github.com/mocheer/droid/pkg/gadb"
)

type Droid struct {
	adb     gadb.Client
	devices []gadb.Device
	vms     []*Vm
}

// New
func New() *Droid {
	return NewWithEmulator()
}

// NewWithEmulator
func NewWithEmulator() *Droid {
	adb, err := gadb.NewClientWith("localhost")
	if err != nil {
		log.Println("连接错误", err)
	}
	return &Droid{adb: adb}
}

// FindDevice
// 当 serial 为空且设备数量大于 0 时，获取第一个设备
func (m *Droid) FindDevice(serial string) (device gadb.Device) {
	//
	if len(m.devices) == 0 {
		log.Println("当前设备为空，正在重新查找设备...")
		devices, err := m.adb.DeviceList()
		if err != nil {
			log.Println(err)
		}
		m.devices = devices
		log.Println("找到设备数量：", len(devices))
	}
	//
	if serial != "" {
		i := slices.IndexFunc(m.devices, func(d gadb.Device) bool {
			return d.Serial() == serial
		})
		if i != -1 {
			device = m.devices[i]
		} else {
			log.Println("没有找到设备：", serial)
		}
	} else {
		device = m.devices[0]
	}
	return device
}

// Run
func (m *Droid) Run(options DroidRunOptions) {
	// 模拟器多开，或者有多个模拟器的情况下，这里设备可能不止一台，需要选择想要操作的设备
	device := m.FindDevice(options.Serial)
	log.Println("目标模拟器：", device.Serial(), device.DeviceInfo())
	vm := &Vm{device: device, options: options}
	vm.Run()
	m.vms = append(m.vms, vm)
}

// Stop
func (m *Droid) Stop() {
	for _, vm := range m.vms {
		vm.Stop()
	}
}

// Has
func (m *Droid) Has(name string) bool {
	for _, vm := range m.vms {
		if vm.options.Name == name {
			return true
		}
	}
	return false
}

// Remove
func (m *Droid) Remove(name string) {
	for _, vm := range m.vms {
		if vm.options.Name == name {
			vm.Stop()
			m.vms = slices.DeleteFunc(m.vms, func(toDeleteVm *Vm) bool {
				return toDeleteVm == vm
			})
			return
		}
	}
}
