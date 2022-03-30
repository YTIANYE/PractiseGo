/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为O(n) 的算法解决此问题。



示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9


提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109

*/

package main

// 我实现的官方题解:哈希表
/*
执行用时：424 ms, 在所有 Go 提交中击败了21.09%的用户
内存消耗：10.1 MB, 在所有 Go 提交中击败了17.25%的用户
*/
func longestConsecutive(nums []int) int {
	hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}

	maxLen := 0
	for _, num := range nums {
		_, ok := hash[num-1]
		if !ok { // num为一段连续数字的开头
			count := 0
			// 注意以后判断hash中是否存在某个元素，
			// 用 hash[x] != 0，来判断，不要用 hash[x] == 1来判断，
			// 因为hash初始化的时候遇见一个x,hash[x]++
			for hash[num] != 0 {
				num++
				count++
			}
			if count > maxLen {
				maxLen = count
			}
		}
	}
	return maxLen
}
