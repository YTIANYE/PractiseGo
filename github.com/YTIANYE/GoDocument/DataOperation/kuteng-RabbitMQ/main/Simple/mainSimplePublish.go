package main

import (
	RabbitMQ2 "awesomeProject1/github.com/YTIANYE/GoDocument/DataOperation/kuteng-RabbitMQ/main/Simple/RabbitMQ"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ2.NewRabbitMQSimple("" + "kuteng")
	rabbitmq.PublishSimple("Hello kuteng222!")
	fmt.Println("发送成功！")
}