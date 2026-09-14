package gadb

import (
	"path"
	"path/filepath"
)

// 支持修改默认的共享目录
var SharedMuMu = `C:\Users\Administrator\Documents\MuMu共享文件夹`

// screencap:
// usage: screencap [-hp] [-d display-id] [FILENAME]
//    -h: this message
//    -p: save the file as a png.
//    -d: specify the physical display ID to capture (default: 4619827820427265280)
//        see "dumpsys SurfaceFlinger --display-id" for valid display IDs.
// If FILENAME ends with .png it will be saved as a png.
// If FILENAME is not given, the results will be printed to stdout.

// Screencap
// 不指定文件时会直接返回二进制数据
// 不指定-p的情况下，生成的图片数据为原始帧缓冲区格式（Raw Framebuffer Dump），这种数据是未经压缩的二进制流，包含屏幕每个像素的 RGBA 值，无文件头或压缩结构
// 未压缩的数据量比png大很多
func (d Device) Screencap() ([]byte, error) {
	return d.RunShellCommandWithBytes("screencap -p")
}

// ScreencapSave
// adb screencap -p --crop=w:h:x:y
// 一般要带上根目录：sdcard
// d.ScreencapSave("/sdcard/a.png")
func (d Device) ScreencapSave(filename string) (string, error) {
	return d.RunShellCommand("screencap", filename)
}

// ScreencapSaveSD
func (d Device) ScreencapSaveSD(filename string) (string, error) {
	return d.RunShellCommand("screencap", path.Join("/sdcard", filename))
}

// ScreencapSaveMuMuShared
// 直接截图到共享目录下的Screenshots目录，方便读取
func (d Device) ScreencapSaveMuMuShared(filename string) (string, error) {
	return d.RunShellCommand("screencap", path.Join("/sdcard/$MuMu12Shared/Screenshots", filename))
}

// GetWindoweMuMuSharedPath
func (d Device) GetWindoweMuMuSharedScreenshotsPath(filename string) string {
	return filepath.Join(SharedMuMu, "Screenshots")
}
