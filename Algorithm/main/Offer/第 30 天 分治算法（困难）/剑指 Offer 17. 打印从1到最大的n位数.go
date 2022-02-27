/*
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]
 

说明：

用返回一个整数列表来代替打印
n 为正整数

*/
package main

import (
	"fmt"
	"math"
	"strconv"
)

// 我的题解：暴力解法
func printNumbers1(n int) []int {
	res := []int{}
	end := int(math.Pow(10, float64(n)))
	for i := 1; i < end; i++ {
		res = append(res, i)
	}
	return res
}

// 我的题解：分治
/*
执行用时：36 ms, 在所有 Go 提交中击败了7.52%的用户
内存消耗：13.1 MB, 在所有 Go 提交中击败了5.94%的用户
*/
func printNumbers(n int) []int {

	var fenzhi func(string, int) []string
	fenzhi = func(str string, n int) []string {
		if n == 0 {
			return []string{str}
		}
		res := []string{}
		for i := 0; i < 10; i++ {
			temp := fenzhi(str+strconv.Itoa(i), n-1)
			res = append(res, temp...)
		}
		return res
	}

	chars := fenzhi("", n)
	// fmt.Println(chars)
	res := []int{}
	for _, str := range chars {
		num, _ := strconv.Atoi(str)
		res = append(res, num)
	}
	return res[1:]
}

func main() {
	n := 2
	fmt.Println(printNumbers(n))
}
