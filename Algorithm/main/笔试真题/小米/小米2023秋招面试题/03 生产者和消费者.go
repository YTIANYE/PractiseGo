/**
实现一个生产者和消费者
 */
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	n := 5
	chan1 := make(chan int, n)
	//
	var cus func()
	cus = func() {
		for i:=0;i<n;i++ {
			fmt.Println(<- chan1)
		}
		wg.Done()
	}
	//
	var gen func()
	gen = func() {
		for i:=0;i<n;i++ {
			chan1 <- i
		}
		wg.Done()
	}
	//
	wg.Add(2)
	go gen()
	go cus()
	wg.Wait()

}

