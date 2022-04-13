/*
题目：

给定一个字符串
如果两个相邻位置的字符相等或为相邻字符，可以标记一下并记录分数
如： ’a' 分数为1， ‘z'分数为26
不可以重复标记
问：最大得分是多少？
*/

package main

import "fmt"

/*
abcd d

f(n) = f(n-1)				f(n-1)的最后一个被标记，且不拆散最后两个
f(n) = f(n-3) + fenshu


f(n) = f(n-2) + fenshu 		f(n-1)的最后一个没被标记

拼合
f(n) = f(n-2) + fenshu  ||  f(n-3) + fenshu
不拼合
f(n) = f(n-1)

*/

// 我的题解：未完全通过样例

func main() {
	var s string
	fmt.Scan(&s)

	if len(s) == 0 || len(s) == 1 {
		fmt.Println(0)
	} else if len(s) == 2 {
		fenshu := isRight(s[0], s[1])
		fmt.Println(fenshu)
	} else if len(s) == 3 {
		fenshu := max(isRight(s[0], s[1]), isRight(s[1], s[2]))
		fmt.Println(fenshu)
	} else { // len(s) > 3
		dp := make([]int, len(s))
		dp[1] = isRight(s[0], s[1])
		dp[2] = max(isRight(s[0], s[1]), isRight(s[1], s[2]))

		for i := 3; i < len(s); i++ {
			fenshu := isRight(s[i], s[i-1])
			if fenshu != 0 {
				dp[i] = max(dp[i-2], dp[i-3]) + fenshu // 使用分数
				// 可能是少了以下的情况就没有AC
				dp[i] = max(dp[i], dp[i-1]) // 不使用分数

				// a b c d
			} else {
				dp[i] = dp[i-1]
			}
		}
		fmt.Println(dp[len(s)-1])
	}
}

func isRight(a, b uint8) int {
	if a-b == 1 || b-a == 1 || a == b {
		return int(a - 'a' + b - 'a' + 2)
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
