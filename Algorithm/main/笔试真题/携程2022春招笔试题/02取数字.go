/**
输入 n 个数，
一个长度为n的字符串，包含RB表示数字是红色或者蓝色

问：选取两个值相同颜色不同的数，有多少种可能？
 */

package main

import (
	"fmt"
)

// 通过样例 31.25%（不知道为什么还会有超时）

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	var color string
	fmt.Scan(&color)

	// 分类
	hashR := make(map[int]int) // key 数字  value: 个数
	hashB := make(map[int]int)
	for i := 0; i < n; i++ {
		if color[i] == 'R' {
			hashR[nums[i]]++
		} else {
			hashB[nums[i]]++
		}
	}

	// 计算
	res := 0
	for i := range hashB {
		res += hashB[i] * hashR[i]
	}
	fmt.Println(res)

}
