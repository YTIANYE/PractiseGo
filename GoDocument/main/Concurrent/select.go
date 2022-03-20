package main

import (
	"fmt"
	"time"
)

func main() {
	selectFunc2()

}

// 多个 chan 同时 ready ,随机选择一个执行
func selectFunc1() {
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	go func() {
		intChan <- 1
	}()
	go func() {
		stringChan <- "hello"
	}()

	select {
	case val := <-intChan:
		fmt.Println("int:", val)
	case str := <-stringChan:
		fmt.Println("string:", str)
	}
	fmt.Println("end!")
}

// 判断管道是不是满了
func selectFunc2() {

	writeChan := make(chan string, 2)
	go func(writeChan chan string) {
		for {
			select {
			case writeChan <- "hello":
				fmt.Println("write")
			default:
				fmt.Println("满了")
				break
			}
			time.Sleep(time.Millisecond * 500)
		}

	}(writeChan)

	for s := range writeChan {
		fmt.Println("read:", s)
		time.Sleep(time.Second)
	}

}
