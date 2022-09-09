/**
题目：
n中配料的价值存在一个数组中，
随机选取3中配料，搭配一杯奶茶
问一杯奶茶平均配料的价值是多少，输出分子和分母
 */

package main

import "fmt"

// 100%
func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		sum += num * 3
	}
	sum, n = fun(sum, n)
	fmt.Println(sum, n)

}

func fun(sum, n int) (int, int) {

	for i := 2; i <= n; i++ {
		if sum%i == 0 && n%i == 0 {
			sum /= i
			n /= i
			i = 1
		}
	}
	return sum, n
}
