package heartbeat

import (
	"distribute/rabbitmq"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

/*
@Time    : 2021/8/8 19:02
@Author  : austsxk
@Email   : austsxk@163.com
@File    : heartbeat.go
@Software: GoLand
*/

var dataServers = make(map[string]time.Time)
var mutex sync.RWMutex

func ListenHeartbeat() {
	fmt.Println("RabbitMQ server", os.Getenv("RABBIT_SERVER"))
	q := rabbitmq.New(os.Getenv("RABBIT_SERVER"))

	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()
	// 异步处理过期数据节点
	go removeExpiredDataPoint()
	for msg := range c {
		dataServer, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

// 处理过期数据节点
func removeExpiredDataPoint() {
	for {
		time.Sleep(time.Second * 5)
		mutex.RLock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.RUnlock()
	}
}

// 获取数据节点列表
func GetServersList() []string {
	serversList := make([]string, 0)
	mutex.RLock()
	defer mutex.RUnlock()
	for k, _ := range dataServers {
		serversList = append(serversList, k)
	}
	return serversList
}

// 随机选择一个服务节点
func RandomSelectDataServersPoint() string {
	ds := GetServersList()
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
