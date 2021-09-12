package main

import (
	"fmt"
	"math"
)

/*
@Time    : 2021/9/12 19:25
@Author  : austsxk
@Email   : austsxk@163.com
@File    : float.go
@Software: GoLand
*/

func main() {
	bias := 127
	sing := 1 << 31
	fmt.Println(bias, sing)
	fmt.Println(bias & sing)

	var n float32 = 0.085
	bits := math.Float32bits(n)
	binary := fmt.Sprintf("%.32b", bits)
	fmt.Println(bits)
	// 右移动23位 获取阶码
	exponentRaw := int(bits >> 23)
	fmt.Println(exponentRaw)

	// 阶码=指数+127
	expoent := exponentRaw - bias
	fmt.Println(expoent)

	// 小数部分 binary[9:]
	var m float64
	for index, value := range binary[9:] {
		if byte(value) == '1' {
			position := index + 1
			bitValue := math.Pow(2, float64(position))
			m += 1 / bitValue
		}
	}
	fmt.Println(m)
	// 最后值
	value := (1 + m) * math.Pow(2, float64(expoent))
	fmt.Println(value)
}
