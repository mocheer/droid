package gadb

import (
	"fmt"
	"strconv"
)

// 这里的坐标不要直接用模拟器界面的像素，还要通过模拟器设置的分辨率和DPI才能计算出来
// 480/96=5

// Tap
// adb shell input tap 100 200
func (d Device) Tap(x, y int) (string, error) {
	xArgs := strconv.Itoa(x)
	yArgs := strconv.Itoa(y)
	return d.RunShellCommand("input", "tap", xArgs, yArgs)
}

// Keydown 按下（不松开）
func (d Device) Keydown(x, y int) (string, error) {
	xArgs := fmt.Sprintf("--x=%d", x)
	yArgs := fmt.Sprintf("--y=%d", y)
	return d.RunShellCommand("input", "--action=1", "--pointerId=1", xArgs, yArgs)
}

func (d Device) Move(x, y int) (string, error) {
	xArgs := fmt.Sprintf("--x=%d", x)
	yArgs := fmt.Sprintf("--y=%d", y)
	return d.RunShellCommand("input", "--action=3", "--pointerId=1", xArgs, yArgs)
}

// KeyUp 松开
func (d Device) KeyUp(x, y int) (string, error) {
	xArgs := fmt.Sprintf("--x=%d", x)
	yArgs := fmt.Sprintf("--y=%d", y)
	return d.RunShellCommand("input", "--action=0", "--pointerId=1", xArgs, yArgs)
}

// Swip
// 滑动
// 单位毫秒
func (d Device) Swip(sx, sy, ex, ey int, duration int) (string, error) {
	sxArgs := strconv.Itoa(sx)
	syArgs := strconv.Itoa(sy)
	exArgs := strconv.Itoa(ex)
	eyArgs := strconv.Itoa(ey)
	durationArgs := strconv.Itoa(duration)
	return d.RunShellCommand("input", "swipe", sxArgs, syArgs, exArgs, eyArgs, durationArgs)
}

// 这里理应提前获取device分辨率，根据分辨率来滑动

// SwipUpWith1080
// 向上滑动
func (d Device) SwipUpWith1080() (string, error) {
	return d.Swip(540, 1280, 540, 640, 500)
}

// SwipDownWith1080
// 向下滑动
func (d Device) SwipDownWith1080() (string, error) {
	return d.Swip(540, 640, 540, 1280, 500)
}

// 向左滑动
func (d Device) SwipLeftWith1080() (string, error) {
	return d.Swip(720, 960, 360, 960, 500)
}

// 向右滑动
func (d Device) SwipRightWith1080() (string, error) {
	return d.Swip(360, 960, 720, 960, 500)
}
