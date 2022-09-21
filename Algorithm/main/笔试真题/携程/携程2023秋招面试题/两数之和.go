package main

import "sort"

// 最优解：哈希
/**
执行用时：4 ms, 在所有 Go 提交中击败了94.89%的用户
内存消耗：4.1 MB, 在所有 Go 提交中击败了57.89%的用户
 */
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, num := range nums {
		if j, ok := hash[target-num];ok {
			return []int{j, i}
		}else {
			hash[num] = i
		}
	}
	return []int{}
}

func twoSum1(nums []int, target int) []int {
	old := make([]int, len(nums))
	copy(old, nums)

	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	var res []int
	var x, y int
	for i < j {
		x = nums[i]
		y = nums[j]
		if x+y > target {
			j--
		} else if x+y < target {
			i++
		} else {
			break
		}
	}

	for k, v := range old {
		if v == x || v == y {
			res = append(res, k)
		}
	}
	return res
}

/**
执行用时：8 ms, 在所有 Go 提交中击败了43.09%的用户
内存消耗：6.1 MB, 在所有 Go 提交中击败了5.99%的用户
*/
func twoSum2(nums []int, target int) []int {
	hash := make(map[int][]int)
	for i, num := range nums {
		if len(hash[num]) > 0 {
			hash[num] = append(hash[num], i)
		} else {
			hash[num] = []int{i}
		}
	}

	sort.Ints(nums)
	//
	n := len(nums)
	i := 0
	j := n - 1
	for i < j {
		if nums[i]+nums[j] == target {
			var a, b int
			if nums[i] == nums[j] {
				a, b = hash[nums[i]][0], hash[nums[j]][1]
			} else {
				a, b = hash[nums[i]][0], hash[nums[j]][0]
			}
			return []int{a, b}
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}
