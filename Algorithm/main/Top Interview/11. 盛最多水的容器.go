/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

 

示例 1：



输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1
 

提示：

n == height.length
2 <= n <= 105
0 <= height[i] <= 104


*/

package main

import "math"

// 我的题解：暴力超时
func maxArea1(height []int) int {

	n := len(height)
	if n < 2 {
		return 0
	}
	max := math.MinInt32

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if height[i] < height[j] {
				if height[i]*(j-i) > max {
					max = height[i] * (j - i)
				}
			} else {
				if height[j]*(j-i) > max {
					max = height[j] * (j - i)
				}
			}
		}
	}
	return max
}

// 我的题解：暴力剪枝
/**
执行用时：120 ms, 在所有 Go 提交中击败了5.55%的用户
内存消耗：8.5 MB, 在所有 Go 提交中击败了61.85%的用户
*/

func maxArea2(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}
	res := 0
	for i := 0; i < n; i++ {
		if height[i] != 0 {
			for j := i + res/height[i]; j < n; j++ {
				res = max(res, min(height[i], height[j])*(j-i))
			}
		}
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

// 我实现的官方题解：双指针
/*
执行用时：64 ms, 在所有 Go 提交中击败了93.71%的用户
内存消耗：8.7 MB, 在所有 Go 提交中击败了25.29%的用户
*/
func maxArea(height []int) int {
	res := 0
	i, j := 0, len(height)-1
	for i < j {
		res = max(res, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
