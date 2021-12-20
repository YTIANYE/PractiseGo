/*
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

示例 1:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

*/
package main

import "fmt"

/*我的题解： 动态规划
注意：我的题解是选择子矩阵中和最大的一个，由于礼物值均大于0，所以最终的子矩阵也就是整个大矩阵
*/

// 执行用时：4 ms, 在所有 Go 提交中击败了98.93%的用户
// 内存消耗：3.8 MB, 在所有 Go 提交中击败了100.00%的用户

func maxValue(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
/*			sum := 0
			if j > 0 {
				sum = max(grid[i][j], grid[i][j-1]+grid[i][j])
			}
			if i > 0 {
				grid[i][j] = max(grid[i][j], grid[i-1][j]+grid[i][j])
			}
			grid[i][j] = max(sum, grid[i][j])*/
			if i==0 && j== 0{continue}
			sum1, sum2 := 0,0
			if j>0{
				sum1 = grid[i][j-1] + grid[i][j]
			}
			if i>0{
				sum2 = grid[i-1][j] + grid[i][j]
			}
			grid[i][j] = max(sum1, sum2)
		}
	}
	return grid[row-1][col-1]
}

/*我实现的官方题解：动态规划*/
// 执行用时：4 ms, 在所有 Go 提交中击败了98.93%的用户
// 内存消耗：3.8 MB, 在所有 Go 提交中击败了83.38%的用户
func maxValue1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	for i := 1; i < row; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < col; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			grid[i][j] += max(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[row-1][col-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1}}

	fmt.Println(maxValue1(nums))
}
