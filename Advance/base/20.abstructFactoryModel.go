package main

import "fmt"

/*
@Time    : 2021/10/4 14:29
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 20.abstructFactoryModel.go
@Software: GoLand
*/

/*
抽象工厂模式:
将具备一类属性和方法的对象进行抽象公有的接口，通过标示创建对象
*/

// Stud
type Stud struct {
}

func (this *Stud) GetRole() string {
	return "学生"
}

// Teach
type Teach struct {
}

func (this *Teach) GetRole() string {
	return "教师"
}

// Lead
type Lead struct {
}

func (this *Lead) GetRole() string {
	return "领导"
}

// 抽象接口，都具备获取角色方法
type Person interface {
	GetRole() string
}

type PersonType int

const (
	STUDENT PersonType = iota
	TEACHER
	LEADER
)

func CreatePerson(t PersonType) Person {
	switch t {
	case STUDENT:
		return new(Stud)
	case TEACHER:
		return new(Teach)
	case LEADER:
		return new(Lead)
	default:
		panic("type error")
	}
}

func main() {
	fmt.Println(CreatePerson(1).GetRole())
	fmt.Println(CreatePerson(2).GetRole())
}
