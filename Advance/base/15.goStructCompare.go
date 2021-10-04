package main

import "fmt"

/*
@Time    : 2021/10/4 12:59
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 15.goStructCompare.go
@Software: GoLand
*/

/*
go 中struct比较需要满足以下情况:
1. struct结构相同，字段相同，类型相同
2. 不能包含不可比较的成员 map slice func
3. 不能使用指针进行比较
*/

type U1 struct {
	Name string
}

type U2 struct {
	Name string
}

type U3 struct {
	Name  string
	Scope []int
}

func main() {
	a1 := U1{Name: "a1"}
	a2 := U2{Name: "a1"}
	// fmt.Println(a1 == a2) // invalid operation: a1 == a2 (mismatched types U1 and U2)
	fmt.Println(a1 == U1(a2)) // true

	a3 := U3{Name: "a", Scope: []int{1, 2, 3}}
	a4 := U3{Name: "a", Scope: []int{1, 2, 3}}
	fmt.Println(a3, a4)
	// fmt.Println(a3 == a4) // invalid operation: a3 == a4 (struct containing []int cannot be compared)

}
