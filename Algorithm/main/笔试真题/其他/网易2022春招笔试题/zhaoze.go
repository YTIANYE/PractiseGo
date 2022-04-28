package main

/*
题目：
n行m列的地图，每个位置的值是0 或者1 ，同值之间行走代价是1，不同值是2
从左上角走到右下角，只能是向左、向右、向下走，
问最少代价是多少？


例子：
3 3
1 0 0
1 1 1
0 0 1

例子：
3 3
1 1 1
0 0 0
1 1 0

*/

import (
	"fmt"
	"math"
)

// 我的题解：当时没写完

func main() {
	var n, m int // 行数 列数
	fmt.Scan(&n, &m)
	maps := [][]int{}
	tags := [][]int{} // 走到当前位置的最小步数
	for i := 0; i < n; i++ {
		tag := []int{}
		nums := []int{}
		for j := 0; j < m; j++ {
			var num int
			fmt.Scan(&num)
			nums = append(nums, num)
			tag = append(tag, math.MaxInt32)
		}
		maps = append(maps, nums)
		tags = append(tags, tag)
	}

	// 走
	var step func(a, b, x, y int) int
	step = func(a, b, x, y int) int { // (x, y) -> (a, b)
		if x < 0 || x >= n || y < 0 || y >= m {
			return math.MaxInt32
		} else if tags[x][y] == math.MaxInt32 {
			return math.MaxInt32
		} else {
			if maps[x][y] == maps[a][b] {
				return 1 + tags[x][y]
			} else {
				return 2 + tags[x][y]
			}
		}
	}

	// 计算
	tags[0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			left := step(i, j, i, j-1)
			right := step(i, j, i, j+1)
			up := step(i, j, i-1, j)
			tags[i][j] = min(min(min(left, right), up), tags[i][j])
		}
		for j:=m-1;j>=0;j--{
			left := step(i, j, i, j-1)
			right := step(i, j, i, j+1)
			tags[i][j] = min(min(left, right), tags[i][j])
		}
	}

	fmt.Println(tags[n-1][m-1])

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
