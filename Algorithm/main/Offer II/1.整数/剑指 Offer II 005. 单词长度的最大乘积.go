/**
给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。



示例 1:

输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"]
输出: 16
解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
示例 2:

输入: words = ["a","ab","abc","d","cd","bcd","abcd"]
输出: 4
解释: 这两个单词为 "ab", "cd"。
示例 3:

输入: words = ["a","aa","aaa","aaaa"]
输出: 0
解释: 不存在这样的两个单词。


提示：

2 <= words.length <= 1000
1 <= words[i].length <= 1000
words[i] 仅包含小写字母

*/

package main

// 我的题解：二进制数作为标记
/**
执行用时：116 ms, 在所有 Go 提交中击败了22.72%的用户
内存消耗：6.6 MB, 在所有 Go 提交中击败了7.80%的用户
*/

func maxProduct(words []string) int {
	hashmap := make(map[string]uint32)
	for _, word := range words {
		hashmap[word] = toUint32(word)
	}

	res := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if !hasSameWord(hashmap[words[i]], hashmap[words[j]]) {
				res = max1(res, len(words[i])*len(words[j]))
			}
		}
	}
	return res
}

func toUint32(str string) uint32 {
	var res uint32
	for _, char := range str {
		t := int(char - 'a')
		res = res | (1 << t)
	}
	return res
}

func hasSameWord(a, b uint32) bool {
	if a&b == 0 {
		return false
	}
	return true
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 官方题解：位运算优化

func maxProduct1(words []string) (ans int) {
	masks := map[int]int{} //使用哈希表记录每个位掩码对应的最大单词长度
	for _, word := range words {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		if len(word) > masks[mask] {
			masks[mask] = len(word)
		}
	}

	for x, lenX := range masks {
		for y, lenY := range masks {
			if x&y == 0 && lenX*lenY > ans {
				ans = lenX * lenY
			}
		}
	}
	return
}
