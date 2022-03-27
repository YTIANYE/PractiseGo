/*
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
示例 1：

输入：[10,1,2]
输出：2110
示例 2：

输入：[3,30,34,5,9]
输出：9534330

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.7 MB, 在所有 Go 提交中击败了100.00%的用户
*/

func main() {
	// 获取数组
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	var s []string
	s = strings.Split(input.Text(), ",")
	var strs []string
	for i:=0;i<len(s);i++{
		temp := s[i]
		if i ==0{
			temp = temp[1:]
		}
		if i == len(s)-1{
			temp = temp[:len(temp)-1]
		}
		strs = append(strs, temp)
	}

	for i:=0;i<len(strs)-1;i++{
		for j:= i+1;j<len(strs);j++{
			if !compStrs(strs[i], strs[j]){
				strs[i], strs[j] = strs[j], strs[i]
			}
		}
	}

	res := ""
	for _, str := range strs{
		res = res + str
	}
	fmt.Println(res)
}

func compStrs(a, b string) bool {
	if a+b > b+a {
		return true
	}
	return false
}
