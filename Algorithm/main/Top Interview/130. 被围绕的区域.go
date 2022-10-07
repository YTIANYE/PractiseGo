/**
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
 

示例 1：


输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
示例 2：

输入：board = [["X"]]
输出：[["X"]]
 

提示：

m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] 为 'X' 或 'O'

*/
package main

/**
执行用时：20 ms, 在所有 Go 提交中击败了48.89%的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了49.80%的用户
*/

// 我的题解：将不需要改变的做标记
func solve(board [][]byte) {
	m, n := len(board), len(board[0]) // 行数，列数
	tags := make([][]bool, m)
	for i := range tags {
		tags[i] = make([]bool, n)
	}

	// 将和边框相连的'O' 全部标记为 true
	var search func(x, y int)
	search = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if board[x][y] == 'X' {
			return
		}
		if tags[x][y] == true {
			return
		}
		tags[x][y] = true
		search(x-1, y)
		search(x+1, y)
		search(x, y+1)
		search(x, y-1)
	}

	// 搜索边框
	for i := 0; i < n; i++ {
		search(0, i)
		search(m-1, i)
	}
	for i := 0; i < m; i++ {
		search(i, 0)
		search(i, n-1)
	}

	//
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && tags[i][j] == false {
				board[i][j] = 'X'
			}
		}
	}
}
