/*请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：

s.length <= 40000*/
package main

/*我的题解： 动态规划 + 线性遍历*/
// 执行用时：16 ms, 在所有 Go 提交中击败了13.93%的用户
// 内存消耗：6.1 MB, 在所有 Go 提交中击败了6.26%的用户
func lengthOfLongestSubstring1(s string) int {
	result := 0
	if len(s) == 0 {
		return 0
	}
	curstr := ""
	for _, str := range s {
		index := instr(string(str), curstr)
		if index == -1 {
			curstr += string(str)
		} else {
			curstr = curstr[index+1:] + string(str)
		}

		if result < len(curstr) {
			result = len(curstr)
		}
	}
	return result
}

// 判断一个字符str是不是在字符串curstr中
// 不在返回 -1，在返回索引
func instr(str, curstr string) int {
	for i, s := range curstr {
		if string(s) == str {
			return i
		}
	}
	return -1
}

/*我实现的精选题解：动态规划+ 哈希表*/
// 执行用时：12 ms, 在所有 Go 提交中击败了38.49%的用户
// 内存消耗：4.1 MB, 在所有 Go 提交中击败了10.64%的用户
func lengthOfLongestSubstring2(s string) int {
	hashmap := make(map[string]int)
	result, tmp := 0, 0
	for j, str := range s {
		i, ok := hashmap[string(str)]
		if !ok {
			i = -1
		}
		hashmap[string(str)] = j

		if j-i > tmp {
			tmp += 1
		} else {
			tmp = j - i
		}

		if tmp > result {
			result = tmp
		}
	}
	return result
}

/*我实现的精选题解： 双指针 + 哈希表*/
// 每次在哈希表中记录字符出现的索引，用新索引替代旧索引，两者差作为之间字符串的长度

// 执行用时：4 ms, 在所有 Go 提交中击败了94.41%的用户
// 内存消耗：4.1 MB, 在所有 Go 提交中击败了10.64%的用户
func lengthOfLongestSubstring(s string) int {
	strmap := make(map[string]int)
	result := 0
	i := -1
	for j, str := range s {
		index, ok := strmap[string(str)]
		if ok && i < index { // 遇到连续多个相同的字符时，应及时更新i,保持在最后面的那个字符上
			i = index
		}
		strmap[string(str)] = j
		if j-i > result {
			result = j - i
		}
	}
	return result
}
