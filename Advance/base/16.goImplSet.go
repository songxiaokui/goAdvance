package main

import (
	"bytes"
	"fmt"
	"sync"
)

/*
@Time    : 2021/10/4 13:11
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 16goImplSet.go
@Software: GoLand
*/

/*
使用go实现一个set操作
实现思路:
Map + Mutex
*/

type MySet struct {
	mutex sync.Mutex
	Data  map[interface{}]Empty
}

type Empty struct{}

func (this *MySet) Add(vs ...interface{}) *MySet {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	for _, v := range vs {
		this.Data[v] = Empty{}
	}
	return this
}

func (this *MySet) String() string {
	var buffer bytes.Buffer
	for k, _ := range this.Data {
		if buffer.Len() > 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprintf("%v", k))
	}
	return buffer.String()
}

func NewSet() *MySet {
	return &MySet{
		Data: make(map[interface{}]Empty),
	}
}

func main() {
	set := NewSet()
	set.Add(1, 2, 3, 1, 2, 3, "a", "a", "c", "c", "b", "r")
	fmt.Println(set)
}
