package heartbeat

import (
	mq "distribute/rabbitmq"
	"os"
	"time"
)

/*
@Time    : 2021/8/8 18:08
@Author  : austsxk
@Email   : austsxk@163.com
@File    : heartbeat.go
@Software: GoLand
*/

func StartHeartBeat() {
	q := mq.New(os.Getenv("RABBIT_SERVER"))
	defer q.Close()
	for {
		// 向apiServers的exchange不断推送该服务的心跳，每一个数据节点都要向接口节点推送
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(time.Second * 5)
	}
}
