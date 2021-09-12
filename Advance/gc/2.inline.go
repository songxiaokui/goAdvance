package main

/*
@Time    : 2021/9/6 22:05
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 2.inline.go
@Software: GoLand
*/

// 函数内联
var z *int

//go:noinline
func escape() {
	b := 1
	z = &b
}

// 逃逸分析
