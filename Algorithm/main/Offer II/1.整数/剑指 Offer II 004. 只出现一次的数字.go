/**
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

 

示例 1：

输入：nums = [2,2,3,2]
输出：3
示例 2：

输入：nums = [0,1,0,1,0,1,100]
输出：100
 

提示：

1 <= nums.length <= 3 * 104
-231 <= nums[i] <= 231 - 1
nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
 */

package main


// 官方题解：数字电路设计
func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		b = (b ^ num) &^ a
		a = (a ^ num) &^ b
	}
	return b
}
