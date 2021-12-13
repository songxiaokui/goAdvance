package main

import (
	"fmt"
	"os/exec"
)

/*
@Time    : 2021/11/11 17:35
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

func ExecString(workDir string, cmd string) (string, error) {
	command := exec.Command("sh", "-c", cmd)
	command.Dir = workDir
	data, err := command.CombinedOutput()
	return string(data), err
}
func main() {
	//command := exec.Command("sh", "-c", "sleep.sh > hhsy.log & echo $! > pid.txt")
	////command.Env = append(os.Environ(), config.Custom.Env.MPIEnv, config.Custom.Env.VNCEnv)
	//command.Dir = "/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/Command"
	//bytes, err := command.CombinedOutput()
	//fmt.Println(string(bytes), err)
	d, e := ExecString("/Users/austsxk/github_sxk_repo/goAdvance/goAdvance/Command", "sleep.sh > hhsy.log & echo $! > pid.txt")
	fmt.Println(d, e)
}
