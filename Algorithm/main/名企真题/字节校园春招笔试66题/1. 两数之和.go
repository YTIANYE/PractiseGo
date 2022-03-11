/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]

提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？

*/

package main

import "sort"


// 我的题解：排序 + 双指针
/*
执行用时：8 ms, 在所有 Go 提交中击败了44.97%的用户内存消耗：3.6 MB, 在所有 Go 提交中击败了79.03%的用户
*/
func twoSum(nums []int, target int) []int {
	old := make([]int, len(nums))
	copy(old, nums)

	sort.Ints(nums)
	i:=0
	j := len(nums)-1
	var res []int
	var x, y int
	for i<j{
		x = nums[i]
		y = nums[j]
		if x+y>target{
			j--
		}else if x+y<target{
			i++
		}else{
			break
		}
	}

	for k, v := range old{
		if v == x || v == y{
			res = append(res, k)
		}
	}
	return res
}
