package main

import "fmt"

/*
@Time    : 2021/10/4 13:45
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 18.GoEscapeAnalysis.go
@Software: GoLand
*/

/*
堆：
一般来讲是人为手动进行管理，手动申请、分配、释放。一般所涉及的内存大小并不定，
一般会存放较大的对象。另外其分配相对慢，涉及到的指令动作也相对多

栈：由编译器进行管理，自动申请、分配、释放。一般不会太大，
我们常见的函数参数（不同平台允许存放的数量不同），局部变量等等都会存放在栈上

逃逸分析:
逃逸分析就是确定一个变量要放堆上还是栈上

分析阶段: 编译器阶段 go 在编译阶段确立逃逸，注意并不是在运行时

常见场景:
1. 局部变量，在栈中分配与释放，一旦产生外部引用，便会产生逃逸；如果函数返回外部不使用变量，就是无意的逃逸
2. 指针引用，肯定会发生逃逸（影响的是变量权重）
*/

type Escape struct {
	Name string
	Age  int
}

func New(name string, age int) *Escape {
	return &Escape{
		Name: name,
		Age:  age,
	}
}

func Use(c *Escape) {

}
func main() {
	a := New("austsxk", 18)
	Use(a)
	fmt.Println(a)
}

//  go run  -gcflags "-m " 18.GoEscapeAnalysis.go
