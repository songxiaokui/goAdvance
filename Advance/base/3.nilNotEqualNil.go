package main

import (
	"fmt"
	"reflect"
)

/*
@Time    : 2021/10/3 09:56
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 3.nilNotEqualNil.go
@Software: GoLand
*/

// 业务判断nil，任何数据都会转化为接口，接口在底层分两部分(value, type)
// 只有value、type都为nil，才能判断nil

// 判断方法
// 1.断言
// 2.反射
func main() {
	var a func()
	var b *struct{}
	var c *string
	vList := []interface{}{a, b, c}
	for _, item := range vList {
		fmt.Println(item)
		if item == nil {
			fmt.Printf("%v is nil", item)
		}
		// 方案1
		{
			if v, ok := item.(func()); ok && v == nil {
				fmt.Println("nil func")
			}
			if v, ok := item.(*struct{}); ok && v == nil {
				fmt.Println("nil struct")
			}
			if v, ok := item.(*string); ok && v == nil {
				fmt.Println("nil string")
			}
		}
		// 方案2
		{
			if reflect.ValueOf(item).IsNil() {
				fmt.Printf("%v is nil\n", item)
			}
		}
	}
}
