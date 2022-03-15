/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
 

示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
示例 2：

输入：s = "a", t = "a"
输出："a"
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

提示：

1 <= s.length, t.length <= 105
s 和 t 由英文字母组成

进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？

*/

package main

import (
	"fmt"
	"math"
)

// 我的题解：暴力超时
func minWindow1(s string, t string) string {
	hasht := map[byte]int{}
	for i := range t {
		hasht[t[i]]++
	}

	temp := ""       // 记录重要字符
	index := []int{} // 记录重要字符在s中的位置
	for i := range s {
		if hasht[s[i]] != 0 {
			temp += string(s[i])
			index = append(index, i)
		}
	}

	res := ""

	for i := range temp {
		temphasht := make(map[byte]int)
		for k, v := range hasht {
			temphasht[k] = v
		}

		start := i
		end := len(temp)
		for j := i; j < len(temp); j++ {
			if temphasht[temp[j]] > 0 {
				temphasht[temp[j]]--
			}
			if temphasht[temp[j]] == 0 {
				delete(temphasht, temp[j])
			}
			if len(temphasht) == 0 {
				end = j
				break
			}
		}

		if end < len(temp) {
			tempstr := s[index[start] : index[end]+1]
			if len(res) == 0 || len(tempstr) < len(res) { // 注意判断条件
				res = tempstr
			}
		}
	}
	return res
}

// 我实现的官方题解：滑动窗口
/*
执行用时：168 ms, 在所有 Go 提交中击败了5.50%的用户
内存消耗：2.7 MB, 在所有 Go 提交中击败了80.14%的用户
*/
func minWindow2(s string, t string) string {
	hashT := make(map[byte]int)
	hashCurr := make(map[byte]int)
	for i := range t {
		hashT[t[i]]++
	}

	// 判断hashCurr是否完全包好hashT
	var check func() bool
	check = func() bool {
		for i := range hashT {
			if hashCurr[i] < hashT[i] {
				return false
			}
		}
		return true
	}

	currLen := math.MaxInt32
	left, right := -1, -1
	for p, q := 0, 0; p < len(s); p++ { // p右指针 q 左指针
		if hashT[s[p]] > 0 {
			hashCurr[s[p]]++
		}
		for check() && q <= p {
			if p-q+1 < currLen {
				currLen = p - q + 1
				left, right = q, p
			}
			if _, ok := hashCurr[s[q]]; ok {
				hashCurr[s[q]]--
			}
			q++
		}
	}

	if left == -1 {
		return ""
	}
	return s[left : right+1]
}

// 我的题解：优化滑动窗口
/*
执行用时：160 ms, 在所有 Go 提交中击败了5.50%的用户
内存消耗：6.9 MB, 在所有 Go 提交中击败了5.12%的用户
*/
func minWindow(s string, t string) string {
	hashT := make(map[byte]int)
	hashCurr := make(map[byte]int)
	for i := range t {
		hashT[t[i]]++
	}

	var check func() bool
	check = func() bool {
		for key := range hashT {
			if hashCurr[key] < hashT[key] {
				return false
			}
		}
		return true
	}

	// 建立一个只包含t中的元素的数组和对应索引
	strs := []byte{}
	indexs := []int{}
	for i := range s {
		if _, ok := hashT[s[i]]; ok {
			strs = append(strs, s[i])
			indexs = append(indexs, i)
		}
	}

	currLen := math.MaxInt32
	resLeft, resRight := -1, -1
	for right, left := 0, 0; right < len(strs); right++ {
		hashCurr[strs[right]]++
		for check() && left <= right {
			tempLen := indexs[right] - indexs[left] + 1
			if tempLen < currLen {
				currLen = tempLen
				resLeft = indexs[left]
				resRight = indexs[right]
			}
			hashCurr[strs[left]]--
			left++
		}
	}

	if resLeft == -1 {
		return ""
	}
	return s[resLeft : resRight+1]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
