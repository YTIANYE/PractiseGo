/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。


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
	"fmt"
	"sort"
)

/*
执行用时：56 ms, 在所有 Go 提交中击败了15.22%的用户
内存消耗：7 MB, 在所有 Go 提交中击败了97.89%的用户
*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	res := [][]int{}
	sort.Ints(nums)

	for first := 0; first < n-2; first++ {
		if first != 0 && nums[first] == nums[first-1] {
			continue
		}

		i := first + 1
		j := n - 1
		for i < j {
			num := nums[first] + nums[i] + nums[j]
			if num > 0 {
				j--
				for j > i && nums[j] == nums[j+1] {
					j--
				}
			} else if num < 0 {
				i++
				for j > i && nums[i] == nums[i-1] {
					i++
				}
			} else if num == 0 {
				res = append(res, []int{nums[first], nums[i], nums[j]})
				j--
				for j > i && nums[j] == nums[j+1] {
					j--
				}
				i++
				for j > i && nums[i] == nums[i-1] {
					i++
				}
			}
		}
	}
	return res
}

func main() {
	// nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

// 我的题解：暴力题解(超时)

func threeSum1(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	var res [][]int

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			num := nums[i] + nums[j]
			for k := j + 1; k < n; k++ {
				if num+nums[k] == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					sort.Ints(temp)
					if !inNums(res, temp) {
						res = append(res, temp)
					}
				}
			}
		}
	}
	return res
}

func inNums(res [][]int, temp []int) bool {
	for _, r := range res {
		count := 0
		for i := range r {
			if r[i] == temp[i] {
				count++
			}
		}
		if count == 3 {
			return true
		}
	}
	return false
}

// 我的题解：排序+剪枝（超时）

func threeSum2(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-1; j++ {
			num := nums[i] + nums[j]
			if num > 0 {
				break
			}
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}

			for k := j + 1; k < n; k++ {
				if k+1 != n && nums[k] == nums[k+1] {
					continue
				}
				if num+nums[k] > 0 {
					break // 后边的数都可以不用考虑了
				}
				if num == -nums[k] {
					temp := []int{nums[i], nums[j], nums[k]}

					res = append(res, temp)
					break
				}
			}
		}
	}
	return res
}
