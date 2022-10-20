/**
一个队列n个数，每个数不相等
第i次出队的数是ai（i = 1 ... n）
最终结果是 所有的 ai * 10^(10^n-i) 求和

a,b两个人，每次可以选择从队首或队尾出一个数
a想要最终结果尽可能的大
b想要最终结果尽可能的小
i为奇数时，a做选择
i为偶数时，b做选择

输出最终的出队序列
*/
package main

import "fmt"

// 100%
func main() {
	var n int
	fmt.Scan(&n)
	que := make([]int, n)
	for i := range que {
		fmt.Scan(&que[i])
	}
	// 队首出队
	lpop := func() int {
		num := que[0]
		if len(que) == 1 {
			que = []int{}
		} else {
			que = que[1:]
		}
		return num
	}
	// 队尾出队
	rpop := func() int {
		num := que[len(que)-1]
		que = que[:len(que)-1]
		return num
	}
	//
	res := make([]int, n)
	for i := 1; i <= n; i++ {
		var num int
		if i%2 == 1 {
			if que[0] > que[len(que)-1] {
				num = lpop()
			} else {
				num = rpop()
			}
		} else {
			if que[0] < que[len(que)-1] {
				num = lpop()
			} else {
				num = rpop()
			}
		}
		res[i-1] = num
	}
	//
	fmt.Printf("%d", res[0])
	for i := 1; i < n; i++ {
		fmt.Printf(" %d", res[i])
	}

}
