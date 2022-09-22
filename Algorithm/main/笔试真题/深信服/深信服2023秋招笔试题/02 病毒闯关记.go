/**
n个数，病毒攻击范围是x
每个数有一个值v
病毒有一个值m的话，如果|m-v| <= x的话，病毒攻击成功
病毒需要依次按顺序进行攻击
当病毒的m值在当前情况不能攻击的话，需要初始化成一个新的可以攻击的值
问病毒要至少初始化为一个值几次，才能攻击完成
第一次初始化不算
 */

package main

import "fmt"

// 100%
func main() {
	var n, x int
	fmt.Scan(&n, &x)
	nums := make([]int, n)
	for i:= range nums {
		fmt.Scan(&nums[i])
	}
	//
	res := 0
	down, high := nums[0] - x, nums[0] + x
	for i:=1;i<n;i++ {
		d := max(down, nums[i]-x)
		h := min(high, nums[i]+x)
		if d > h {
			res++
			down, high = nums[i] - x, nums[i] + x
		}else {
			down = d
			high = h
		}
	}
	//
	fmt.Println(res )
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int ) int {
	if b > a  {
		return b
	}
	return a
}
