/*
给定n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

提示：

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105

*/

package main

import "fmt"

// 1. 我的题解：先筛取最高值+再从两端算最高值之间的雨水

/*
执行用时：4 ms, 在所有 Go 提交中击败了96.36%的用户
内存消耗：4.9 MB, 在所有 Go 提交中击败了17.78%的用户
*/

func trap1(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}

	tops := []int{}
	i := 0
	for height[i] < height[i+1] {
		i++
		if i == len(height)-1 {
			return 0
		}
	}
	top := i // 第一个最高点
	tops = append(tops, top)
	for i < len(height)-1 {
		// 下降
		for i < len(height)-1 && height[i] >= height[i+1] {
			i++
		}
		if i == len(height)-1 {
			break
		}
		// 上升
		for i < len(height)-1 && height[i] <= height[i+1] {
			i++
		}

		top = i
		tops = append(tops, top)
	}
	// 8 3 4 7 9 2 6

	if len(tops) == 0 || len(tops) == 1 {
		return 0
	}
	left := tops[0]
	water := 0
	for j := 0; j < len(tops); j++ {
		if height[tops[j]] >= height[left] {
			for k := left; k < tops[j]; k++ {
				num := height[left] - height[k]
				if num >= 0 {
					water += num
				}
			}
			left = tops[j]
		}
	}
	right := tops[len(tops)-1]
	for j := len(tops) - 1; right > left; j-- {
		if height[tops[j]] >= height[right] {
			for k := right; k > tops[j]; k-- {
				num := height[right] - height[k]
				if num >= 0 {
					water += num
				}
			}
			right = tops[j]
		}
	}
	return water
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// 2. 我实现的官方题解：双指针

// 只要 right_max[i]>left_max[i] （元素 0 到元素 6），
// 积水高度将由 left_max 决定，
// 类似地 left_max[i]>right_max[i]（元素 8 到元素 11）。
/*
执行用时：8 ms, 在所有 Go 提交中击败了58.81%的用户内存消耗：4.3 MB, 在所有 Go 提交中击败了34.24%的用户
*/

func trap2(height []int) int {
	i, j := 0, len(height)-1
	water := 0
	leftmax, rightmax := 0, 0

	for i < j {
		if height[i] < height[j] {
			if height[i] >= leftmax {
				leftmax = height[i]
			} else {
				water += leftmax - height[i]
			}
			i++
		} else {
			if height[j] >= rightmax {
				rightmax = height[j]
			} else {
				water += rightmax - height[j]
			}
			j--
		}
	}
	return water
}

// 3. 我实现的官方题解：动态规划:左右各扫描一次
/*
执行用时：220 ms, 在所有 Go 提交中击败了5.07%的用户
内存消耗：8.1 MB, 在所有 Go 提交中击败了5.39%的用户
*/

func trap(height []int) int {
	left, right := []int{}, []int{}

	lmax, rmax := 0, 0
	for i := 0; i < len(height); i++ {
		if height[i] >= lmax {
			lmax = height[i]
		}
		left = append(left, lmax)
	}

	for i := len(height) - 1; i >= 0; i-- {
		if height[i] >= rmax {
			rmax = height[i]
		}
		right = append([]int{rmax}, right...)
	}

	water := 0
	for i := 0; i < len(height); i++ {
		water += min(left[i], right[i]) - height[i]
	}
	return water
}

func main() {
	height := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},    // 6
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 3}, // 11
		{4, 2, 0, 3, 2, 5},                      // 9
		{4, 2, 3},                               // 1
		{5, 4, 1, 2},                            // 1
		{5, 5, 4, 7, 8, 2, 6, 9, 4, 5},          // 10
		{1, 7, 8},                               // 0
	}
	res := []int{}
	for i := range height {
		res = append(res, trap(height[i]))
	}
	fmt.Println(res)
}
