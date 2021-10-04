package main

import "fmt"

/*
@Time    : 2021/10/4 12:51
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 14.goEnmuIota.go
@Software: GoLand
*/

/*
golang 中枚举类型使用iota，没有enum关键词

以下是常见使用方法
*/

type UserType int

const (
	Student UserType = iota
	Teacher
	Leader
)

func (this UserType) String() string {
	switch this {
	case 0:
		return "学生"
	case 1:
		return "教师"
	case 2:
		return "领导"
	default:
		return "类型未知"
	}
}

func main() {
	fmt.Println(UserType(Student), UserType(Teacher), UserType(Leader))
}
