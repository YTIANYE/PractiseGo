/*
0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。

例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。


示例 1：
输入: n = 5, m = 3
输出:3
示例 2：
输入: n = 10, m = 17
输出:2

限制：
1 <= n<= 10^5
1 <= m <= 10^6

*/
package main

import "fmt"

// 我的题解：暴力解法（超时）
func lastRemaining1(n int, m int) int {
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	i := 1
	index := 0
	for len(nums) != 1 {
		if i == m {
			nums = append(nums[0:index], nums[index+1:]...)
			i = 1
			index = index % len(nums)
		} else {
			i++
			index = (index + 1) % len(nums)
		}
	}
	return nums[0]
}

// 官方题解：动态规划
/*
执行用时：4 ms, 在所有 Go 提交中击败了98.94%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了85.68%的用户
*/
func lastRemaining(n int, m int) int {
	res := 0
	for i := 2; i < n+1; i++ {
		res = (res + m) % i
	}
	return res
}

func main() {
	n, m := 5, 3
	// n, m := 70866, 116922
	fmt.Println(lastRemaining(n, m))
}
