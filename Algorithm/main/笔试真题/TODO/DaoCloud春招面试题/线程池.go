// package main
//
// import "fmt"
//
// type Chi struct{
// 	count int
// 	sum int
//
// 	queue chan int
//
// }
//
//
//
// func (this *Chi) Add(){
// 	if this.count < this.sum{
// 		this.count++
// 		this.queue <- this.count
// 		go work()
// 	}
//
// }
//
// func(this *Chi) Done(){
// 	if this.count == 0{
// 		return
// 	}
// 	this.count--
// 	select {
//
// 	}
// }
//
// func work(){
// 	fmt.Println("hello")
// }
//
// func main(){
//
// 	var c Chi
// 	c.sum = 10
// 	c.count = 0
// 	c.queue = make(chan int, c.sum)
//
// 	for {
// 		c.Add()
// 	}
//
//
// }

package main

import (
	"fmt"
	"runtime"
	"time"
)
// istio
func main() {
	dataChan := make(chan int, 100)
	go func() {
		for {
			select {
			case data := <-dataChan:
				fmt.Println("data:", data)
				// 这里延迟是模拟处理数据的耗时
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 填充数据
	curr := 0
	for i := 0; i < curr; i++ {
		dataChan <- i
	}

	// 这里循环打印查看协程个数
	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}