package gadb

// Home
// 相当于点击Home按键
func (d Device) Home() {
	d.RunShellCommand("input", "keyevent", "3")
}

// 返回键
func (d Device) Back() {
	d.RunShellCommand("input", "keyevent", "4")
}
