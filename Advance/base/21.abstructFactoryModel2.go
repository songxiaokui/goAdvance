package main

import "fmt"

/*
@Time    : 2021/10/4 15:41
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 21.abstructFactoryModel2.go
@Software: GoLand
*/

/*
在20上面，通过类型去创建对象，会存在一个问题：
一旦拓展了新的模块，就要修改工厂创建方法

将创建对象方法继续抽象出接口
*/

// Stud
type Stud2 struct {
}

func (this *Stud2) GetRole() string {
	return "学生"
}

// Teach
type Teach2 struct {
}

func (this *Teach2) GetRole() string {
	return "教师"
}

// Lead
type Lead2 struct {
}

func (this *Lead2) GetRole() string {
	return "领导"
}

// 抽象接口，都具备获取角色方法
type Person2 interface {
	GetRole() string
}

// ----------------------------------------------------------------------------
type AbsFactoryPerson interface {
	CreatePerson() Person2
}

// 此时就要多加struct和方法，多代码量减少耦合
type Stud2Factory struct {
}

func (this *Stud2Factory) CreatePerson() Person2 {
	return &Stud2{}
}

type Teach2Factory struct {
}

func (this *Teach2Factory) CreatePerson() Person2 {
	return &Teach2{}
}

type Lead2Factory struct {
}

func (this *Lead2Factory) CreatePerson() Person2 {
	return &Lead2{}
}

func main() {
	var factory AbsFactoryPerson = new(Lead2Factory)
	fmt.Println(factory.CreatePerson().GetRole())
}
