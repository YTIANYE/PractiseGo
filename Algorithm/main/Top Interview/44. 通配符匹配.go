/**
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
示例 3:

输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
示例 4:

输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
示例 5:

输入:
s = "acdcb"
p = "a*c?b"
输出: false

*/

package main

// 我实现的官方题解：动态规划
/**
执行用时：12 ms, 在所有 Go 提交中击败了66.50%的用户
内存消耗：6.3 MB, 在所有 Go 提交中击败了35.79%的用户
*/
func isMatch1(s string, p string) bool {
	m, n := len(s), len(p)
	dp := [][]bool{} // s[i], p[j]之前是否匹配
	for i := 0; i <= m; i++ {
		temp := make([]bool, n+1)
		dp = append(dp, temp)
	}
	// 初始化边界条件
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = true
		} else {
			break
		}
	}
	dp[0][0] = true
	// 迭代
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}




