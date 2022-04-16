/**
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

 

示例 1：

输入：[3,2,3]
输出：3
示例 2：

输入：[2,2,1,1,1,2,2]
输出：2
 

进阶：

尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

*/

package main

/**
执行用时：12 ms, 在所有 Go 提交中击败了93.29%的用户
内存消耗：5.9 MB, 在所有 Go 提交中击败了99.04%的用户
*/
func majorityElement(nums []int) int {
	count := 0
	var res int
	for _, num := range nums {
		if count == 0 {
			res = num
		}
		if num == res {
			count++
		} else {
			count--
		}
	}
	return res
}
