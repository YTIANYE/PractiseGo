/**
题目：将一个数的每一位进行重新排列，如果判断可不可以组成偶数
可以和原来的数相同，不能有前导0
 */

package main

import (
	"fmt"
	"strconv"
)

// 100%
func main() {
	var n int
	fmt.Scan(&n)
	for ;n!=0;n-- {
		var s string
		fmt.Scan(&s)
		fmt.Println(fun(s))
	}
}

func fun(s string) int {
	temp,_ := strconv.Atoi(s)
	if temp % 2 == 0 {
		return temp
	}
	//
	res := -1
	n := len(s)
	if n == 1 {
		return res
	}
	//
	for i:=0;i<n-1;i++ {
		a, _ := strconv.Atoi(string(s[i]))
		if a % 2 == 0 {
			num := gennum(s, i)
			if num[0] == '0' {
				continue
			}
			res,_ = strconv.Atoi(num)
			break
		}
	}
	return res
}

func gennum(s string, i int) string {
	var res string
	if i == 0 {
		res = s[1:] + string(s[0])
	} else{
		res = s[:i] + s[i+1:] + string(s[i])
	}
	return res
}