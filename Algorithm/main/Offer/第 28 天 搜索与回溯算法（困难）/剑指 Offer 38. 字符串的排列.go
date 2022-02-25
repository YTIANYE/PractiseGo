/*
输入一个字符串，打印出该字符串中字符的所有排列。

 

你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

 

示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]
 

限制：

1 <= s 的长度 <= 8

*/
package main

import "fmt"

// 我的题解：暴力法：每次添加字符去重
/*
执行用时：1320 ms, 在所有 Go 提交中击败了5.64%的用户
内存消耗：6.9 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func permutation1(s string) []string {
	if len(s) == 0 || len(s) == 1 {
		return []string{s}
	}

	res := []string{}
	for i := range s {
		res = fun(res, string(s[i]))
	}
	return res
}

func fun(strs []string, str string) []string {
	if len(strs) == 0 {
		return []string{str}
	}
	res := []string{}
	for i := range strs {
		for j := 0; j < len(strs[i])+1; j++ {
			res = append(res, strs[i][:j]+str+strs[i][j:])
		}
	}
	// 去重
	for i := 0; i < len(res)-1; i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] == res[j] {
				if j == len(res)-1 {
					res = res[:j]
				} else {
					res = append(res[:j], res[j+1:]...)
					j-- // 注意index的变化
				}
			}
		}
	}

	return res
}

// 我实现的精选题解：溯回+剪枝
/*
执行用时：36 ms, 在所有 Go 提交中击败了74.09%的用户
内存消耗：7.5 MB, 在所有 Go 提交中击败了72.71%的用户
*/
func permutation(s string) []string {
	res := []string{}
	ss := make([]uint8, len(s)) // 创建副本
	copy(ss, s)

	var def func(int)
	def = func(x int) {
		if x == len(ss)-1 {
			res = append(res, string(ss))
			return
		}
		hashMap := make(map[uint8]bool)
		for i := x; i < len(ss); i++ {
			if hashMap[ss[i]] {
				continue
			}
			hashMap[ss[i]] = true
			ss[x], ss[i] = ss[i], ss[x]
			def(x + 1)
			ss[x], ss[i] = ss[i], ss[x]
		}
	}
	def(0)
	return res
}

func main() {
	// s := "abcdefgh"
	// s := "aab"
	// s := "kzfxxx"
	s := "kzxx"
	fmt.Println(permutation(s))
}
