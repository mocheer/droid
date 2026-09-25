package droid_test

import (
	"testing"
	"time"

	"github.com/mocheer/droid/cmd/droid"
)

func Test(t *testing.T) {
	d := droid.New("localhost")
	d.RunWithOptions(droid.DroidConfig{
		Name:         "baidu_heatmap",
		AppName:      "com.demo.istrongheapmap",
		ActivityName: ".ListenActivity",
		StopInterval: 1000 * 60 * 20 * time.Millisecond,
		Alive:        true,
	})
	select {}
}
