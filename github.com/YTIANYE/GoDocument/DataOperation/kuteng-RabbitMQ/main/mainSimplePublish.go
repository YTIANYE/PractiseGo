package main

import (
	"awesomeProject1/github.com/YTIANYE/GoDocument/DataOperation/kuteng-RabbitMQ/main/RabbitMQ"
	"fmt"


)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "kuteng")
	rabbitmq.PublishSimple("Hello kuteng222!")
	fmt.Println("发送成功！")
}