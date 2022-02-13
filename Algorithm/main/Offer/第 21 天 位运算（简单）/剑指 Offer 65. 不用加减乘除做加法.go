/*
写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
示例:
输入: a = 1, b = 1
输出: 2
提示：
a, b 均可能是负数或 0
结果不会溢出 32 位整数
*/

package main

import (
	"fmt"
)

// 暴力解法：超时
func add1(a int, b int) int {
	if a > 0 {
		for a != 0 {
			b++
			a--
		}
	} else {
		for a != 0 {
			b--
			a++
		}
	}

	return b
}

// 精选题解：
// n=a⊕b
// c=a&b<<1
// 非进位和：异或运算
// 进位：与运算+左移一位
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.8 MB, 在所有 Go 提交中击败了92.70%的用户
*/

func add(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1 // c = 进位
		a ^= b            // a = 非进位和
		b = c             // b = 进位
	}
	return a
}

func main() {
	a := -215049105
	b := -559615654
	// a := 1
	// b := 2
	fmt.Println(add(a, b))
}
