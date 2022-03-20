package main

import (
	"fmt"
	"sync"
	"time"
)



var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

// 互斥锁
func mutexFunc1(){
	var add func(int)
	add = func(id int) {
		for i:=0;i<100;i++{
			lock.Lock()
			x++
			fmt.Println(id, ":", x)
			lock.Unlock()
			time.Sleep(time.Second)
		}
		wg.Done()
	}

	wg.Add(2)
	go add(1)
	go add(2)
	wg.Wait()
	fmt.Println(x)
}

// 读写锁
func rwmutexFunc(){

	var read func()
	read = func(){
		rwlock.Lock()
		time.Sleep(time.Millisecond)
		fmt.Println("read:", x)
		rwlock.Unlock()
		wg.Done()
	}

	var write func()
	write = func() {
		rwlock.Lock()
		time.Sleep(time.Millisecond)
		x++
		fmt.Println("write:", x)
		rwlock.Unlock()
		wg.Done()
	}

	start := time.Now()
	for i:=0;i<10;i++{
		wg.Add(1)
		go write()
	}

	for i:=0;i<50;i++{
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}




func main(){
	// mutexFunc1()
	rwmutexFunc()
}
