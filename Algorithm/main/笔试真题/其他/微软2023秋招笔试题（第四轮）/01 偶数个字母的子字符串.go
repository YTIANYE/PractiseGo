/**
S是长度为N的由小写字母构成的字符串
找出S字符串的子字符串，
使得子字符串中每个字符出现次数是偶数，
求这样的最长的子字符串的长度

bbaaadadb
aa,adad, daaada, aaadad
6

abacb
0

zthtzh
6
*/
package main

func Solution1(S string) int {
	// write your code in Go (Go 1.4)

	// tag是一个整数， 0-25 位置上，有一个位是 1 表示 该位上的字母个数为奇数
	// 最低位（第0位）表示'a'的个数是不是偶数，第25位表示'z'的个数是不是偶数，
	// 0 表示是偶数个，1表示是奇数个
	tag := 0
	n := len(S)

	var change func(char byte) // 新增一个或减少一个字符char,其所在的位变号
	change = func(char byte) {
		index := char - 'a'
		temp := 1 << index
		tag = tag ^ temp // 借助异或，0^1 = 1； 1^1 = 0；
	}

	res := 0
	for i := 0; i < n; i++ {
		tag = 0
		change(S[i])
		for j := i + 1; j < n; j++ {
			change(S[j])   // 新加入字符
			l := j - i + 1 //[i,j]左闭右闭
			if l%2 == 1 {  // 注意change()之后在判断
				continue
			}
			if tag == 0 {
				res = max1(res, l)
			}
		}
	}
	return res
}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
