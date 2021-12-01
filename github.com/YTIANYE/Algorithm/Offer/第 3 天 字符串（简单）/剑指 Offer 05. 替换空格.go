package main

import "fmt"

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."

限制：
0 <= s 的长度 <= 10000
*/

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.3 MB, 在所有 Go 提交中击败了22.91%的用户

func replaceSpace(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' '{
			s = s[:i] + "%20" + s[i+1:]
		}
	}
	return s
}

func main() {
	s := "hello world !"
	fmt.Println(replaceSpace(s))
}
