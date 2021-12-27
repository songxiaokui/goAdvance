package main

import (
	"images/cmd/image"
	"log"
)

/*
@Time    : 2021/12/27 14:46
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

func main() {
	err := image.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err.Error())
	}
}
