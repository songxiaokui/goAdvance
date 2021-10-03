package main

import "fmt"

/*
@Time    : 2021/10/3 10:20
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 5.deferLinkCallAndForLoop.go
@Software: GoLand
*/

// defer链式调用与for循环调用注意问题
/*
结论:
1. 链式调用，在defer定义时，只会延时执行最后调用的一次，也就是说前面调用会在被定义时依次执行
2. for循环调用时，注意最后值，是进行地址拷贝
*/
type myStruct struct {
}

func (this *myStruct) do(index int) *myStruct {
	fmt.Printf("%v->", index)
	return this
}
func main() {
	// 链式调用
	{
		a := myStruct{}
		/*
			defer a.do(1).do(2).do(3)
			a.do(4)
			// 1->2->4->3
		*/
		defer func() {
			a.do(1).do(2).do(3)
		}()
		a.do(4)

	}
	// defer 循环调用
	{
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println(i)
			}() // 3 3 3
		}

		for i := 0; i < 3; i++ {
			defer fmt.Println(i) // 2 1 0
		}
	}
}
