/**
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。



示例 1：

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10

*/
package main

import (
	"sort"
	"strconv"
)

/**
执行用时：4 ms, 在所有 Go 提交中击败了8.58%的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了97.77%的用户
*/
// 递归
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	//
	temp := []int{}
	var dfs func(bool, int)
	dfs = func(choosed bool, index int) {
		if index == n {
			res = append(res, append([]int{}, temp...))
			return
		}
		dfs(false, index+1)
		if index > 0 && !choosed && nums[index-1] == nums[index] {
			return
		}
		temp = append(temp, nums[index])
		dfs(true, index+1)
		temp = temp[:len(temp)-1]
	}
	//
	dfs(false, 0)
	return res
}

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.7 MB, 在所有 Go 提交中击败了5.09%的用户
*/
// 子集中元素及个数相同，排列顺序不同时，算作统一子集
func subsetsWithDup0(nums []int) [][]int {
	sort.Ints(nums)
	hash := make(map[string]bool)
	n := len(nums)
	res := [][]int{}
	res = append(res, []int{})
	for i := 0; i < n; i++ {
		m := len(res)
		for j := 0; j < m; j++ {
			arr := make([]int, len(res[j]))
			copy(arr, res[j])
			arr = append(arr, nums[i])
			strs := toString(arr)
			if !hash[strs] {
				hash[strs] = true
				res = append(res, arr)
			}
		}
		num := strconv.Itoa(nums[i])
		if !hash[num] {
			hash[num] = true
			res = append(res, []int{nums[i]})
		}
	}

	return res
}

// 全部子集
func subsetsWithDup1(nums []int) [][]int {
	hash := make(map[string]bool)
	n := len(nums)
	res := [][]int{}
	res = append(res, []int{})
	for i := 0; i < n; i++ {
		m := len(res)
		for j := 0; j < m; j++ {
			arr := make([]int, len(res[j]))
			copy(arr, res[j])
			arr = append(arr, nums[i])
			strs := toString(arr)
			if !hash[strs] {
				hash[strs] = true
				res = append(res, arr)
			}
		}
		num := strconv.Itoa(nums[i])
		if !hash[num] {
			hash[num] = true
			res = append(res, []int{nums[i]})
		}
	}

	return res
}

// 连续子集
func subsetsWithDup2(nums []int) [][]int {
	hash := make(map[string][]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j <= n; j++ { // 注意 []也符合，所以j:=i
			str := toString(nums[i:j])
			hash[str] = []int{i, j}
		}
	}
	//
	res := [][]int{}
	for _, val := range hash {
		l, r := val[0], val[1]
		res = append(res, nums[l:r])
	}
	return res
}

func toString(nums []int) string {
	res := ""
	for i := range nums {
		num := strconv.Itoa(nums[i])
		if i == 0 {
			res += num
		} else {
			res += "," + num
		}
	}
	return res
}
