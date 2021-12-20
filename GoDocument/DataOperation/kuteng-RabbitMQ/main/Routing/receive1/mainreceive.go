package main

import (
	RabbitMQ2 "awesomeProject1/GoDocument/DataOperation/kuteng-RabbitMQ/main/Routing/RabbitMQ"
)

func main() {
	kutengtwo := RabbitMQ2.NewRabbitMQRouting("kuteng", "kuteng_偶数")
	kutengtwo.RecieveRouting()
}