package main

import (
	"log"

	"github.com/mocheer/droid/pkg/gadb"
)

func main() {
	// 测试5037端口是否正常，并实例化客户端对象
	adbClient, err := gadb.NewClient()
	if err != nil {
		log.Println(err)
	}
	//
	devices, err := adbClient.DeviceList()
	if err != nil {
		log.Println(err)
	}
	if len(devices) == 0 {
		log.Println("没有连接，开始连接127.0.0.1:16384")
		err := adbClient.Connect("127.0.0.1", 16384)
		if err != nil {
			panic(err)
		}
		// 重新获取设备信息
		devices, err = adbClient.DeviceList()
		if err != nil {
			log.Println(err)
		}
	}
	// devices[0]
	// 这里设备可能不止一台，需要选择想要操作的设备
	d := devices[0]
	// gadb.SetDebug(true)
	// 向上滑动
	d.SwipUpWith1080()

}
