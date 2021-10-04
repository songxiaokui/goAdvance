package main

import (
	"fmt"
	"sync"
	"time"
)

/*
@Time    : 2021/10/4 10:42
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 12.kvMapAndHasExpireTime.go
@Software: GoLand
*/

/*
实现一个线程安全的 k-v map，并可是设置过期时间

实现策略: sync.Map + time.afterFunc
*/

var kv sync.Map

func Set(key string, value interface{}, expire time.Duration) {
	kv.Store(key, value)
	time.AfterFunc(expire, func() {
		kv.Delete(key)
	})
}

func main() {
	Set("id", 1, time.Second*3)
	Set("name", "austsxk", time.Second*5)
	for {
		fmt.Println(kv.Load("id"))
		fmt.Println(kv.Load("name"))
		time.Sleep(time.Second * 1)
	}
}
