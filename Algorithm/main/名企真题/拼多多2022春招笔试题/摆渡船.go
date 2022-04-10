/*
题目：
m辆汽车，摆渡船每次运送n辆汽车，一趟时间是x， 往返2x,问最快多长时间运送完毕

输入： m, n, x
每辆汽车到达的时间
*/

// 事例

/*
例子1：
输入：
3 2 10
10
40
30
输出
50

例子2：
输入
10 2 10
0
10
20
30
40
50
60
70
80
90
输出
100

例子3：
输入
8 2 10
0
0
0
0
0
0
0
0
输出
70

例子4：
输入
4 2 10
0
0
0
100
输出
110
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

// 我的题解：当时未完成

func main() {
	var m, n, x int
	fmt.Scan(&m, &n, &x)
	cars := []int{}
	for i := 0; i < m; i++ {
		var time int
		fmt.Scan(&time)
		cars = append(cars, time)
	}
	sort.Ints(cars)

	// 动态规划
	dp := []int{}
	for i := 0; i <= m; i++ {
		dp = append(dp, math.MaxInt32)
	}
	dp[0] = 0

	for i := 1; i <= m; i++ {
		for j := i - 1; j >= 0 && j >= i-n; j-- { // 第i辆车加入
			lasttime := dp[j]
			dp[i] = min(dp[i], max(lasttime, cars[i-1])+2*x)
		}
		if i == m { // 最后一次不用返回
			dp[m] -= x
		}
	}
	fmt.Println(dp[m])
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
