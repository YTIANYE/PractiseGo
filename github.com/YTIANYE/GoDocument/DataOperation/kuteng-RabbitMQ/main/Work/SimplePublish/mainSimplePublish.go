package main

import (
	"fmt"
	"strconv"
	"time"

	"awesomeProject1/github.com/YTIANYE/GoDocument/DataOperation/kuteng-RabbitMQ/main/Work/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "kuteng")

	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello kuteng!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}