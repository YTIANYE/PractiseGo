/*
数字以0123456789101112131415…的格式序列化到一个字符序列中。
在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。

示例 1：

输入：n = 3
输出：3
示例 2：

输入：n = 11
输出：0

限制：

0 <= n < 2^31

*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
0		1			0
9		9*1			1-9
189		90*2		10-99
2889	900*3		100-999
9000*4		1000-9999

*/

// 我的题解：

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了71.11%的用户
*/

func findNthDigit1(n int) int {
	if n < 10 {
		return n
	}

	last := 0
	curr := 0
	count := 0 // 当前是几位数

	for curr < n {
		last = curr
		curr += 9 * int(math.Pow(10, float64(count))) * (count + 1)
		count++
	}

	start := int(math.Pow(10, float64(count-1))) // 开始的数
	var num int

	num = start - 1 + (n-last+1)/count // 当前的数

	// 【有第0位，第n位是第n+1个位】
	weizhi := (n - last - 1) % count // num中的第几位,从0开始
	str := strconv.Itoa(num)
	res, _ := strconv.Atoi(string(str[weizhi]))
	return res
}

// 我实现的精选题解
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.8 MB, 在所有 Go 提交中击败了93.70%的用户
*/

func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	var num int
	for n > count { //  1.
		n -= count
		start *= 10
		digit += 1
		count = 9 * start * digit
	}
	num = start + (n-1)/digit                        // 2.
	return int(strconv.Itoa(num)[(n-1)%digit] - '0') // 3.
}

func main() {

	// n := 30 // 2
	// n := 10 // 1
	// n := 11 // 0
	// n := 270 // 6
	n := 30

	fmt.Println(findNthDigit(n))

}
