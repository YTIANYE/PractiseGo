/*
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。


示例 1：

输入：n = 12
输出：5
示例 2：

输入：n = 13
输出：6


限制：

1 <= n < 2^31

*/

package main

import (
	"fmt"
	"math"
)

// 精选题解
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.8 MB, 在所有 Go 提交中击败了87.57%的用户
*/

func countDigitOne(n int) int {
	res := 0       // 结果
	curr := n % 10 // 当前位
	high := n / 10
	low := 0
	dig := 1 // 当前 位

	for high != 0 || curr != 0 {
		if curr == 0 {
			res += high * dig
		} else if curr == 1 {
			res += high*dig + low + 1
		} else {
			res += (high + 1) * dig
		}
		low += curr * dig
		curr = high % 10
		high /= 10
		dig *= 10
	}
	return res
}

func main() {
	num := int(math.Pow(2, 31))
	fmt.Println(num)
	// 12   1  10 11 12
	fmt.Println(countDigitOne(13))
}
