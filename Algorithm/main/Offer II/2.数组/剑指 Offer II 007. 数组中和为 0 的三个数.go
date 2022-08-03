/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c ，使得 a + b + c = 0 ？请找出所有和为 0 且 不重复 的三元组。



示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]


提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

package main

import (
	"sort"
	"strconv"
)

// 我实现的官方题解：排序+双指针
/**
执行用时：32 ms, 在所有 Go 提交中击败了67.59%的用户
内存消耗：7.3 MB, 在所有 Go 提交中击败了57.93%的用户
*/
func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	//
	sort.Ints(nums)
	n := len(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] { // 注意去重
			continue
		}
		l, r := i+1, n-1
		target := 0 - nums[i]
		for l < r {
			if l > i+1 && nums[l-1] == nums[l] { // 注意这里的去重
				l++
			} else if nums[l]+nums[r] < target {
				l++
			} else if nums[l]+nums[r] > target {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
			}
		}
	}
	return res
}

// 我的题解：二分查找，超时
func threeSum1(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	// 二分查找
	var biaFind func(int, int, int) bool
	biaFind = func(num, left, right int) bool {
		if left > right {
			return false
		}
		mid := (left + right) / 2
		if nums[mid] > num {
			return biaFind(num, left, mid-1)
		} else if nums[mid] < num {
			return biaFind(num, left+1, right)
		}
		return true

	}
	// 计算
	sort.Ints(nums)
	n := len(nums)
	hashmap := make(map[string]bool)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			num := 0 - nums[i] - nums[j]
			if biaFind(num, j+1, n-1) {
				nums2 := []int{nums[i], nums[j], num}
				nums2Str := toString(nums2)
				if _, ok := hashmap[nums2Str]; !ok {
					res = append(res, nums2)
					hashmap[nums2Str] = true
				}
			}
		}
	}
	return res
}

func toString(nums []int) string {
	res := ""
	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	return res
}
