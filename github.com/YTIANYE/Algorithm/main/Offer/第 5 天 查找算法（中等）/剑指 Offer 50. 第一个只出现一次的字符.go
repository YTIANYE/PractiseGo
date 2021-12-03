package main

import "fmt"

/*
在字符串 s 中找出第一个只出现一次的字符。
如果没有，返回一个单空格。 s 只包含小写字母。

示例 1:
输入：s = "abaccdeff"
输出：'b'

示例 2:
输入：s = ""
输出：' '

限制：
0 <= s 的长度 <= 50000

*/

/*我的题解：暴力法*/

// 执行用时：64 ms, 在所有 Go 提交中击败了6.09%的用户
// 内存消耗：5.4 MB, 在所有 Go 提交中击败了41.48%的用户

func firstUniqChar1(s string) byte {
	charmap := make(map[byte]int) // 单个字符类型 int32
	for _, ch := range s {
		// _, ok := charmap[byte(ch)]
		// if ok {
		// 	charmap[byte(ch)]++
		// } else {
		// 	charmap[byte(ch)] = 1
		// }

		// 直接 ++ 即可，默认从0开始+1
		charmap[byte(ch)]++
	}

	for _, ch := range s {
		if charmap[byte(ch)] == 1 {
			return byte(ch)
		}
	}
	return ' '
}

/*我实现的官方题解：哈希表存储索引*/

// 执行用时：28 ms, 在所有 Go 提交中击败了47.02%的用户
// 内存消耗：5.4 MB, 在所有 Go 提交中击败了24.12%的用户
func firstUniqChar(s string) byte {
	charsmap := make(map[int32]int)
	for i, ch := range s {

		if charsmap[ch] == 0 {
			charsmap[ch] = i + 1 // 避免s中第一个本身索引就是0
		} else {
			charsmap[ch] = -1
		}
	}
	result := len(s)
	for _, value := range charsmap { // map遍历方法
		if value != -1 && value-1 < result {
			result = value - 1
		}
	}
	if result == len(s) {
		return ' '
	}
	return s[result]
}



func main() {
	s := "abcacdb"
	fmt.Println(firstUniqChar(s))
}
