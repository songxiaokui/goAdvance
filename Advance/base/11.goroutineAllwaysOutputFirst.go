package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
@Time    : 2021/10/3 13:05
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 11.goroutineAllwaysOutputFirst.go
@Software: GoLand
*/

// goroutine总是输出的第一个值是最后一个数据

/*
这种情况只存在在单核运行的情况，go默认使用当前机器的全部CPU执行

原因是 goroutine是有GPM模型的一个系统级线程执行，线程维护一个运行上下文P
P准备就绪后，线程才开始执行，默认执行的是最后创建的那一个协程
然后再执行其他的协程 并且依次按照创建时的顺序执行

*/

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 9; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	// resolve
	go func() {
		defer wg.Done()
		fmt.Println("准备输出...")
	}()
	wg.Wait()

}
