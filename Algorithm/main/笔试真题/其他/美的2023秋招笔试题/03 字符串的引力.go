/**
一个字符串中字符的种类数目称之为 引力
求一个字符串的所以子字符串的引力之和
 */
package main

import "fmt"

//func main() {
//	var s string
//	fmt.Scan(&s)
//	n := len(s)
//	res := 0
//	for i:=0;i<n;i++ {
//		for j:=i+1;j<=n;j++ {
//			res += yili(s[i:j])
//		}
//	}
//	fmt.Println(res)
//}
//
//
//func yili (s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//	hash := make(map[byte]bool)
//	for i := range s {
//		hash[s[i]] = true
//	}
//	return len(hash)
//}

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	res := 0


	for i:=0;i<n;i++ {
		hash := make(map[byte]bool)
		for j:=i;j<n;j++ {
			hash[s[j]] = true
			res += len(hash)
		}
	}
	fmt.Println(res)
}

