/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

 
示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1
 

提示：

1 <= nums.length <= 2500
-104 <= nums[i] <= 104
 

进阶：

你能将算法的时间复杂度降低到 O(n log(n)) 吗?

*/

package main

import "fmt"

// 我的题解：动态规划， 新加入一个结点，以当前结点为结尾的最长递增子序列
/*
执行用时：56 ms, 在所有 Go 提交中击败了59.89%的用户
内存消耗：4.1 MB, 在所有 Go 提交中击败了5.03%的用户
*/

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	// 初始化记录表
	zui := []int{}
	for i := 0; i < n; i++ {
		zui = append(zui, 1)
	}
	zui[0] = 1
	// 迭代更新
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- { // zui 的记录表遍历
			if nums[i] > nums[j] {
				zui[i] = max(zui[i], zui[j]+1)
			}
		}
	}
	// 计算结果
	res := 0
	for _, v := range zui {
		if res < v {
			res = v
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}
