package gadb

// https://developer.android.google.cn/tools/adb?hl=zh-cn#am
// https://developer.android.google.cn/tools/adb?hl=zh-cn#pm

// UninstallWithSystem
// 卸载系统应用
// adb shell pm uninstall --user 0 com.mumu.store
func (d Device) UninstallWithSystem(packageName string) (string, error) {
	return d.RunShellCommand("pm uninstall", packageName)
}

// 启动应用
func (d Device) StartWithActivity(packageName string, activityName string) (string, error) {
	return d.RunShellCommand("am start", packageName+"/"+activityName)
}

// TODO 检查应用是否启动成功
func (d Device) IsAppStarted(packageName string) (bool, error) {
	return true, nil
}

// 启动应用
// -e 传递参数
// -S 启动前先强制停止目标应用（确保冷启动）。
func (d Device) StartWithArgs(packageName string, activityName string, args string) (string, error) {
	return d.RunShellCommand("am start -S -n", packageName+"/"+activityName, args)
}

// 强制停止应用关联的所有进程。
// adb shell am force-stop <包名>
// adb shell am force-stop com.android.settings
func (d Device) Stop(packageName string) (string, error) {
	return d.RunShellCommand("am force-stop", packageName)
}

// 清除指定应用的所有用户数据（包括数据库、SharedPreferences、缓存文件等），使应用恢复到首次安装时的状态
// 不会主动终止进程，但清除数据可能导致应用因状态丢失而自行崩溃或退出
// adb shell pm clear <包名>
func (d Device) Clear(packageName string) (string, error) {
	return d.RunShellCommand("pm clear  ", packageName)
}

func (d Device) StopAndClear(packageName string) (string, error) {
	d.Stop(packageName)
	return d.Clear(packageName)
}

func (d Device) GetRunAppInfo(packageName string) (string, error) {
	result, err := d.RunShellCommand("dumpsys activity activities | grep", packageName)
	return result, err
}

// PackageInfo
// 查看包信息
// adb shell dumpsys package package_name
// 可以查看包的版本、Activity（用于启动）等信息
// android:name属性对应的值，即为启动Activity的完整类名。
func (d Device) PackageInfo(packageName string) (string, error) {
	return d.RunShellCommand("dumpsys package", packageName)
}
