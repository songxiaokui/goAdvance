package main

import (
	"fmt"
	"sync"
)

/*
@Time    : 2021/10/4 14:16
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 19.SingletonMode.go
@Software: GoLand
*/

/*
go 内置sync.Once 可以实现单例模式
*/

type Conf struct {
	Port int
}

func NewConf(port int) *Conf {
	return &Conf{port}
}

var c *Conf
var once sync.Once

func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			once.Do(func() {
				c = NewConf(8080)
				fmt.Println("current index:", i)
			})
		}(i)
	}

	fmt.Println(*c)
}
