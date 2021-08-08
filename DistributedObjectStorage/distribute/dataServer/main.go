package main

import (
	"distribute/dataServer/handler"
	"distribute/dataServer/heartbeat"
	"distribute/dataServer/locate"
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

// 数据层

func main() {
	// 将本数据层推送至接口层
	go heartbeat.StartHeartBeat()
	// 响应绑定在同一个exchange上的信息，查看是否存在该对象
	go locate.StartLocate()

	http.HandleFunc("/objects/", handler.StorageHandler)

	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
