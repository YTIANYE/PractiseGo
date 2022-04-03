/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]
 

提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109

 */
package main

import "sort"

// 我的题解：二分查找
/**
执行用时：8 ms, 在所有 Go 提交中击败了26.40%的用户
内存消耗：3.8 MB, 在所有 Go 提交中击败了99.88%的用户
 */

func searchRange(nums []int, target int) []int {
	n := len(nums)
	begin := -1
	end := -1

	var biaFind func (left, right int)
	biaFind = func(left, right int) {
		if left > right{
			return
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			i:=mid
			for ;i<n && nums[i] == target;i++{}
			end = i-1
			j := mid
			for ;j>=0 && nums[j] == target;j--{}
			begin = j+1
			return
		}
		biaFind(left, mid-1)
		biaFind(mid+1, right)
	}
	biaFind(0, n-1)
	return []int{begin, end }
}

// 官方题解：调用库函数
func searchRange1(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target + 1) - 1
	return []int{leftmost, rightmost}
}
