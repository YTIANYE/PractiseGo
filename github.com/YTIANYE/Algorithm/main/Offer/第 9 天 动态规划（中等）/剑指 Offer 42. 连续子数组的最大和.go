/*
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。


示例1:
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释:连续子数组[4,-1,2,1] 的和最大，为6。

提示：
1 <=arr.length <= 10^5
-100 <= arr[i] <= 100
*/

package main

/*我的题解：暴力方法（超时）*/
func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxsumcur := -100
	// maxsumpre := 0
	maxsum := -100
	for i, _ := range nums {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			maxsumcur = max(sum, maxsumcur)
		}
		maxsum = max(maxsumcur, maxsum)
	}
	return maxsum
}

func maxSubArray(nums []int) int {
	maxpre := 0
	maxsum := -100
	for _, num := range nums {
		maxpre = max(num, maxpre+num) // 以当前值做结尾的子数组的最大和 取决于 前一个值做结尾的子数组的最大和 和当前值
		maxsum = max(maxsum, maxpre)
	}
	return maxsum
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
