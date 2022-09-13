/**
给出一个字符串只有0 和 1
每次操作是反转相邻的两个字符
0变1 1变0
问这个字符串能不能变成所有字符都相同的情况


比如
101
可以变成011
在变成000
 */
package main

import "fmt"

func main() {
	var q int
	fmt.Scan(&q)
	for ; q != 0; q-- {
		var s string
		fmt.Scan(&s)
		if panduan(s) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

// 100%
func panduan(s string) bool {
	// n := len(s)
	num0 := 0
	num1 := 0
	for i := range s {
		if s[i] == '0' {
			num0++
		} else {
			num1++
		}
	}
	if num0%2 == 1 && num1%2 == 1 {
		return false
	}
	return true
}
