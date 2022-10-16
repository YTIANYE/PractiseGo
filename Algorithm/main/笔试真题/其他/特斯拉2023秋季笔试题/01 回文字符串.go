/**
删除字符串中的一些字符
使其是一个不含连续重复字符的字符串

输入：
一个字符串
修改每个字符对应的代价
 */

package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, C []int) int {
	// write your code in Go 1.13.8
	res := 0
	n := len(S)
	for i:=1;i<n;i++ {
		if S[i] == S[i-1] {
			res += min(C[i], C[i-1])
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
