package main

import (
	RabbitMQ2 "awesomeProject1/github.com/YTIANYE/GoDocument/DataOperation/kuteng-RabbitMQ/main/Simple/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ2.NewRabbitMQSimple("" + "kuteng")
	rabbitmq.ConsumeSimple()
}