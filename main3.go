package main

import "net/http"

/*
@Time    : 2021/12/2 16:15
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main3.go
@Software: GoLand
*/

func main() {
	r := &http.Response{}
	r.Body.Close()
}
