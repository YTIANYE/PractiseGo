/**
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。



示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0


提示：

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105


进阶：

如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。

*/

package main

import "math"

// 我的题解：滑动窗口
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：3.6 MB, 在所有 Go 提交中击败了64.92%的用户
*/

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	i, j := 0, 0 // 左闭右开
	res := math.MaxInt
	sum := 0
	for i < n || j < n {
		if sum < target {
			if j == n {
				break
			}
			sum += nums[j]
			j++
		} else {
			res = min(res, j-i)
			sum -= nums[i]
			i++
		}
	}
	//
	if res == math.MaxInt {
		return 0
	}
	return res
}

// 官方题解：更简洁
func minSubArrayLen1(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
