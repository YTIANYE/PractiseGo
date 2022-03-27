/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。


示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'

*/

package main

import "fmt"

// 我的题解：深度遍历
/*
执行用时：4 ms, 在所有 Go 提交中击败了90.44%的用户
内存消耗：4.1 MB, 在所有 Go 提交中击败了27.50%的用户
*/


func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	var visit [][]bool
	for i := 0; i < m; i++ {
		visit = append(visit, []bool{})
		for j := 0; j < n; j++ {
			visit[i] = append(visit[i], false)
		}
	}

	var dfsgrid func(int, int)
	dfsgrid = func(i int, j int) {
		if i >= m || i < 0 || j >= n || j < 0 { // 注意终止条件
			return
		}

		if visit[i][j] == false {
			visit[i][j] = true
			if grid[i][j] == '1' {
				dfsgrid(i+1, j)
				dfsgrid(i, j+1)
				dfsgrid(i-1, j) // 注意可以往左找
				dfsgrid(i, j-1)
			}
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visit[i][j] == false {
				if grid[i][j] == '1' {
					res++
				}
				dfsgrid(i, j)
			}
		}
	}

	return res

}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	fmt.Println(numIslands(grid))

}
