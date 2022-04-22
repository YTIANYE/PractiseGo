/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。

 

示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]
 

提示：

1 <= s.length <= 16
s 仅由小写英文字母组成

*/

package main

import "fmt"

// 官方题解：动态规划+回溯
/**
执行用时：208 ms, 在所有 Go 提交中击败了99.91%的用户
内存消耗：23 MB, 在所有 Go 提交中击败了82.98%的用户
*/
func partition(s string) [][]string {
	// 初始化
	n := len(s)
	res := [][]string{}

	ishuiwen := make([][]bool, n)
	for i := 0; i < n; i++ {
		ishuiwen[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			ishuiwen[i][j] = true
		}
	}

	// 动态规划
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			ishuiwen[i][j] = s[i] == s[j] && ishuiwen[i+1][j-1]
		}
	}

	// 递归
	queue := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]string{}, queue...))
			return
		}
		for j := i; j < n; j++ {
			if ishuiwen[i][j] {
				queue = append(queue, s[i:j+1])
				dfs(j + 1)
				queue = queue[:len(queue)-1]
			}
		}
	}
	dfs(0)

	return res
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
