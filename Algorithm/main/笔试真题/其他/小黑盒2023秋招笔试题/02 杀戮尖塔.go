package main

import "sort"

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param N int整型
 * @param cost int整型一维数组
 * @param damage int整型一维数组
 * @return int整型
 */
// 100%
func Calc( N int ,  cost []int ,  damage []int ) int {
	// write code here
	res := 0
	nums1 := []int{}
	nums2 := 0
	nums3 := 0
	for i:=0;i<N;i++ {
		if cost[i] == 1 {
			nums1 = append(nums1, damage[i])
		}else if cost[i] == 2 {
			nums2 = max(nums2, damage[i])
		}else if cost[i] == 3 {
			nums3 = max(nums3, damage[i])
		}
	}
	//
	sort.Ints(nums1)
	l := len(nums1)

	res = max(res, nums3)
	res = max(res, nums2)
	if l >= 1 {
		res = max(res, nums2 + nums1[l-1])
	}
	if l >= 3 {
		res = max(res, nums1[l-1] + nums1[l-2] + nums1[l-3])
	}else if l == 2 {
		res = max(res, nums1[l-1] + nums1[l-2] )
	}else if l == 1 {
		res = max(res, nums1[l-1])
	}

	return res



}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}