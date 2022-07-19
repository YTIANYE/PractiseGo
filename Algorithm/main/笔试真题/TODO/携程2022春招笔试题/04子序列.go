/*

一个字符串只包含数字，其子串是9的倍数的子串有多少个（子串数值相同但是字符位置不同，为不同子串）

*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	hash := make(map[int]int) // key ： 数字  value:出现次数
	for i := range s {
		num, _ := strconv.Atoi(string(s[i]))
		hash[num]++
	}

	res := 0


	fmt.Println(res)

}

func isnine(num int) bool {
	if num%9 == 0 {
		return true
	}
	return false
}

func mod(num int) int {
	return num % 1000000007
}

// 字符串每一位相加是9的倍数，则这个数是9的倍数
/**
99
108
117
126
135
144
999
1008
1017
1998
1989
1980
*/
