/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
 

示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
示例 4：

输入：s = "([)]"
输出：false
示例 5：

输入：s = "{[]}"
输出：true
 

提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成

*/

package main

// 我的题解：
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了60.76%的用户
*/

func isValid(s string) bool {
	stack := []int32{}
	p := -1
	hashMap := map[int32]int32{')': '(', ']': '[', '}': '{'}

	for _, char := range s{
		if p >= 0 && stack[p] == hashMap[char]{
			stack = stack[:p]
			p--
		}else{
			p++
			stack = append(stack, char)
		}
	}

	if p >=0{
		return false
	}
	return true
}