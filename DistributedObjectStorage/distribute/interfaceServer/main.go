package main

import (
	"distribute/interfaceServer/heartbeat"
	"distribute/interfaceServer/locate"
	"distribute/interfaceServer/objects"
	"log"
	"net/http"
	"os"
)

/*
@Time    : 2021/8/8 15:52
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

// 接口层

func main() {
	// 存储检查更新自动注册的心跳
	go heartbeat.ListenHeartbeat()

	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
