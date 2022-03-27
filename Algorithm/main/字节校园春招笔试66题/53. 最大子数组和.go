/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

 

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23
 

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
 

进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

*/

package main

import "fmt"

// 我的题解：动态规划
/*
执行用时：120 ms, 在所有 Go 提交中击败了7.27%的用户内存消耗：24.8 MB, 在所有 Go 提交中击败了5.13%的用户
*/
func maxSubArray(nums []int) int {
	res := nums[0]

	var maxSub func([]int) int
	maxSub = func(nums []int) int{
		lengh := len(nums)
		if lengh == 1{
			return nums[0]
		}
		last := maxSub(nums[:lengh-1])
		temp := max(nums[lengh-1], last+ nums[lengh-1])
		res = max(temp, res )
		return temp
	}

	maxSub(nums)
	return res
}



func max(a, b int) int{
	if a>b{
		return a
	}
	return b
}

func main(){
	nums := []int{1}
	fmt.Println(maxSubArray(nums))
}