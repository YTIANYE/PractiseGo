package main

import "awesomeProject1/github.com/YTIANYE/GoDocument/DataOperation/kuteng-RabbitMQ/main/Work/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "kuteng")
	rabbitmq.ConsumeSimple()
}
