package main

/*
@Time    : 2021/9/7 22:38
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 5.elf.go
@Software: GoLand
*/

import (
	"debug/elf"
	"log"
)

func main() {
	f, err := elf.Open("main")
	if err != nil {
		log.Fatal(err)
	}
	for _, section := range f.Sections {
		log.Println(section)
	}
}

// objdump -s -j .note.go.buildid 4.link
// 查找每个Go程序唯一id
