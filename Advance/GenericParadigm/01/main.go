package main

import "fmt"

/*
@Time    : 2021/8/19 17:03
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/
type Addable interface {
	type int, int64, string
}

func add[T Addable](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add[int](1, 3))
	fmt.Println(add[string]("hello", "world"))
}
