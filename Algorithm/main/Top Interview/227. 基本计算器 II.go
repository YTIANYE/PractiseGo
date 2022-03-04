/*
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
整数除法仅保留整数部分。

示例 1：
输入：s = "3+2*2"
输出：7
示例 2：
输入：s = " 3/2 "
输出：1
示例 3：
输入：s = " 3+5 / 2 "
输出：5

提示：

1 <= s.length <= 3 * 105
s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
s 表示一个 有效表达式
表达式中的所有整数都是非负整数，且在范围 [0, 231 - 1] 内
题目数据保证答案是一个 32-bit 整数

*/

package main

import (
	"fmt"
)

// 我实现的官方题解：
/*
执行用时：4 ms, 在所有 Go 提交中击败了81.02%的用户
内存消耗：5.9 MB, 在所有 Go 提交中击败了87.70%的用户
*/
func calculate(s string) int {

	stack := []int{}
	num := 0
	pretable := '+'

	for i, char := range s {
		isNum := char >= '0' && char <= '9'
		// 先判断是不是一个数
		if isNum {
			num = num*10 + int(char-'0')
		}

		// 后判断是符号的情况、或是最后一个数
		if !isNum && char != ' ' || i == len(s)-1 { // 注意最后一个数的判断
			switch pretable {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num) // 减法这样处理很好
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			pretable = char
			num = 0
		}
	}

	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}

func main() {
	s := "1  +1"
	fmt.Println(calculate(s))
}
