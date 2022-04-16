/**
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

 

示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"
 

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 109

*/

package main

import (
	"sort"
	"strconv"
)

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.6 MB, 在所有 Go 提交中击败了29.98%的用户
*/

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return comp(nums[i], nums[j])
	})

	var res string
	for _, num := range nums {
		res += strconv.Itoa(num)
	}

	// 去除前导 0
	for len(res) != 1 && res[0] == '0' {
		res = res[1:]
	}

	return res
}

func comp(a, b int) bool {
	s1 := strconv.Itoa(a)
	s2 := strconv.Itoa(b)

	if s1+s2 > s2+s1 {
		return true
	}
	return false
}
