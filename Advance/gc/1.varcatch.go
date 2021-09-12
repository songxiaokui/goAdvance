package main

import "fmt"

/*
@Time    : 2021/9/6 21:59
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 1.varcatch.go
@Software: GoLand
*/

// 变量捕获
// 变量捕获是对闭包场景而言 变量捕获需要明确变量在闭包中是值引用 还是地址引用

func main() {
	a := 1
	b := 2
	go func() {
		fmt.Println(a, b)
	}()
	a = 100
}

// go tool compile -m=2 1.varcatch.go |grep capturing
/*
1.varcatch.go:20:15: main.func1 capturing by ref: a (addr=true assign=true width=8)
1.varcatch.go:20:18: main.func1 capturing by value: b (addr=false assign=false width=8)

在闭包内引入闭包外的 a b变量，在闭包后对变量a进行赋值操作；所以在闭包中对a、b处理不同
a: 地址引用传递
b: 值传递
*/
