package main

import (
	"fmt"
	"net"
)

/*
@Time    : 2021/12/1 10:13
@Author  : austsxk
@Email   : austsxk@163.com
@File    : server.go
@Software: GoLand
*/

func HandlerMessageAndSendMessage(conn net.Conn) {
	// receive message
	buf := make([]byte, 512)
	_, _ = conn.Read(buf[0:])
	fmt.Println("接收的消息为: ", string(buf))
	_, _ = conn.Write([]byte("hello world"))
}

func main() {
	// init
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:5555")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("tcp address:", tcpAddr)

	// tcp
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return
	}
	// accept message
	for {
		connect, _ := listener.Accept()
		go HandlerMessageAndSendMessage(connect)
	}

}
