/**
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。



示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。


提示：

1 <= nums.length <= 3 * 104
0 <= nums[i] <= 105

*/
package main

/**
执行用时：1060 ms, 在所有 Go 提交中击败了5.08%的用户
内存消耗：6.7 MB, 在所有 Go 提交中击败了86.13%的用户
*/
// 我的题解：标记法-暴力法
func canJump1(nums []int) bool {
	n := len(nums)
	tags := make([]bool, n)
	tags[0] = true
	for i := 0; i < n; i++ {
		if tags[i] == false {
			return false
		}

		for num := nums[i]; num > 0; num-- {
			if i+num < n {
				tags[i+num] = true
			}
		}
	}
	return tags[n-1]
}

/**
执行用时：68 ms, 在所有 Go 提交中击败了23.19%的用户
内存消耗：6.7 MB, 在所有 Go 提交中击败了15.71%的用户
*/
// 我的题解：倒序方法
func canJump2(nums []int) bool {
	n := len(nums)
	minNum := n - 1
	for i := n - 2; i >= 0; i-- {
		if i+nums[i] >= minNum {
			minNum = i
		}
	}
	if minNum == 0 {
		return true
	}
	return false
}
/**
执行用时：52 ms, 在所有 Go 提交中击败了49.10%的用户
内存消耗：6.7 MB, 在所有 Go 提交中击败了82.49%的用户
 */
// 官方题解：贪心（正序方法）
func canJump(nums []int) bool {
	n := len(nums)
	right := 0
	for i := 0; i < n; i++ {
		if i <= right {
			if i+nums[i] > right {
				right = i + nums[i]
			}
			if right >= n-1 {
				return true
			}
		}
	}
	return false
}
