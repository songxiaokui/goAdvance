package main

import (
	"fmt"
	"os"
	"strings"
)

/*
@Time    : 2021/11/16 16:38
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

func ReadLogContainsEndMark(path, mark string, threshold int64) bool {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return false
	}
	buf := make([]byte, threshold)
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	start := stat.Size() - threshold
	if start < 0 {
		start = 0
	}
	_, err = file.ReadAt(buf, start)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		if strings.Contains(line, mark) {
			return true
		}
	}
	return false
}

func readConvertFile(name string) bool {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return false
	}
	buf := make([]byte, 250)
	stat, err := os.Stat(name)
	if err != nil {
		return false
	}
	fmt.Println(stat.Size())
	start := stat.Size() - 250
	if start < 0 {
		start = 0
	}
	_, err = file.ReadAt(buf, start)

	fmt.Println("---", string(buf))
	lines := strings.Split(string(buf), "\n")
	fmt.Println(lines)
	for _, line := range lines {

		if strings.Contains(line, "END OF CALCULATION") {
			fmt.Println("ok")
			return true
		}

	}
	return false
}

func main() {
	fmt.Println(ReadLogContainsEndMark("run_solver.log", "Destroying", 200))
}
