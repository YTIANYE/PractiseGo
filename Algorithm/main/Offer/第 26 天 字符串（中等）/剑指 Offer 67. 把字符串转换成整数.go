/*
写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。

当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：

假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。

示例1:

输入: "42"
输出: 42
示例2:

输入: "   -42"
输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。
    我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
示例3:

输入: "4193 with words"
输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
示例4:

输入: "words and 987"
输出: 0
解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
     因此无法执行有效的转换。
示例5:

输入: "-91283472332"
输出: -2147483648
解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
因此返回 INT_MIN (−231) 。

*/

package main

import (
	"fmt"
	"math"
)

// 我的题解：直观解法
/*
执行用时：4 ms, 在所有 Go 提交中击败了50.42%的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了5.04%的用户
*/

func strToInt(str string) int32 {
	INT_MAX := int32(math.Pow(2, 31) - 1)
	INT_MIN := int32(-1 * math.Pow(2, 31))

	if len(str) == 0 {
		return 0
	}
	// 去除空格
	str = rmSpace(str)
	if len(str) == 0 {
		return 0
	}
	// 判断符号
	var sym int32
	sym = 1
	if str[0] == '-' || str[0] == '+' {
		if str[0] == '-' {
			sym = -1
		}
		if len(str) == 1 {
			return 0
		}
		str = str[1:]
	} else if !isInt(str[0]) {
		return 0
	}
	// 判断连续数字
	states := map[uint8]int32{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	var res int32
	res = 0
	for i := range str {
		if isInt(str[i]) {
			if res*sym >= INT_MAX/10 && str[i] >8 {
				return INT_MAX
			}
			if res*sym <= INT_MIN/10 && str[i] > 7 {
				return INT_MIN
			}
			res = res*10 + states[str[i]]
		} else {
			break
		}
	}
	// 返回

	return res * sym
}

// 去除开头空格
func rmSpace(str string) string {

	for str[0] == ' ' {
		if len(str) == 1 {
			return ""
		}
		str = str[1:]
	}
	return str
}

// 判断数字
func isInt(num uint8) bool {
	if num >= '0' && num <= '9' {
		return true
	}
	return false
}

func main() {
	//nums := []string{"42", "   -42", "4193 with words", "words and 987", "-91283472332", "  -0012a42"}
	nums := []string{"9223372036854775808"}
	for _, num := range nums {
		fmt.Println(strToInt(num))
	}
}
