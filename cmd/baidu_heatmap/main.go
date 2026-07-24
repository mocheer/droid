package main

import (
	"log"

	"github.com/mocheer/droid/gadb"
	"github.com/mocheer/pluto/pkg/ts/img"
)

func main() {
	adb, err := gadb.NewClientWith("localhost")
	if err != nil {
		log.Println("错误")
	}
	//
	devices, err := adb.DeviceList()
	if err != nil {
		log.Println(err)
	}
	data, _ := devices[0].Screencap()
	i, _, _ := img.FromBytes(data)
	i.Save("a.png", "png")
	// ds.Save("rgba", data)
}
