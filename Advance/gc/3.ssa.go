package main

/*
@Time    : 2021/9/6 23:27
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 3.ssa.go
@Software: GoLand
*/

// 静态单赋值

var d uint8

func main() {
	var a uint8 = 1
	a = 2
	if true {
		a = 3
	}
	d = a
}
