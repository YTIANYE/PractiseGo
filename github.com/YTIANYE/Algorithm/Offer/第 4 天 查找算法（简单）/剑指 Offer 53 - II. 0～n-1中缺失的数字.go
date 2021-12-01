package main

import "fmt"

/*
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

示例 1:
输入: [0,1,3]
输出: 2

示例2:
输入: [0,1,2,3,4,5,6,7,9]
输出: 8

限制：
1 <= 数组长度 <= 10000
*/

/*我的题解：暴力解法*/

// 执行用时：16 ms, 在所有 Go 提交中击败了79.91%的用户
// 内存消耗：6.1 MB, 在所有 Go 提交中击败了99.91%的用户

func missingNumber1(nums []int) int {
	num := len(nums)
	for i, n := range nums {
		num += i - n
	}
	return num
}

/*我实现的官方题解*/

// 执行用时：16 ms, 在所有 Go 提交中击败了79.91%的用户
// 内存消耗：6.1 MB, 在所有 Go 提交中击败了99.91%的用户
func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	nums := []int{0}
	fmt.Println(missingNumber(nums))

}
