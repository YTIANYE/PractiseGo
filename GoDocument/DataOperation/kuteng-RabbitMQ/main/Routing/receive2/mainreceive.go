package main

import (
	RabbitMQ2 "awesomeProject1/GoDocument/DataOperation/kuteng-RabbitMQ/main/Routing/RabbitMQ"
)

func main() {
	kutengone := RabbitMQ2.NewRabbitMQRouting("kuteng", "kuteng_奇数")
	kutengone.RecieveRouting()
}