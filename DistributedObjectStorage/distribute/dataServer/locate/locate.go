package locate

import (
	mq "distribute/rabbitmq"
	"os"
	"strconv"
)

/*
@Time    : 2021/8/8 18:14
@Author  : austsxk
@Email   : austsxk@163.com
@File    : locate.go
@Software: GoLand
*/

func Locate(objectName string) bool {
	_, err := os.Stat(objectName)
	return !os.IsNotExist(err)
}

// 后台不断接受dataServer的exchange，判断是否存在对象
func StartLocate() {
	q := mq.New(os.Getenv("RABBIT_SERVER"))
	defer q.Close()
	// 绑定到dataServers,本地队列接受exchange所以消息
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		// 判断对象名称是否在当前节点数据服务存在
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
