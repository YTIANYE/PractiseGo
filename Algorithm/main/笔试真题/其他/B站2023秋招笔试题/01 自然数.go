/**
一个数可以拆分成若干个数相乘，求这些数相加的最小结果是什么？
 */

package main

import "fmt"

// 100%
func main() {
	var n int
	fmt.Scan(&n)
	var res int
	var i int
	for  {
		n, i = comp(n)
		//fmt.Println(n, i)
		if i == 1 {
			break
		}
		res += i
	}
	res += n
	fmt.Println(res)
}

func comp(n int ) (int, int) {
	for i := 2;i<n;i++ {
		if n % i == 0 {
			return n/i, i
		}
	}
	return n, 1
}
