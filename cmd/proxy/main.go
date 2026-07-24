package main

import (
	"log"

	"github.com/mocheer/droid/gadb"
)

func main() {
	// adb shell settings put global http_proxy [代理IP]:[端口]
	// 测试5037端口是否正常，并实例化客户端对象
	adbClient, err := gadb.NewClient()
	if err != nil {
		log.Println(err)
	}
	//
	devices, err := adbClient.DeviceList()
	if err != nil {
		log.Println(err)
	}
	if len(devices) == 0 {
		log.Println("没有连接，开始连接127.0.0.1:16384")
		err := adbClient.Connect("127.0.0.1", 16384)
		if err != nil {
			panic(err)
		}
		// 重新获取设备信息
		devices, err = adbClient.DeviceList()
		if err != nil {
			log.Println(err)
		}
	}
	// devices[0]
	// 这里设备可能不止一台，需要选择想要操作的设备
	d := devices[0]
	// gadb.SetDebug(true)
	// 向上滑动
	d.ScreencapSave("")

}

// // 修改请求的函数
// func modifyRequest(r *http.Request) {
// 	// 示例：修改请求的 URL
// 	// if strings.Contains(r.URL.String(), "example.com") {
// 	// 	r.URL, _ = url.Parse("https://newexample.com/newpath")
// 	// }

// 	// // 示例：修改请求头
// 	// r.Header.Set("User-Agent", "Custom-User-Agent")

// 	// 示例：修改请求体（如果需要）
// 	// 注意：修改请求体需要读取并重新设置 r.Body
// }

// // 反向代理处理函数
// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	// 打印原始请求
// 	log.Printf("Received request: %s %s", r.Method, r.URL.String())

// 	// 修改请求
// 	modifyRequest(r)

// 	// 创建一个反向代理
// 	proxy := httputil.NewSingleHostReverseProxy(r.URL)

// 	// 转发请求
// 	proxy.ServeHTTP(w, r)
// }

// func main() {
// 	// 启动代理服务器
// 	http.HandleFunc("/", handleRequest)
// 	log.Println("Starting proxy server on :8888")
// 	if err := http.ListenAndServe(":8888", nil); err != nil {
// 		log.Fatal("Error starting proxy server: ", err)
// 	}
// }
