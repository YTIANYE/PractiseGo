/**
字符串匹配
str,
mod中‘？’表示某个字符，从str最后一个倒着配，发挥能够匹配mod的str中的最后一个子字符串的开头字符的索引,
 */

package main

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * findStr从反方向查找子串，返回最后一次匹配成功时，子串的偏移位置，如果没有匹配成功返回-1
 * @param str string字符串 待匹配的字符串
 * @param mod string字符串 匹配模式字符串
 * @return int整型
 */

// 100%
func findStr(str string, mod string) int {
	// write code here
	n := len(str) // 11
	m := len(mod) // 2
	res := -1
	for j := n - 1; j+1-m >= 0; j-- {
		if isSame(str[j-m+1:j+1], mod) {
			res = j
		}
	}
	return res
}

func isSame(str, mod string) bool {
	for i := 0; i < len(mod); i++ {
		if mod[i] == '*' {
			continue
		}
		if mod[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
