package main

import "fmt"

/*
@Time    : 2021/10/3 10:35
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 6.deferPanic.go
@Software: GoLand
*/

// defer panic 嵌套的执行情况
func main1() {
	defer fmt.Println("前")
	defer fmt.Println("中")
	defer fmt.Println("后")
	panic("异常触发1")

}

func main2() {
	func() {
		defer fmt.Println("前")
		defer fmt.Println("中")
		defer fmt.Println("后")
		panic("异常触发1")
	}()

	// panic不会执行
	panic("异常触发2")
}

func main() {
	defer func() {
		defer fmt.Println("前")
		defer fmt.Println("中")
		defer fmt.Println("后")
		panic("异常触发1")
	}()

	// panic2会执行
	panic("异常触发2")
}
