package main

/*
@Time    : 2021/10/4 12:39
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 13.goContainerList.go
@Software: GoLand
*/
import (
	"container/list"
	"fmt"
)

/*
go 中的链表操作，双向链表，定义在container.list中
*/

func main() {
	l := list.New()
	l.Init()
	fmt.Println("链表初始化: ", l.Len(), l.Front(), l.Back())
	l.PushBack(1)
	x := l.PushBack(2)
	l.PushBack(3)
	fmt.Println("链表增加元素: ", l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("元素:", e.Value)
	}

	// 首部插入
	l.PushFront(0)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("元素:", e.Value)
	}

	// 在某个元素之前插入
	l.InsertBefore(1.5, x)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("元素:", e.Value)
	}
	// 在某个元素之后插入
	l.InsertAfter(2.5, x)
	fmt.Println("--------")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("元素:", e.Value)
	}
}
