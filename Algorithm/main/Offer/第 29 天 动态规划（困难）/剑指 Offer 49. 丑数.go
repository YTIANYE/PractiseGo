package main

import (
	"fmt"
)

/*
丑数计算
因子：2 3 5
[1]
[2,3,5]
[4,6,9,15,25]
[]

*/

// 我实现的精选题解：动态规划
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：6.3 MB, 在所有 Go 提交中击败了25.82%的用户
*/
func nthUglyNumber(n int) int {

	res := []int{1}
	a, b, c := 0, 0, 0

	for i := 1; i < n; i++ {
		num := min(res[a]*2, res[b]*3, res[c]*5)
		if num == res[a]*2 {
			a++
		}
		if num == res[b]*3 {
			b++
		}
		if num == res[c]*5 {
			c++
		}
		res = append(res, num)
	}
	return res[n-1]

}

func min(a, b, c int) int {
	res := a
	if res>b {
		res = b
	}
	if res > c {
		res = c
	}
	return res
}

func main() {
	n := 10
	fmt.Println(nthUglyNumber(n))
}
