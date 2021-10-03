package main

import "fmt"

/*
@Time    : 2021/9/16 21:33
@Author  : austsxk
@Email   : austsxk@163.com
@File    : floatFormat.go
@Software: GoLand
*/

func main() {
	var f float32 = 0.085
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%F\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
	fmt.Printf("%8.2f\n", f) // 占8位，保留两位小数
	var c int = 9
	// 用补码进行计算
	fmt.Println(^c)
	fmt.Println(-1 >> 2)

	var d int8 = -8
	fmt.Printf("%b", d)
}
