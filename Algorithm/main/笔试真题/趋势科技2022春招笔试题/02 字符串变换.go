/*
题目：
一个字符串a替换成另一个字符串b
可以删除一个字符，可以添加一个字符，
也可以修改一个字符（先删除，再添加，算2次）
问，需要经过多少次替换
*/

package main

import "fmt"

func main() {
	// 	var a, b string
	// 	fmt.Scan(&a)
	// 	fmt.Scan(&b)
	// 	res := solution2(a, b)
	// 	fmt.Println(res)

	strs := [][]string{
		{"com", "cn"},
		{"com", "ccom"},
		{"comn", "cnom"},
	}
	for _, val := range strs {
		a := val[0]
		b := val[1]
		res := solution2(a, b)
		fmt.Println(a, b, res)
	}
}

// 错误方式
// func solution1(a, b string) int {
// 	n, m := len(a), len(b)
//
// 	dp := make([]int, m+1)
// 	for i := 0; i < n; i++ {
// 		for j := 1; j <= m; j++ {
// 			if a[i] == b[j-1] {
// 				dp[j] = max(dp[j], dp[j-1]) + 1
// 			} else {
// 				dp[j] = dp[j-1]
// 			}
// 		}
// 	}
//
// 	count := dp[len(dp)-1]
// 	res := n - count + m - count
// 	return res
// }

func solution2(a, b string) int {
	n, m := len(a), len(b)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// 计算最长公共字串的长度（字串：可以不连续，字符出现的顺序应该一致）
	// 动态规划
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] { // 注意 -1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 放弃a[i-1] 或者 b[j-1]
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	count := dp[n][m]
	res := n - count + m - count
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
