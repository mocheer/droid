package gadb

// Install
// adb -s emulator-5555 install path_to_apk
func (d Device) Install(filename string) {
	d.adbClient.executeCommand("-s " + "" + d.serial + "install " + filename)
}
