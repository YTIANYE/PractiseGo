/*
golang两个协程交替打印1-100的奇数和偶数，一个协程打印奇数（1 3 5 7 9 ... 999 ），一个协程打印偶数（2 4 6 8 ... 100），要求

1. 打印的数字顺序为1 2 3 4 5，顺序打印
2. main函数主要工作为起协程（也可以初始化一些变量），然后阻塞等待协程执行。
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	// method1()
	// method2()
	method3()
}

// 方法一：利用chan的阻塞	和	for{}
func method1() {
	cha1 := make(chan int, 1)
	cha2 := make(chan int, 1)

	cha1 <- 1

	go func() {
		for i := 1; i < 100; i += 2 {
			<-cha1
			fmt.Println(i)
			cha2 <- 1
		}
	}()

	go func() {
		for i := 2; i <= 100; i += 2 {
			<-cha2
			fmt.Println(i)
			cha1 <- 1
		}
	}()

	// 方式一：阻止main协程的关闭
	for {

	}

	// 方式二：等待一些时间，使得两个协程完成
	// time.Sleep(time.Minute)
}

// 方法二：利用chan传递数据	和	sync.WaitGroup
func method2() {
	// wg := sync.WaitGroup{}
	jiChan := make(chan int, 1)
	ouChan := make(chan int, 1)

	wg.Add(1)
	go func() {
		for {
			if len(jiChan) == 1 {
				time.Sleep(time.Millisecond * 100)
				x := <-jiChan
				fmt.Println(x)
				x++
				ouChan <- x
				if x == 100 {
					wg.Done()
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		for {
			if len(ouChan) == 1 {
				time.Sleep(time.Millisecond * 100)
				x := <-ouChan
				fmt.Println(x)
				x++
				jiChan <- x
				if x == 101 {
					wg.Done()
				}
			}
		}
	}()

	jiChan <- 1
	wg.Wait()
}

// 方法三：利用互斥锁
func method3() {
	x := 1

	wg.Add(1)
	go func() {
		for {
			if x > 100 {
				break
			}
			if x%2 == 1 {
				lock.Lock()
				fmt.Println(x)
				x++
				lock.Unlock()
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for {
			if x > 100 {
				break
			}
			if x%2 == 0 {
				lock.Lock()
				fmt.Println(x)
				x++
				lock.Unlock()
			}
		}
		wg.Done() // 注意不要写在for循环中
	}()

	wg.Wait()
}
