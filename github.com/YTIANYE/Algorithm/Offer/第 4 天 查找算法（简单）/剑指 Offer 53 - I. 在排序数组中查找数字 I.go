package main

import "fmt"

/*
统计一个数字在排序数组中出现的次数。

示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: 2

示例2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

提示：
0 <= nums.length <= 105
-109<= nums[i]<= 109
nums是一个非递减数组
-109<= target<= 109
*/

/*我的题解1：找出一个后向两边查找*/

// 执行用时：8 ms, 在所有 Go 提交中击败了41.74%的用户
// 内存消耗：3.9 MB, 在所有 Go 提交中击败了60.67%的用户
func search(nums []int, target int) int {
	count := 0
	mid := bisearch(nums, 0, len(nums)-1, target)
	if mid == -1{
		return 0
	}
	for i := mid; i < len(nums) && nums[i] == target; i++ {
		count++
	}
	for i := mid; i >-1 && nums[i] == target; i-- {
		count++
	}
	return count - 1
}

func bisearch(nums []int, left, right, target int) int {
	if left > right{
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] < target {
		return bisearch(nums, mid+1, right, target)
	} else if nums[mid] > target {
		return bisearch(nums, left, mid-1, target)
	} else {
		return mid
	}
}

func main(){
	// nums := []int{1}
	// target := 1
	nums := []int{5,7,7,8,8,10}
	target := 6
	fmt.Println(search(nums, target))
}
