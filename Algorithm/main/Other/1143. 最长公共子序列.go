package main

import "fmt"

// 我的题解：动态规划
/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：10.5 MB, 在所有 Go 提交中击败了88.70%的用户
*/

func longestCommonSubsequence1(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	//  表示[i][j]之前的字符串的最大公共字串的长度
	dp := make([][]int, n+1) // 注意 +1
	for i := range dp {
		dp[i] = make([]int, m+1) // 注意 +1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] { // 注意 -1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
		fmt.Println(dp[i])
	}

	res := dp[n][m]
	return res
}

// 错误方式：一维dp数组
func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)
	dp := make([]int, m+1) // 注意 +1

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] { // 注意 -1
				dp[j] = dp[j-1] + 1 // 存在错误，容易计算过量
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
		}
		fmt.Println(dp)
	}

	res := dp[m]
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a, b := "abcba", "abcbcba" // 5

	// 错误方式
	res := longestCommonSubsequence(a, b)
	fmt.Println("res:", res)

	// 正确方式
	res1 := longestCommonSubsequence1(a, b)
	fmt.Println("res1:", res1)
}
