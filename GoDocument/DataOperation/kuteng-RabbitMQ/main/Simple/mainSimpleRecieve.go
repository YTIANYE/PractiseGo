package main

import (
	"awesomeProject1/GoDocument/DataOperation/kuteng-RabbitMQ/main/Simple/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "kuteng")
	rabbitmq.ConsumeSimple()
}