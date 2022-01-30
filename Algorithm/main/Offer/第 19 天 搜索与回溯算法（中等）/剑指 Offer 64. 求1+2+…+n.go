/*
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

 

示例 1：

输入: n = 3
输出: 6
示例 2：

输入: n = 9
输出: 45
 

限制：

1 <= n <= 10000

*/

package main

import "math"

// 我的题解:平方+右移
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.8 MB, 在所有 Go 提交中击败了99.78%的用户
*/

func sumNums1(n int) int {
	// n(n+1)/2 = (n*n + n)/2

	return (int(math.Pow(float64(n), 2)) + n)>>1
}

// 我实现的官方题解:递归
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了81.42%的用户
*/
func sumNums(n int) int{
	res := 0
	var sum func(int) bool
	sum = func(n int) bool{
		res += n
		return n>0 && sum(n-1)
	}
	sum(n)
	return res
}