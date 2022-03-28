/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。


示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
 

提示：

你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间。
*/

package main

// 我的题解：二分查找
/*
执行用时：36 ms, 在所有 Go 提交中击败了22.07%的用户
内存消耗：6.6 MB, 在所有 Go 提交中击败了39.48%的用户
*/
func search(nums []int, target int) int {
	res := -1

	var bioserch func(int, int)
	bioserch = func(left, right int) {
		if left > right {
			return
		}

		mid := (left + right) / 2

		if nums[mid] == target {
			res = mid
			return
		}
		bioserch(left, mid-1)
		bioserch(mid+1, right)

	}
	bioserch(0, len(nums)-1)

	return res
}
