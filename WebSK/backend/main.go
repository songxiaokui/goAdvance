package main

/*
@Time    : 2021/12/2 10:15
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

import (
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	s.GET("/sk")
	if err := s.Run(":8888"); err != nil {
		panic("server failed!")
	}

}
