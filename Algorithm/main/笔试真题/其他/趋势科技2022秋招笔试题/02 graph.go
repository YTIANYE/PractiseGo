/**
题目： 一幅地图中的1连在一起构成一个图形，问地图中有几个图形，所有图形内各个点之间最大距离是多少？


示例：
输入
[
[0, 0, 0, 0, 0, 1, 1, 1, 1, 1],
[0, 1, 0, 0, 0, 1, 0, 0, 0, 1],
[0, 0, 0, 0, 0, 1, 0, 0, 0, 1],
[0, 0, 0, 0, 0, 1, 1, 1, 1, 1],
[0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
[0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
[0, 0, 0, 1, 1, 1, 1, 1, 0, 0],
[1, 1, 1, 1, 1, 1, 1, 1, 1, 0]
]

输出：图形个数和最大距离
[]int{3,8}
*/

package main

import "fmt"

// 当时没做完
func tuxing(graph [][]int) []int {
	if len(graph) == 0 || len(graph[0]) == 0 {
		return []int{0, 0}
	}
	n, m := len(graph), len(graph[0])
	// 访问足迹
	tag := make([][]bool, n)
	for i := 0; i < n; i++ {
		tag[i] = make([]bool, m)
	}
	// 坐标是否越界
	var isRight func(int, int) bool
	isRight = func(x, y int) bool {
		if x < 0 || x >= n || y < 0 || y >= m {
			return false
		}
		return true
	}
	// 找到一个图形
	var find func(int, int, *[][]int)
	find = func(x int, y int, ptu *[][]int) {
		if !isRight(x, y) || tag[x][y] || graph[x][y] == 0 {
			return
		}
		//
		tag[x][y] = true
		*ptu = append(*ptu, []int{x, y})
		find(x+1, y, ptu)
		find(x-1, y, ptu)
		find(x, y+1, ptu)
		find(x, y-1, ptu)
		return
	}
	// 求图形内部最远距离
	var maxDistence func([][]int) int
	maxDistence = func(tu [][]int) int {
		res := 0
		for i := 0; i < len(tu); i++ {
			for j := i; j < len(tu); j++ {
				a := tu[i]
				b := tu[j]
				dis := abs(a[0]-b[0]) + abs(a[1]-b[1])
				res = max(res, dis)
			}
		}
		return res
	}
	// 求解
	count, maxNum := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tu := [][]int{}
			find(i, j, &tu)
			if len(tu) != 0 {
				maxNum = max(maxNum, maxDistence(tu))
				count++
			}
		}
	}
	return []int{count, maxNum}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	graph := [][]int{
		{0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 1, 0, 0, 0, 1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
	}
	fmt.Println(tuxing(graph))
}
