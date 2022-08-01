/**
给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。

输入为 非空 字符串且只包含数字 1 和 0。



示例 1:

输入: a = "11", b = "10"
输出: "101"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"


提示：

每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。


*/

package main

import (
	"fmt"
	"strconv"
)

// 我的题解 :
/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.4 MB, 在所有 Go 提交中击败了47.00%的用户
 */


func addBinary(a string, b string) string {
	if len(a)< len(b) {
		a, b = b, a
	}
	m, n := len(a), len(b)
	// 两数相加
	res := ""
	jinwei := 0

	var jia func(num int)
	jia = func(num int) {
		if num >= 2 {
			jinwei = 1
			num -= 2
		} else {
			jinwei = 0
		}
		res = strconv.Itoa(num) + res
	}
	// 计算
	i, j := m-1, n-1
	for i >= 0 {
		var num int
		if j >= 0 {
			num = int(a[i]-'0'+b[j]-'0') + jinwei
			j--
		} else { // 处理长的数
			num = int(a[i]-'0') + jinwei
		}
		i--
		jia(num)
	}
	// 处理进位
	if jinwei == 1 {
		res = "1" + res
	}
	return res

}

// 官方题解：优化了 func jia(num int)
func addBinary1(a string, b string) string {
	ans := ""
	carry := 0 // 进位
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		// 分开计算两个数
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		// 进位的处理比我的题解更简单
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := "11"
	b := "10"
	fmt.Println(addBinary(a, b))
}
