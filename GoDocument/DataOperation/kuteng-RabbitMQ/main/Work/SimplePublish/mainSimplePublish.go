package main

import (
	RabbitMQ2 "awesomeProject1/GoDocument/DataOperation/kuteng-RabbitMQ/main/Work/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ2.NewRabbitMQSimple("" + "kuteng")

	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello kuteng!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}