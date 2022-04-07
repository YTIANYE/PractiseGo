/**
从一个字符串中找到第一个不重复的字符，返回它的索引。如果不存在，则返回 -1。
案例:
s = "qwert"
返回 0.
s = "testengine",
返回 2.
s = "helloolleh",
返回 -1.
s = "abababe",
返回 6.
注意：字符串中只会包含小写字母。


 */
package main

import "fmt"

func firstUniqChar(s string) int {
	hashString := make(map[uint8]int)// 次数
	hashIndex := make(map[uint8]int)// 索引

	for i  := range s{
		hashString[s[i]]++
		hashIndex[s[i]] = i
	}

	for i := range s{
		if hashString[s[i]] == 1{
			return hashIndex[s[i]]
		}
	}
	return -1
}

func main(){
	strs := []string{"qwert", "testengine", "helloolleh" , "abababe"}
	for _, s := range strs{
		fmt.Println(firstUniqChar(s))
	}

}

