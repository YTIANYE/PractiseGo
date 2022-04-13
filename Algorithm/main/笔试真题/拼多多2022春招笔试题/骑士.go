/**
n行m列的地图中，马从K位置，跳向T位置，问最少需要多少步？
马走日
地图中0位置代表石头，不能跳，且会蹩马脚；1的位置代表平地可以走，开始和终止位置都是平地

输入
测试样例组数T
每组中输入n, m
还有n行长度为m的字符串，表示地图
输出：
最少的步数， 如果不能到达，输出-1
*/

// 例子
/**
6
4 4
K111
1111
1111
111T
2 3
K10
00T
2 3
K01
11T
4 4
K111
1101
1011
111T
3 3
KT1
111
111
3 3
K11
1T1
111
*/

package main

import (
	"fmt"
	"math"
)

// 我的题解：60%通过，部分超时
func main() {
	var T int
	fmt.Scan(&T)
	for ; T != 0; T-- {
		var n, m int
		fmt.Scan(&n, &m)

		var kx, ky, tx, ty int
		ditu := [][]uint8{}
		for i := 0; i < n; i++ {
			var temp []uint8
			var s string
			fmt.Scan(&s)
			for j := range s {
				if s[j] == 'T' {
					tx, ty = i, j
				} else if s[j] == 'K' {
					kx, ky = i, j
				}
				temp = append(temp, s[j])
			}
			ditu = append(ditu, temp)
		}
		// fmt.Println(ditu)
		fmt.Println(moveToT(kx, ky, tx, ty, ditu))
	}
}

// bfs
func moveToT(kx, ky, tx, ty int, ditu [][]uint8) int { // 传入起始位置、目标位置和地图
	// 初始化
	n, m := len(ditu), len(ditu[0])
	// fmt.Println(n,m )
	dp := [][]int{}
	for i := 0; i < n; i++ {
		var temp []int
		for j := 0; j < m; j++ {
			temp = append(temp, math.MaxInt32)
		}
		// fmt.Println(temp)
		dp = append(dp, temp)
	}
	// fmt.Println(dp)
	dp[kx][ky] = 0

	// move
	var isright func(int, int) bool // 该位置是否可以跳到上面
	isright = func(a, b int) bool {
		if 0 <= a && a < n && 0 <= b && b < m && (ditu[a][b] == '1' || ditu[a][b] == 'T') { // 注意目标位置可以跳到上面
			return true
		}
		return false
	}

	var isStone func(int, int) bool // 该位置是否是石头
	isStone = func(a int, b int) bool {
		if ditu[a][b] == '0' {
			return true
		}
		return false
	}

	var nextlocal func(int, int) [][]int // 返回下一步可以跳到的全部位置
	nextlocal = func(x, y int) [][]int {
		res := [][]int{}
		next := [][]int{
			{x + 2, y + 1, x + 1, y}, // [0][1]下一步的位置， [2][3]马脚的位置
			{x + 2, y - 1, x + 1, y},
			{x - 2, y + 1, x - 1, y},
			{x - 2, y - 1, x - 1, y},
			{x - 1, y + 2, x, y + 1},
			{x + 1, y + 2, x, y + 1},
			{x - 1, y - 2, x, y - 1},
			{x + 1, y - 2, x, y - 1},
		}

		for _, step := range next {
			if isright(step[0], step[1]) && !isStone(step[2], step[3]) {
				res = append(res, []int{step[0], step[1]})
			}
		}
		return res
	}

	// 层次遍历
	queue := [][]int{}
	queue = append(queue, []int{kx, ky})
	for len(queue) != 0 {
		tqueue := make([][]int, len(queue))
		copy(tqueue, queue)
		// fmt.Println(tqueue)
		// tqueue := queue
		queue = [][]int{}
		for _, local := range tqueue {
			x, y := local[0], local[1]
			if x == tx && y == ty { // 到达
				return dp[x][y]
			}

			nextstep := nextlocal(x, y) // 下一步可走的位置
			for _, step := range nextstep {
				nextx, nexty := step[0], step[1]
				if dp[nextx][nexty] > dp[x][y]+1 {
					dp[nextx][nexty] = dp[x][y] + 1
					queue = append(queue, []int{nextx, nexty})
				}
			}
		}
	}
	return -1
}
