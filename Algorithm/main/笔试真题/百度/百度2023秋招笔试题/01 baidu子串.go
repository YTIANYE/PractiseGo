/**
统计一个字符串的子字符串是“baidu”类型的字符串的数量
“baidu”字符串即，第1 4 是辅音， 2 3 5 是元音
 */
package main

import "fmt"
// 100%
func main() {
	var s string
	fmt.Scan(&s)
	//
	fmt.Println(fun(s))

}

func fun(s string) int {
	res := 0
	n := len(s)
	if n <5 {
		return res
	}
	//
	hash := make(map[byte]bool)
	hash['a'] = true
	hash['e'] = true
	hash['i'] = true
	hash['o'] = true
	hash['u'] = true
	//
	isright := func(str string) bool  {
		s1, s2, s3, s4, s5 := str[0], str[1], str[2], str[3], str[4]
		thash := make(map[byte]int)
		for i:= range str {
			thash[str[i]]++
		}
		if len(thash) != 5 {
			return false
		}
		if !hash[s1] && !hash[s4] && hash[s2] && hash[s3] && hash[s5] {
			return true
		}
		return false
	}
	//
	for i:=4;i<n;i++ {
		j := i - 4
		 if isright(s[j:i+1]) {
			 res++
		 }
	}
	return res
}