/**
输入的字符串均为小写字符
好串：一个字符串中仅有一个字符出现次数为奇数次，其他都为偶数次
一个字符串的子串中有多少个是好串（注意是都少个，不是多少种）

 */
package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(fun3(s))
}
// 20%
func fun3(s string) int {
	res := 0
	n := len(s)
	for i:=0;i<n;i++ {
		hash := make(map[byte]bool) // true : 奇数
		hash[s[i]] = true
		res++

		for j:=i+2;j<n;j+=2 {

			hash[s[j]] = !hash[s[j]]
			hash[s[j-1]] = !hash[s[j-1]]

			//
			var count int
			for k := range hash{
				if hash[k] {
					count++
				}
			}
			if count == 1 {
				res++
			}
		}
	}
	return res
}
