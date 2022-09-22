// 必须定义一个 包名为 `main` 的包，并实现 `main()` 函数。
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg  sync.WaitGroup

	ch := make(chan int, 5)

	cus := func() {
		for {
			var num int
			num = <- ch
			fmt.Println(num)
			if num == 4 {
				wg.Done()
			}
		}

	}

	gen := func() {
		for i:=0;i<5;i++ {
			ch <- i
		}
		wg.Done()
	}
	wg.Add(2)
	go gen()
	go cus()

	wg.Wait()
}


