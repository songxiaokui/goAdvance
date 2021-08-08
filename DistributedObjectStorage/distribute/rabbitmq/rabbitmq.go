package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

/*
@Time    : 2021/8/8 15:55
@Author  : austsxk
@Email   : austsxk@163.com
@File    : rabbitmq.go
@Software: GoLand
*/

// RabbitMQ

type RabbitMQ struct {
	channel  *amqp.Channel
	Name     string
	Exchange string
}

// New 创建一个RabbitMQ
func New(s string) *RabbitMQ {
	// 使用amqp url创建一个消息队列连接
	conn, err := amqp.Dial(s)
	if err != nil {
		panic(err)
	}
	// 获取当前连接对象的channel用来接受处理消息
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	q, err := ch.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	mq := new(RabbitMQ)
	mq.channel = ch
	mq.Name = q.Name
	return mq
}

// Bind 将自己的消息队列与Exchange绑定，所有发往该exchange的消息都可以被该消息队列接受
func (rm *RabbitMQ) Bind(exchange string) {
	err := rm.channel.QueueBind(
		rm.Name,  // 队列名称
		"",       // 路由key
		exchange, // 交换机名称
		false,    // 是否等待
		nil,      // 参数
	)
	if err != nil {
		panic(err)
	}
	rm.Exchange = exchange
}

// Send 向某个消息队列中发送消息
func (rm *RabbitMQ) Send(queue string, body interface{}) {
	data, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	err = rm.channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ReplyTo: rm.Name,
			Body:    []byte(data),
		},
	)
	if err != nil {
		panic(err)
	}
}

// Publish 向某个exchange中推送消息
func (rm *RabbitMQ) Publish(exchange string, body interface{}) {
	data, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	err = rm.channel.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ReplyTo: rm.Name,
			Body:    []byte(data),
		},
	)
	if err != nil {
		panic(err)
	}
}

// Consume 用来接受消息
func (rm *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, err := rm.channel.Consume(
		rm.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	return c
}

// Close 关闭channel
func (rm *RabbitMQ) Close() {
	rm.channel.Close()
}
