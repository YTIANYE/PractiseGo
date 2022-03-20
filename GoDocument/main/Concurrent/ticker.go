package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			// <-ticker.C
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				// 停止
				ticker.Stop() // 不在继续执行
			}
		}
	}()
	for {
	}
}
