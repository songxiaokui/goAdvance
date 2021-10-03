package main

import (
	"fmt"
	"sync"
)

/*
@Time    : 2021/10/3 10:55
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 8.ConcurrencyModel1.go
@Software: GoLand
*/

/*
1. 最基本的并发模型 sync.WaitGroup + go
2. 使用通信 channel + go 将业务结果通过channel传递
*/

// waitGroup + go
func concurrencyModel1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("-", i)
		go func(b int) {
			defer wg.Done()
			fmt.Println(b * 2)

		}(i)
	}
	wg.Wait()
}

// channel + go
func concurrencyModel2() {
	c := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			go func(b int) {
				c <- b * 2
			}(i)
		}
	}()

	for i := 0; i < cap(c); i++ {
		fmt.Println(<-c)
	}
}

func main() {
	// concurrencyModel1()
	concurrencyModel2()
}
