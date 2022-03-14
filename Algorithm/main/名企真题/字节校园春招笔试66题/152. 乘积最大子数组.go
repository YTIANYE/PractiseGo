/*
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 32-位 整数。

子数组 是数组的连续子序列。

 

示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 

提示:

1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数


*/

package main

import "math"

// 我的题解：动态规划
/*
执行用时：4 ms, 在所有 Go 提交中击败了94.22%的用户
内存消耗：4.7 MB, 在所有 Go 提交中击败了5.12%的用户
*/
func maxProduct(nums []int) int {
	res := math.MinInt32
	var maxpro func([]int) (int, int) // 返回nums[-1]为结尾的子数组乘积的最大值和最小值
	maxpro = func(nums []int) (int, int) {
		if len(nums) == 1 {
			res = max(res, nums[0])
			return nums[0], nums[0]
		}

		length := len(nums)
		lastmin, lastmax := maxpro(nums[:length-1])
		curr := nums[length-1]
		currmax := max(max(lastmin*curr, lastmax*curr), curr)
		currmin := min(min(lastmin*curr, lastmax*curr), curr)
		res = max(currmax, res)
		return currmax, currmin
	}
	maxpro(nums)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
