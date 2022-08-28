/**
一个长度为n的数组中删除一个元素后，再计算长度为k的子数组和的最大值，

输入一个数组
输出最大值
 */

package main

import (
	"fmt"
)

// 50%
func main() {
	var n,k int
	fmt.Scan(&n, &k)
	arrs:= make([]int, n)
	for i:=0;i<n;i++ {
		fmt.Scan(&arrs[i])
	}
	//
	res := 0
	var sum int
	for i:=0;i<n;i++ {
		nums := remove(arrs, i)
		sum = he(nums, k)
		res = max(res, sum)
	}
	fmt.Println(res )
}

func remove(arrs []int, i int) []int{
	nums := []int{}
	for j:= range arrs {
		if i == j {
			continue
		}
		nums = append(nums, arrs[j])
	}
	return nums
}

func he(nums []int, k int) int {
	p:=0
	q := k-1
	sum := 0
	for i:=0;i<=q;i++ {
		sum += nums[i]
	}
	res := sum // 最大和
	// 左闭右闭
	q++
	p++
	for q < len(nums) {
		sum -= nums[p-1]
		sum += nums[q]
		p++
		q++
		res = max(res, sum)
	}
	return res
}

func max(a, b int ) int {
	if a > b {
		return a
	}
	return b
}
