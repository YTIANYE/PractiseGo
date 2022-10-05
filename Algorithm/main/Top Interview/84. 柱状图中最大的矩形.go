/**
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。



示例 1:



输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10
示例 2：



输入： heights = [2,4]
输出： 4


提示：

1 <= heights.length <=105
0 <= heights[i] <= 104

*/
package main

/**
执行用时：96 ms, 在所有 Go 提交中击败了40.59%的用户
内存消耗：8.8 MB, 在所有 Go 提交中击败了55.88%的用户
*/
// 我实现的官方题解：单调栈
func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)

	//
	stack := []int{}
	for i, height := range heights {
		for j := len(stack) - 1; j >= 0 && heights[stack[j]] >= height; j-- {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	//
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		height := heights[i]
		for j := len(stack) - 1; j >= 0 && heights[stack[j]] >= height; j-- {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	res := 0
	for i, height := range heights {
		res = max84(res, height*(right[i]-left[i]-1))
	}
	return res
}

/**
执行用时：88 ms, 在所有 Go 提交中击败了87.83%的用户
内存消耗：8.7 MB, 在所有 Go 提交中击败了60.45%的用户
*/
// 单调栈优化——我实现的官方题解
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)

	for i := range right {
		right[i] = n
	}
	stack := []int{}
	for i := 0; i < n; i++ {
		j := len(stack) - 1 // 指向栈中的最后一个数
		for j >= 0 && heights[stack[j]] >= heights[i] {
			right[stack[j]] = i
			stack = stack[:j]
			j--
		}
		if j < 0 {
			left[i] = -1
		} else {
			left[i] = stack[j]
		}
		stack = append(stack, i)
	}
	//
	res := 0
	for i, height := range heights {
		res = max84(res, height*(right[i]-left[i]-1))
	}
	return res
}

func max84(a, b int) int {
	if a > b {
		return a
	}
	return b
}
