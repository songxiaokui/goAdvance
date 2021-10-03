package main

import (
	"fmt"
	"time"
)

/*
@Time    : 2021/10/3 11:10
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 9.concurrencyModel2PubSub.go
@Software: GoLand
*/

// 并发模型 - 生产者消费者模型

func Prod(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i * 2
		time.Sleep(time.Second * 1)
	}
}

func Consumer(c chan int) (single chan struct{}) {
	single = make(chan struct{})
	go func() {
		defer func() {
			single <- struct{}{}
		}()
		for d := range c {
			fmt.Println("接收的数据: ", d)
		}
	}()
	return single
}

func main() {
	c := make(chan int)
	go Prod(c)
	<-Consumer(c)
}
