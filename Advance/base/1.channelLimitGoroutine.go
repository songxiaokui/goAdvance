package main

import (
	"fmt"
	"sync"
	"time"
)

/*
@Time    : 2021/10/3 09:28
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 1.channelLimitGoroutine.go
@Software: GoLand
*/

func job(index int) {
	time.Sleep(time.Second * 1)
	fmt.Printf("job id is: %d\n", index)
}

var pool chan struct{}

func main() {
	pool = make(chan struct{}, 10)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		pool <- struct{}{}
		go func(i int) {
			defer wg.Done()
			job(i)
			<-pool
		}(i)
	}
	wg.Wait()
}
