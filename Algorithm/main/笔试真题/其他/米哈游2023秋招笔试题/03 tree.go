/**
题目：
给一个树的结点涂色，保证除叶子结点以外，每个结点与其相邻结点颜色不同，且所有相邻节点有两种不同于自己的颜色，输出一种情况即可


样例：
5
0 1
2 0
1 3
0 4

RGBBR
*/

package main

import (
	"fmt"
	"io"
)

// 地图着色问题，应该有更好的最优解

// 我的题解：未得到验证
func main() {

	// 输入
	var n int
	fmt.Scan(&n)
	tree := [][]int{}
	for i := 0; i < 4; i++ { // 为了测试
		var a, b int
		if _, ok := fmt.Scan(&a, &b); ok != io.EOF {
			tree = append(tree, []int{a, b})
		} else {
			break
		}
	}

	// 存储每个点的相邻结点
	sons := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		sons[i] = []int{}
	}
	fathers := make(map[int]int, n)
	for _, bian := range tree {
		a, b := bian[0], bian[1]
		a, b = min(a, b), max(a, b)
		sons[a] = append(sons[a], b)
		fathers[b] = a
	}

	// 记录每个点颜色情况
	colors := make([]byte, n)
	for key := 0; key < n; key++ {
		colors[key] = 'N' // 当前没有颜色
	}

	// 涂色
	queue := []int{0}
	colors[0] = 'R'
	for len(queue) != 0 {
		tqueue := queue
		queue = []int{}
		for len(tqueue) != 0 {
			node := tqueue[0]
			tqueue = tqueue[1:]
			rightColor := []byte{}
			colormap := make(map[byte]bool)
			colormap['R'] = true
			colormap['G'] = true
			colormap['B'] = true
			if colors[fathers[node]] == 'N' {
				delete(colormap, colors[node])
				for key := range colormap {
					rightColor = append(rightColor, key)
				}
				colors[sons[node][0]] = rightColor[0]
				for i := 1; i < len(sons[node]); i++ {
					colors[sons[node][i]] = rightColor[1]
				}
			} else {
				delete(colormap, colors[node])
				delete(colormap, colors[fathers[node]])
				for key := range colormap {
					rightColor = append(rightColor, key)
				}
				for i := 0; i < len(sons[node]); i++ {
					colors[sons[node][i]] = rightColor[0]
				}
			}
			queue = append(queue, sons[node]...)
		}
	}

	// 涂色结果
	for i := 0; i < n; i++ {
		fmt.Print(string(colors[i]))
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
