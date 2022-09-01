package main

import "fmt"

func kill(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	left := 2
	for left < n {
		left *= 2
	}
	left /= 2
	num := left + 1
	res := 2
	for num < n {
		res += 2
		num++
	}

	return res
}

/**

8
# 8
9
# 2


16
# 16
*/

func main() {
	for i := 1; i < 17; i++ {
		fmt.Println(i,kill(i))
	}
	//fmt.Println(kill(16))
}
