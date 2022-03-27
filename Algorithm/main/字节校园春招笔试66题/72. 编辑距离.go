/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
 

示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

提示：
0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成

*/

package main

// 我实现的官方题解:动态规划
/*
执行用时：4 ms, 在所有 Go 提交中击败了79.20%的用户
内存消耗：5.4 MB, 在所有 Go 提交中击败了71.20%的用户
*/

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	// 有一个字符串是空
	if n*m == 0 {
		return n + m
	}

	// DP数组
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	// 边界状态初始化
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}

	// 计算所有DP的值
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftdown := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftdown += 1
			}
			dp[i][j] = min(left, min(leftdown, down))
		}
	}
	return dp[n][m]
}
