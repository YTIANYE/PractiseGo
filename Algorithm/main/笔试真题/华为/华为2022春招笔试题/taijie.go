/**
题目：
n个台阶，从第一个台阶开始出发，每个台阶上有魔法值，表示在这个台阶上接下来能走的最大台阶数
从第一个台阶出发，能否在k步以内到达第n个台阶，不可以输出-1，可以输出步数

输入：
台阶数
魔法值
k

例子1：
输入：
6
2 1 5 0 3 2
2
输出
2

例子2：
6
2 1 2 0 0 1
2
输出
-1
 */


package main

import (
	"fmt"
	"math"
)

// 我的题解：动态规划： 通过率100%

func main() {
	var n int
	fmt.Scan(&n)
	nums := []int{}
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	var k int
	fmt.Scan(&k)

	if n == 1 {
		fmt.Println(0)
	}

	count := make([]int, n)
	for i := range count{
		count[i] = math.MaxInt32// 注意初始值
	}
	count[0] = 0

	// 计算  ： 动态规划

	for i := 0; i < n; i++ {
		for j := 0; j <= nums[i]; j++ {
			if i+j <n{
				count[i+j] = min(count[i+j], count[i] +1)
			}else{
				break
			}
		}
	}

	if count[len(count)-1] <= k && count[len(count)-1] > 0{
		fmt.Println(count[len(count)-1])
	}else{
		fmt.Println(-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
