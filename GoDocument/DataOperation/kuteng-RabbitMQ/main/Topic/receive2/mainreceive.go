package main

import (
	RabbitMQ2 "awesomeProject1/GoDocument/DataOperation/kuteng-RabbitMQ/main/Topic/RabbitMQ"
)

func main() {
	kutengOne := RabbitMQ2.NewRabbitMQTopic("exKutengTopic", "kuteng.*.two")
	kutengOne.RecieveTopic()
}