package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
@Time    : 2021/10/3 12:53
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 10.concurrencyModel3SurvivalOfTheFittest.go
@Software: GoLand
*/
// go 并发模型之优胜劣汰模式
/*
所谓优胜劣汰就是服务器资源不可控，获取某些资源，使用多goroutine请求，只要获取到一个，即结束
*/

func job2() int {
	rand.Seed(time.Now().Unix())
	result := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(result))
	return result
}

func main() {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go func() {
			c <- job2()
		}()
	}
	fmt.Println("最快获取数据时间为: ", <-c)
}
