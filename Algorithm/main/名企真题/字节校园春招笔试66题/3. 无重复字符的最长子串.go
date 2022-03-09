/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

 

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 

提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

*/

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	} else if l == 1 {
		return 1
	}

	max := 0
	tempStr := ""
	for i := 0; i < l; i++ {
		for j:=0;j<len(tempStr);j++ {
			if s[i] == tempStr[j] {
				tempStr = tempStr[j+1:]
			}
		}
		tempStr = tempStr + string(s[i])
		if len(tempStr) > max {
			max = len(tempStr)
		}
	}
	return max
}

func main() {
	s := "abcabcbb"
	// s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
