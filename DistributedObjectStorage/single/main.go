package main

import (
	"log"
	"net/http"
	"single/handler"
)

/*
@Time    : 2021/8/8 12:16
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

func main() {
	http.HandleFunc("/objects/", handler.StorageHandler)
	err := http.ListenAndServe("127.0.0.1:9909", nil)
	if err != nil {
		log.Fatal(err)
	}
}
