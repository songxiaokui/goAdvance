package main

import (
	"fmt"
	"net"
)

/*
@Time    : 2021/12/1 10:26
@Author  : austsxk
@Email   : austsxk@163.com
@File    : client.go
@Software: GoLand
*/

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:5555")
	if err != nil {
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}

	// send message
	_, err = conn.Write([]byte("我是austsxk！")[0:])
	if err != nil {
		return
	}

	// receive message
	buf := make([]byte, 512)
	_, _ = conn.Read(buf[0:])
	fmt.Println("接收的消息:", string(buf))

}
