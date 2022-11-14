/**
给定一个字符串 s ， 找出其中不含有重复字符的 最长字串 的长度

示例1：
输入 s = "abcabcbb"
输出 3
解释：abc

示例2 ：
输入 s = "bbbbbb"
输出 1
解释： 因为无重复字符的最长字串是“b”, 所以其长度为 1

示例3：
输入 s = "pwwkew"
输出 3
解释：因为无重复字符的最长字串是“wke”, 所以其长度是3

// 注意：答案必须是 子串 的长度，“pwke”是一个子序列，不是子串
*/
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	//
	p, q := 0, 0 // 左闭右闭
	hash := make(map[byte]int)
	res := 0 // 长度
	for p < n {
		if i, ok := hash[s[p]]; ok {
			for t := q; t < i+1; t++ {
				delete(hash, s[t])
			}
			q = i + 1
		}
		hash[s[p]] = p
		p++
		res = max(res, p-q)
	}
	//
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
a b c a d
  q     p
*/

func main() {
	strs := []string{
		"abcabcbb",
		"bbbbbb",
		"pwwkew",
	}
	for _, s := range strs {
		fmt.Println(lengthOfLongestSubstring(s))
	}
}
