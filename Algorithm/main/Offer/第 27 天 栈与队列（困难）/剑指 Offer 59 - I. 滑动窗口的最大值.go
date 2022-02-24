package main

import (
	"fmt"
	"math"
)

// 我的题解:滑动窗口 暴力解法

func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}

	window := []int{}
	res := []int{nums[0]}

	for i := 0; i < k; i++ {
		if nums[i] > res[0] {
			res[0] = nums[i]
		}
		window = append(window, nums[i]) // {位置， 值}
	}

	for i := k; i < len(nums); i++ {
		window = append(window, nums[i])
		if window[0] == res[i-k] {
			res = append(res, max(window[1:])) // 这里降低算法效率
		} else if nums[i] >= res[i-k] {
			res = append(res, nums[i])
		} else {
			res = append(res, res[len(res)-1])
		}
		window = window[1:]
	}

	return res
}

func max(nums []int) int {
	max := int(-1 * math.Pow(2, 31))
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

// 我实现的精选题解： 维护一个最大值队列
/*
执行用时：12 ms, 在所有 Go 提交中击败了98.00%的用户
内存消耗：6.4 MB, 在所有 Go 提交中击败了53.01%的用户
*/
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}

	res := []int{}
	queue := []int{} // 递减
	for i := 0; i < k; i++ {
		for len(queue) != 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	res = append(res, queue[0])

	for i := k; i < len(nums); i++ {
		if nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		for len(queue) != 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		res = append(res, queue[0])
	}
	return res
}

func main() {
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
