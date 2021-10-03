package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
@Time    : 2021/10/3 10:43
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 7.n++.go
@Software: GoLand
*/

// data race 变量自增的坑与解决方案
func increase() {
	n := 0
	for i := 0; i < 1000000; i++ {
		n++
	}
	fmt.Printf("last n is %d", n)
}

// 并发执行会产生data race n++不是原则操作
func increase2() {
	n := 0
	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n++
		}()

	}
	wg.Wait()
	fmt.Printf("last n is %d", n)
}

// 解决方案1 加锁
func increase3() {
	n := 0
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mutex.Unlock()
			mutex.Lock()
			n++
		}()

	}
	wg.Wait()
	fmt.Printf("last n is %d", n)
}

// 解决方案2 原子操作 atomic
func increase4() {
	var n int64
	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&n, 1)
		}()

	}
	wg.Wait()
	fmt.Printf("last n is %d", n)
}
func main() {
	// increase()
	// increase2()
	// increase3()
	increase4()
}
