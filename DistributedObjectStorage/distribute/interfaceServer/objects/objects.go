package objects

import (
	"log"
	"net/http"
)

/*
@Time    : 2021/8/8 19:50
@Author  : austsxk
@Email   : austsxk@163.com
@File    : objects.go
@Software: GoLand
*/

func Handler(w http.ResponseWriter, r *http.Request) {
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
