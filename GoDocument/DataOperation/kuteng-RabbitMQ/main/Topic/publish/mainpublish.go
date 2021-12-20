package main

import (
	RabbitMQ2 "awesomeProject1/GoDocument/DataOperation/kuteng-RabbitMQ/main/Topic/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	kutengOne := RabbitMQ2.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.one")
	kutengTwo := RabbitMQ2.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.two")
	for i := 0; i <= 100; i++ {
		kutengOne.PublishTopic("Hello kuteng topic one!" + strconv.Itoa(i))
		kutengTwo.PublishTopic("Hello kuteng topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}