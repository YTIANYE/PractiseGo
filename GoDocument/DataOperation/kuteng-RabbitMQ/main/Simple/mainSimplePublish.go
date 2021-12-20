package main

import (
	"awesomeProject1/GoDocument/DataOperation/kuteng-RabbitMQ/main/Simple/RabbitMQ"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "kuteng")
	rabbitmq.PublishSimple("Hello kuteng222!")
	fmt.Println("发送成功！")
}