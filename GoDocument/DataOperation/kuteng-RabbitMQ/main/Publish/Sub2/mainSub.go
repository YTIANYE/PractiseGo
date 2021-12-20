package main

import (
	RabbitMQ2 "awesomeProject1/GoDocument/DataOperation/kuteng-RabbitMQ/main/Publish/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ2.NewRabbitMQPubSub("" +
		"newProduct")
	rabbitmq.RecieveSub()
}