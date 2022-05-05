/**
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

 

示例 1：

输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]
示例 2：

输入：nums = [1,1]
输出：[2]
 

提示：

n == nums.length
1 <= n <= 105
1 <= nums[i] <= n
进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。

*/

package main

import "fmt"

// 我的题解：交换位置，比对nums[i]是否在nums中nums[i]的位置上
/**
执行用时：44 ms, 在所有 Go 提交中击败了52.04%的用户
内存消耗：7.4 MB, 在所有 Go 提交中击败了85.48%的用户
 */

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	res := []int{}

	for i := 0; i < n; i++ {
		// fmt.Println(nums)
		for nums[i]-1 != i && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// fmt.Println(nums)
	for i, num := range nums {
		if i+1 != num {
			res = append(res, i+1)
		}
	}
	return res
}

/**
4 3 2 7 8 2 3 1
7 3 2 4 8 2 3 1
3 3 2 4 8 2 7 1
2 3 3 4 8 2 7 1
3 2 3 4 8 2 7 1 // 3的位置相等
*/

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
