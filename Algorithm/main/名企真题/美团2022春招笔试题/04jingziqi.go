/*
题目：
一个田字的四个方框涂颜色，
相同的颜色不能在一行或者一列
给定n个颜色，问有多少种涂颜色的方式？

例子：
n = 2

1 0
0 1

0 1
1 0

共有2种
*/

package main

import "fmt"

/**
2 个颜色
2种

3个颜色
12 种

4个颜色
4*3*2 = 24

*/

/*

0
0
2
3 18 = 6 + 12
4 84 = 6 * 2 + 4 * 12 + 24
260
630
1302
2408
4104
6570
*/

// 我的题解：通过样例 91%

func main() {
	var n int
	fmt.Scan(&n)

	res := 0
	if n <= 1 {
		fmt.Println(0)
	} else { // 1000000007
		if n >= 2 {
			res = (res + (chose(n, 2)*2)%1000000007) % 1000000007
		}
		if n >= 3 {
			res = (res + (chose(n, 3)*12)%1000000007) % 1000000007
		}
		if n >= 4 {
			res = (res + (chose(n, 4)*24)%1000000007) % 1000000007
		}
	}
	fmt.Println(res)

}

func chose(n, m int) int { // 从n个数中选取m个
	// C 5 2 = 5 * 4 / 2
	// C 7 3 = 7 * 6 * 5 / (3 * 2 * 1 )
	num1 := 1
	num2 := 1
	for i := 1; i <= m; i++ {
		num1 *= n
		n--
		num2 *= i
	}
	return num1 / num2
}
