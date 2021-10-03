package main

import "fmt"

/*
@Time    : 2021/10/3 10:08
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 4.deferParams.go
@Software: GoLand
*/

// defer语句在定义时，对参数进行一次快照操作，如果是匿名函数 则不会确定值，突破defer机制，传递指针
func deferBreak(a *int) {
	fmt.Println(*a)
}

func main() {
	a := 1
	defer fmt.Println(a) // 此时a在defer定义时已经被复制 a=1
	defer func() {       // defer定义没传参数 a为全局变量 a=2
		fmt.Println(a)
	}()
	// 突破defer机制,传递指针
	{
		defer deferBreak(&a) // a=2
	}
	a++
}
