package handler

import (
	"log"
	"net/http"
)

/*
@Time    : 2021/8/8 12:16
@Author  : austsxk
@Email   : austsxk@163.com
@File    : handler.go
@Software: GoLand
*/

// StorageHandler 处理单机REST接口Handler
func StorageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("接受请求: ", r.RemoteAddr)
	switch r.Method {
	case http.MethodPut:
		// up load
		put(w, r)
		return
	case http.MethodGet:
		// download object
		get(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
