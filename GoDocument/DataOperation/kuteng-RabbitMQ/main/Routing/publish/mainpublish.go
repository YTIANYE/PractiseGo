package main

import (
	RabbitMQ2 "awesomeProject1/GoDocument/DataOperation/kuteng-RabbitMQ/main/Routing/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	kutengone := RabbitMQ2.NewRabbitMQRouting("kuteng", "kuteng_奇数")
	kutengtwo := RabbitMQ2.NewRabbitMQRouting("kuteng", "kuteng_偶数")
	for i := 0; i <= 100; i++ {
		if i % 2 == 0{
			kutengtwo.PublishRouting("Hello kuteng 偶数!" + strconv.Itoa(i))
		}else{
			kutengone.PublishRouting("Hello kuteng 奇数!" + strconv.Itoa(i))
		}
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}