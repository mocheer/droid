package gadb

// 应用权限管理

// 运行后台运行
// adb shell appops set <包名> RUN_IN_BACKGROUND allow
func (d Device) AllowRunInBackground(packageName string) (string, error) {
	return d.RunShellCommand("adb shell appops set", packageName, "RUN_IN_BACKGROUND allow")
}

func (d Device) IsAllowRunInBackground(packageName string) (bool, error) {
	result, err := d.RunShellCommand("adb shell appops get", packageName, "RUN_IN_BACKGROUND")
	// 返回 allow 表示允许后台运行；deny 或 ignore 表示被限制 。
	return result == "RUN_IN_BACKGROUND: allow", err
}
