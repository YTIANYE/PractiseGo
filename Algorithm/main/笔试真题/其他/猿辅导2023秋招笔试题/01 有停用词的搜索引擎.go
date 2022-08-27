/**
统计一组单词中出现频率最高的单词，忽略大小写
注意不用统计停用词，停用词中‘?’表示一个字符

输入 M ，表示组数
每组中，第一行输入单词个数和相应的单词
第二行输入停用词个数及相应停用词


2
12 A tidy tiger tied a tie tighter to tidy her tiny tail
1 a
16 A big black bug bit a big black bear made the big black bear bleed blood
2 b?? b???
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 我的题解： 50%
func isRight(s string, ting string) bool { // 匹配，返回true
	if len(s) != len(ting) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if ting[i] == '?' {
			continue
		}
		if s[i] != ting[i] {
			return false
		}
	}
	return true
}

func fun(s1, s2 string) {
	hashWords := make(map[string]int)
	hashTings := make(map[string]int)

	words := strings.Split(s1, " ")
	words = words[1:]
	for i := range words {
		words[i] = strings.ToLower(words[i])
		hashWords[words[i]]++
	}

	tings := strings.Split(s2, " ")
	tings = tings[1:]
	for i := range tings {
		tings[i] = strings.ToLower(tings[i])
		hashTings[tings[i]]++
	}

	res := 0
	for s, num := range hashWords {
		// 判断是否是停用词
		tag := false
		for ting := range hashTings {
			if isRight(s, ting) {
				tag = true // 是停用词
				break
			}
		}
		if tag {
			continue
		}
		//
		res = max(res, num)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var M int
	fmt.Scan(&M)
	input := bufio.NewScanner(os.Stdin)
	var words, tings string
	for i := 0; i < M*2; i++ {
		input.Scan()
		if i%2 == 0 {
			words = input.Text()
		} else {
			tings = input.Text()
			fun(words, tings)
		}
	}
}
