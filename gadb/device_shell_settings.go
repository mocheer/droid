package gadb

// SetProxy
// 此方法会为设备所有网络连接（包括 Wi-Fi 和移动数据）设置代理
// # 设置全局代理（IP 和端口需替换为实际值）
// adb shell settings put global http_proxy <代理IP>:<端口>

// # 示例：设置代理为 192.168.1.100:8080
// adb shell settings put global http_proxy 192.168.1.100:8080
func (d Device) SetProxy(addr string) (string, error) {
	return d.RunShellCommand("adb shell settings put global http_proxy", addr)
}

// 清除全局代理
// adb shell settings put global http_proxy :0
func (d Device) ResetProxy() (string, error) {
	return d.RunShellCommand("adb shell settings put global http_proxy", ":0")
}

// 查询代理设置
func (d Device) GetProxy() (string, error) {
	return d.RunShellCommand("adb shell settings get global http_proxy")
}
