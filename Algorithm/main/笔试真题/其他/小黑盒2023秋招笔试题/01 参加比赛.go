package main

import "sort"

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param N int整型
 * @param results int整型一维数组
 * @return int整型一维数组
 */

// 100%

func CalcPoints(N int, results []int) []int {
	// write code here
	res := make([]int, N)
	nums := [][]int{}
	for i, r := range results {
		if r == -1 {
			res[i] = 0
		}else {
			nums = append(nums, []int{i, r})
		}
	}
	//
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][1] < nums[j][1]
	})
	//

	lastv := -1
	n := N
	for index, num := range nums {
		i, v := num[0], num[1]

		if lastv != v {
			n = N-index
			lastv = v

		}
		res[i] = n
	}
	return res
}
