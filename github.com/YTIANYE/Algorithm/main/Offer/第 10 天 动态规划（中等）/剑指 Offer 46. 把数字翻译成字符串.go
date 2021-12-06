/*
给定一个数字，我们按照如下规则把它翻译为字符串：
0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

*/
package main

import (
	"fmt"
	"strconv"
)

/*我的题解*/

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了79.27%的用户
func translateNum1(num int) int {
	if num >= 0 && num <= 9 {
		return 1
	}
	ppre := 1          // 截至到上上一个数，种类个数
	pre := 1           // 截止到上一个数，种类个数
	cur := 0           // 到当前数，种类的个数
	prenum := num % 10 // 当前数的前一个数
	num /= 10

	for num != 0 {
		curnum := num % 10 // 当前数
		num /= 10
		if isparam(curnum*10 + prenum) {
			cur = ppre + pre
		} else {
			cur = pre
		}
		ppre = pre
		pre = cur
		prenum = curnum
	}
	return pre

}

func isparam(num int) bool {
	if num >= 10 && num <= 25 {
		return true
	} else {
		return false
	}
}

/*我实现的官方题解:转换成字符串后正向计算*/

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func translateNum(num int) int {
	str := strconv.Itoa(num) // int转换成字符串
	pp, p, cur := 0, 0, 1
	for i, _ := range str {
		pp, p, cur = p, cur, 0
		cur += p
		if i == 0 {
			continue
		}
		x, _ := strconv.ParseInt(str[i-1:i+1], 10, 64) // string 转 int
		if x >= 10 && x <= 25 {
			cur += pp
		}
	}
	return cur
}

func main() {
	num := 12258
	fmt.Println(translateNum(num))
}
