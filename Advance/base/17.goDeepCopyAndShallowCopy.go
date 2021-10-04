package main

import "fmt"

/*
@Time    : 2021/10/4 13:28
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 17.goDeepCopyAndShallowCopy.go
@Software: GoLand
*/

/*
golang 中的切片在底层是有一个底层数组存储数据

使用 := 实现切片赋值，实际上是一次浅拷贝 如果在没有进行扩容操作，则被拷贝的变量的数据会受到影响

可以使用copy实现深拷贝
*/

func main() {
	a := make([]int, 3, 3)
	a = []int{1, 2, 3}
	// 浅拷贝, 未扩容
	b := a
	a[1] = 100
	fmt.Println(a, b)
	// 对a进行append造成扩容
	a = append(a, 4)
	a[1] = 200
	fmt.Println(a, b) // a [1 200 3 4]  b [1 100 3]
	// 实现深拷贝
	e := make([]int, len(a), cap(a))
	copy(e, a)
	a[1] = 0
	fmt.Println(a, e)

}
