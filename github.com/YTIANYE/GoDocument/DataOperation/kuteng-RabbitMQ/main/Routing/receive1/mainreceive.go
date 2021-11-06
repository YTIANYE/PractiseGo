package main

import "awesomeProject1/github.com/YTIANYE/GoDocument/DataOperation/kuteng-RabbitMQ/main/Routing/RabbitMQ"

func main() {
	kutengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_偶数")
	kutengtwo.RecieveRouting()
}