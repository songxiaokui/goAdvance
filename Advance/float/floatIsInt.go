package main

import (
	"fmt"
	"math"
)

/*
@Time    : 2021/9/12 19:42
@Author  : austsxk
@Email   : austsxk@163.com
@File    : floatIsInt.go
@Software: GoLand
*/

// 判断浮点数是否是整数
// bits 小数二进 bias 偏移量 32 127 64 1023
func FloatIsInt(bits uint32, bias int) bool {
	// 获取阶码
	jm := int(bits >> 23)
	// 计算指数
	s := jm - bias

	exponet := s - 23
	fmt.Println(exponet)
	return false

}

func main() {
	var n float32 = 234523.000
	bits := math.Float32bits(n)
	binary, _ := fmt.Printf("%.32b\n", bits)
	fmt.Println(binary)
	c := bits >> 23
	fmt.Println(c)
	d := (bits & ((1 << 23) - 1)) | (1 << 23)
	fmt.Printf("%.32b\n", d)
	fmt.Println(d)
	FloatIsInt(bits, 127)

	// 构造正无穷 Inf 负无穷 -Inf 无效数 NaN
	var c1 float64
	fmt.Println(c1, 1/c1, -1/c1, c1/c1)

	// 单精度当阶码小于-126时 系数部分小于0
	// 双精度当解码小于-1022 系数部分小于0
	// 此时把系数小于1的称为非常规数
	// 系数在1-2之间的为常规数
}
