/*
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。
例如输入字符串"I am a student. "，则输出"student. a am I"。



示例 1：

输入: "the sky is blue"
输出:"blue is sky the"
示例 2：

输入: " hello world! "
输出:"world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good example"
输出:"example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


说明：
无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

*/

package main

import "fmt"

/*我的题解：暴力题解*/
/*执行用时：8 ms, 在所有 Go 提交中击败了27.54%的用户
内存消耗：7.1 MB, 在所有 Go 提交中击败了37.83%的用户*/
func reverseWords1(s string) string {

	var strs string
	tag := false // 还没确定j的位置
	j := 0
	i := 0

	for ; i < len(s); i++ {
		if s[i] == ' ' {
			if tag { // 截取一个单词
				strs = s[j:i] + " " + strs
				tag = false
			}
		} else {      // s[i] != ' '
			if !tag { // 还没确定j的位置
				j = i
				tag = true
				if i == len(s)-1 { // s = "a"的情况
					strs = s[j:len(s)] + " " + strs
				}
			} else if i == len(s)-1 { // 确定了j的位置，i到达s最后一个字符
				strs = s[j:len(s)] + " " + strs
			}
		}
	}

	if len(strs) > 0 {
		return strs[:len(strs)-1]
	} else {
		return ""
	}

}

/*我实现的精选题解*/
func reverseWords(s string) string {
	// 删除首尾空格
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			s = s[i:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			s = s[:i+1]
			break
		}
	}
	// 截取字符
	strs := ""
	i, j := len(s)-1, len(s)-1
	for i >= 0 {
		for i >= 0 && s[i] != ' ' {
			i--
		}
		strs += s[i+1:j+1] + " "
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
	}
	// 返回结果
	if len(strs) > 0 {
		return strs[:len(strs)-1]
	} else {
		return ""
	}
}

func main(){
	// s := "a good   example"
	// s := "a 1"
	// s := "  hello world!  "
	s := "s"
	fmt.Println(reverseWords(s))
}
