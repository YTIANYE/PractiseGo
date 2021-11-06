package main

import "awesomeProject1/github.com/YTIANYE/GoDocument/DataOperation/kuteng-RabbitMQ/main/Topic/RabbitMQ"

func main() {
	kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "#")
	kutengOne.RecieveTopic()
}