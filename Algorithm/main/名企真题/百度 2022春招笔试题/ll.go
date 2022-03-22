/*
给定一个字符串，全部由0和1构成、
一个字符不大于相邻的字符，
则， 可以删掉字符串中的该字符得到一个字串
求全部字串的数目

*/

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	resHash := make(map[string]int)
	lastmap := make(map[string]int)
	lastmap[s]++

	for len(lastmap) != 0 { // hashmap 中的长度
		hashmap := make(map[string]int)
		for skey := range lastmap{
			s = skey
			for i := 0; i < len(s); i++ {
				if (i-1 >= 0 && s[i] <= s[i-1]) || (i+1 < len(s) && s[i] <= s[i+1]) {
					temps := s[:i] + s[i+1:]
					if _, ok := hashmap[temps]; !ok {
						hashmap[temps]++
					}
				}
			}
		}

		lastmap = make(map[string]int)
		for key, val := range hashmap{
			resHash[key] = val
			lastmap[key] = val
		}
	}

	fmt.Println(resHash)
	fmt.Println(len(resHash))
}
