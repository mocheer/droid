package gadb

import (
	"strconv"
	"time"
)

// Screenrecord 录频
// adb shell screenrecord path
// adb shell screenrecord --size 1280*720 path
// adb shell screenrecord --time-limit 130 path //指定时长为130秒
// 一般要带上根目录：sdcard
// d.Screenrecord("/sdcard/a.mp4")
func (d Device) Screenrecord(filename string) (string, error) {
	return d.RunShellCommand("screenrecord", filename)
}

func (d Device) ScreenrecordWithTimeLimit(filename string, timeLimit time.Duration) (string, error) {
	return d.RunShellCommand("screenrecord --time-limit", strconv.Itoa(int(timeLimit)), filename)
}

func (d Device) ScreenrecordWithSize(filename string, size string) (string, error) {
	return d.RunShellCommand("screenrecord --size", size, filename)
}
