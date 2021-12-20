package main

import (
	RabbitMQ2 "awesomeProject1/GoDocument/DataOperation/kuteng-RabbitMQ/main/Work/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ2.NewRabbitMQSimple("" + "kuteng")
	rabbitmq.ConsumeSimple()
}
