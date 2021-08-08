package locate

import (
	"distribute/rabbitmq"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
@Time    : 2021/8/8 19:50
@Author  : austsxk
@Email   : austsxk@163.com
@File    : locate.go
@Software: GoLand
*/

func Handler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// 查询对象所在的节点
	objectName := strings.Split(r.URL.EscapedPath(), "/")[2]
	locate := Locate(objectName)
	if len(locate) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(locate)
	w.Write(b)
}

// Locate 根据对象名称获取节点地址
func Locate(objectName string) string {
	q := rabbitmq.New(os.Getenv("RABBIT_SERVER"))
	q.Publish("dataServers", objectName)
	c := q.Consume()
	// 从rabbitMq中发布对象名称 查找对象 如果1s没返回 则直接返回空
	go func() {
		time.Sleep(time.Second * 1)
		q.Close()
	}()
	// 从channel中获取数据信息
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(objectName string) bool {
	return Locate(objectName) != ""
}
