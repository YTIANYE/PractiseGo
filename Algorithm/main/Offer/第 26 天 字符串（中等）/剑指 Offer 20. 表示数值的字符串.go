/*
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。

数值（按顺序）可以分成以下几个部分：

若干空格
一个 小数 或者 整数
（可选）一个 'e' 或 'E' ，后面跟着一个 整数
若干空格
小数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
下述格式之一：
至少一位数字，后面跟着一个点 '.'
至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
一个点 '.' ，后面跟着至少一位数字
整数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
至少一位数字
部分数值列举如下：

["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
部分非数值列举如下：

["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]
 

示例 1：

输入：s = "0"
输出：true
示例 2：

输入：s = "e"
输出：false
示例 3：

输入：s = "."
输出：false
示例 4：

输入：s = "    .1  "
输出：true
 

提示：

1 <= s.length <= 20
s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，空格 ' ' 或者点 '.' 。

*/

package main

import "fmt"

// 我的题解：暴力解法
/*
执行用时：4 ms, 在所有 Go 提交中击败了64.50%的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了93.81%的用
*/
// 去掉首尾空格
func removeSpace(s string) string {
	// 去除尾空格
	for j := len(s) - 1; len(s) != 0; j-- {
		if s[j] != ' ' {
			break
		}
		s = s[:j]
	}
	if len(s) == 0 {
		return ""
	}

	for s[0] == ' ' {
		if len(s) == 1 {
			return ""
		}
		s = s[1:]
	}
	return s
}

// 判断是不是一个数值
func isNumber(s string) bool {

	s = removeSpace(s)
	if len(s) == 0 {
		return false
	}

	// 科学计数
	for i, str := range s {
		if str == 'e' || str == 'E' {
			if i == len(s)-1 || i == 0 { // e处在字符串首尾
				return false
			}
			s1 := removeSymbol(s[:i])
			s2 := removeSymbol(s[i+1:])
			if len(s1) == 0 || len(s2) == 0 {
				return false
			}
			if (isDecimal(s1) || isInteger(s1)) && isInteger(s2) {
				return true
			}
			return false
		}
	}

	// 整数或小数
	s = removeSymbol(s)
	if len(s) == 0 {
		return false
	}
	if isDecimal(s) || isInteger(s) {
		return true
	}
	return false
}

// 判断是不是小数（不算符号）
func isDecimal(s string) bool {
	// 小数点
	for i, str := range s {
		if str == '.' {
			var s1, s2 string
			if i == 0 {
				s2 = s[i+1:]
			} else if i == len(s)-1 {
				s1 = s[:i]
			}
			s1 = s[:i]
			s2 = s[i+1:]
			if len(s1) == 0 && len(s2) == 0 {
				return false
			}
			if len(s1) == 0 && isInteger(s2) {
				return true
			}
			if len(s2) == 0 && isInteger(s1) {
				return true
			}
			if isInteger(s1) && isInteger(s2) {
				return true
			}
		}
	}
	return false
}

// 判断是不是整数（不算符号）
func isInteger(s string) bool {
	for _, str := range s {
		if str < 48 || str > 57 {
			return false
		}
	}
	return true
}

// 去除加减号
func removeSymbol(s string) string {
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	return s
}

// 我实现的精选题解：有限状态自动机
/*
执行用时：4 ms, 在所有 Go 提交中击败了64.50%的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了93.81%的用户
*/
func isNumber1(s string) bool {
	states := []map[uint32]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4}, //  0. start with 'blank'
		{'d': 2, '.': 4},                 //  1. 'sign' before 'e'
		{'d': 2, '.': 3, 'e': 5, ' ': 8}, //  2. 'digit' before 'dot'
		{'d': 3, 'e': 5, ' ': 8},         //  3. 'digit' after 'dot'
		{'d': 3},                         // 4. 'digit' after 'dot' (‘blank’ before 'dot')
		{'s': 6, 'd': 7},                 // 5. 'e'
		{'d': 7},                         // 6. 'sign' after 'e'
		{'d': 7, ' ': 8},                 // 7. 'digit' after 'e'
		{' ': 8},                         // 8. end with 'blank'
	}

	p := 0
	for _, c := range s{
		var t uint32
		if  c >= '0' && c <= '9'{
			t = 'd'
		}else if c == '-' || c == '+'{
			t = 's'
		}else if c == 'e' || c == 'E'{
			t = 'e'
		}else if c == '.' || c == ' '{
			t = uint32(c)
		}else{
			t = '?'
		}

		var ok bool
		p, ok = states[p][t]
		if ok == false{
			return false
		}
		// p = states[p][t]
	}
	if p == 2 || p == 3 || p == 7 || p == 8{
		return true
	}
	return false
}

func main() {

	// strs := []string{"+100", "5e2", "-123", "3.1416", "-1E-16", "0123"}
	// strs := []string{"12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"}
	// strs := []string{"0", "e", ".", "    .1"}
	strs := []string{" ", ". 1"}
	for _, s := range strs {
		fmt.Println(isNumber1(s))
	}


}
