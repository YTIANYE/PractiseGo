/**
O（1）空间复杂度反转列表中的元音字母, 返回反转后的字符串

例如：
输入: aeiou
输出：uoiea

输入: hellio
输出：hollie

输入: aecdui
输出：iucdea
*/
package main

import "fmt"

func rev(s string) string {
	hash := make(map[byte]bool)
	hash['a'] = true
	hash['e'] = true
	hash['i'] = true
	hash['o'] = true
	hash['u'] = true
	n := len(s)
	// 取代
	fun := func(i int, str string) {
		if i== 0 {
			if n != 1 {
				s = str + s[i+1:]
			}else {
				s = str
			}
		}else if i == n-1 {
			s = s[:i] + str
		}else {
			temp :=  str + s[i+1:]
			s = s[:i] +temp
		}
	}
	//  换位置
	tihuan := func(i, j int) {
		a := string(s[i])
		b := string(s[j])
		fun(i, b)
		fun(j, a)
	}
	//

	i:= 0
	j := n-1
	for i<j {
		for i< j && hash[s[j]] == false {
			j--
		}
		for i<j && hash[s[i]] == false {
			i++
		}
		if i < j {
			tihuan(i, j)
			j--
			i++
		}
	}
	return s
}

func main() {
	s1 := "aeiou"
	res1 := rev(s1)
	fmt.Println(res1)

	s2 := "hellio"
	res2 := rev(s2)
	fmt.Println(res2)

	s3 := "aecdui"
	res3 := rev(s3)
	fmt.Println(res3)
}


