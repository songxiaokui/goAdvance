package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
@Time    : 2021/11/26 09:59
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main1.go
@Software: GoLand
*/

func DoParse(output string) (string, error) {
	if len(output) == 0 {
		return "", errors.New("结果错误")
	}

	data := strings.Fields(strings.ReplaceAll(output, "\n", ""))
	fmt.Println("转化结果未: ", data)
	if len(data) == 2 {
		return data[1], nil
	}
	return "", errors.New("结果错误")
}

const (
	SlurmStatShellScript = `/usr/bin/sacct -ncbj %v | awk '{printf "%%s %%s\n",$1,$3}'; /usr/bin/squeue -hj %v |awk '{printf "%%s %%s\n", $1, $5}' | paste --`
)

func main() {

	//a := fmt.Sprintf(SlurmStatShellScript, 24, 24)
	//command := exec.Command("sh", "-c", a)
	//command.Env = append(os.Environ())
	//data, err := command.CombinedOutput()
	//fmt.Println("结果:",string(data), "错误:", err)
	//fmt.Println("结果长度:",len(string(data)), "错误:", err)

	s := `44 COMPLETED`
	s1 := `44 COMPLETED 
`
	s2 := ``
	s3 := `d`
	fmt.Println(DoParse(s))
	fmt.Println(DoParse(s1))
	fmt.Println(DoParse(s2))
	fmt.Println(DoParse(s3))
}
