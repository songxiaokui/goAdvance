package main

import (
	"fmt"
	"time"
)

/*
@Time    : 2021/10/3 09:37
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 2.funcTimeOut.go
@Software: GoLand
*/

/*
实现方式:
1. 业务放入到goroutine中
2. 业务结果放入到channel中

func job() string {
	// 业务逻辑很耗时
	// time.Sleep(time.Second*5)
	// return "success"
}
*/

func jobs() chan interface{} {
	result := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 2)
		result <- "success"
	}()
	return result
}

// add a middle func to handle time out
func runHandle() (interface{}, error) {
	select {
	case d := <-jobs():
		return d, nil
	case <-time.After(time.Second * 3):
		return nil, fmt.Errorf("time out\n")
	}
}

func main() {
	fmt.Println(runHandle())
}
