/*
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成

*/

package main

import "fmt"

// 我的题解：中心扩散法
/*
执行用时：8 ms, 在所有 Go 提交中击败了68.08%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了100.00%的用户
*/

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	res := ""
	for i := 0; i < len(s); i++ {
		res = huiwen(s, res, i, i)
		if i+1 <= len(s)-1 {
			res = huiwen(s, res, i, i+1)
		}
	}

	return res
}

// 以一个或两个字符向外扩散，查找是否是回文字符串
func huiwen(s, res string, left, right int) string {
	for left >= 0 && right <= len(s)-1 {
		if s[left] != s[right] {
			break
		}
		temp := s[left : right+1]
		if len(temp) > len(res) {
			res = temp
		}
		left--
		right++
	}
	return res
}

func main() {
	s := "ac"
	fmt.Println(longestPalindrome(s))
}
