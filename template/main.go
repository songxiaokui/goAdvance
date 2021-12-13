package main

import (
	"fmt"
	"path"
	"path/filepath"
)

/*
@Time    : 2021/11/17 10:01
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/
var config = `
app.name = ${appName}
app.ip = ${appIP}
app.port = ${appPort}
`

var dev = map[string]string{
	"appName": "my_app",
	"appIP":   "0.0.0.0",
	"appPort": "8080",
}

type A struct {
	appName string
	appIP   string
}

func main() {
	// /home/a.py/ss/s/s/ss
	a := "home/a/s/sss.py"

	s := filepath.Dir(a)

	d := path.Join("1", s)
	fmt.Println(d)
}
