package main

import "net/http"

/*
@Time    : 2021/10/4 15:58
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 22.decoratorHandle.go
@Software: GoLand
*/

/*
装饰器模式
*/

// decorator add some function
func CheckLoginDecorator(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("token") == "" {
			writer.Write([]byte("token error"))
		} else {
			f(writer, request)
		}
	}
}

func show(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", CheckLoginDecorator(show))
	http.ListenAndServe(":9999", nil)
}
