/**
2.给定一个字符串 s（小写英文字母） 和一些长度相同的单词 words 。
找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置
注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序

输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案

示例 2：
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]

示例 3：
输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出：[6,9,12]

*/

package main

import "fmt"

func main() {
	//s := "barfoothefoobarman"
	//words := []string{"foo", "bar"}
	//
	//s := "wordgoodgoodgoodbestword"
	//words := []string{"foo", "bar"}

	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}
	res := []int{}

	n := len(s)
	m := len(words)
	k := len(words[0])
	for i := 0; i < n-k*m; i++ {
		hashmap := make(map[string]bool)
		for _, word := range words {
			hashmap[word] = false // 不存在
		}
		for j := i; j < i+k*m; j += k {
			str := s[j : j+k]
			if _, ok := hashmap[str]; ok {
				hashmap[str] = true // 存在
			}
		}
		// 判断words中的Word是否都存在在该子串中
		tag := true
		for str := range hashmap {
			if hashmap[str] == false {
				tag = false
				break
			}
		}
		if tag {
			res = append(res, i)
		}
	}

	//
	fmt.Println(res)

}
