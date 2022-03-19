/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

 

示例 1：


输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12
 

提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100

*/

package main

// 我的题解：动态规划

/*
执行用时：4 ms, 在所有 Go 提交中击败了97.85%的用户
内存消耗：3.7 MB, 在所有 Go 提交中击败了100.00%的用户
*/

func minPathSum(grid [][]int) int {

	n := len(grid)    // 行数
	m := len(grid[0]) // 列数
	if n == 1 && m == 1 {
		return grid[0][0]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[n-1][m-1]

}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
