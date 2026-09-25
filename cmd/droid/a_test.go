package droid_test

import (
	"testing"
	"time"

	"github.com/mocheer/droid/cmd/droid"
)

func Test(t *testing.T) {
	d := droid.New()
	d.RunWithOptions(droid.DroidRunOptions{
		Name:         "baidu_heatmap",
		AppName:      "com.demo.istrongheapmap",
		ActivityName: ".ListenActivity",
		StopInterval: 1000 * 60 * 20 * time.Millisecond,
		KeepAlive:    true,
	})
	select {}
}
