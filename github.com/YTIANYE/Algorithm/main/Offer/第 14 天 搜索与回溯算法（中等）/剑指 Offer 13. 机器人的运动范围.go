/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。
但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

示例 1：
输入：m = 2, n = 3, k = 1
输出：3
示例 2：
输入：m = 3, n = 1, k = 0
输出：1
提示：

1 <= n,m <= 100
0 <= k<= 20
*/

package main

import (
	"fmt"
)

// 我的题解：深度遍历
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：3.2 MB, 在所有 Go 提交中击败了17.54%的用户
*/
func movingCount1(m int, n int, k int) int {
	// 地图初始化
	var board [][]int
	for i := 0; i < m; i++ {
		var nums []int
		for j := 0; j < n; j++ {
			nums = append(nums, 0)
		}
		board = append(board, nums)
	}
	// 走
	find(m, n, 0, 0, k, board)
	// 统计足迹
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != 0 {
				res++
			}
		}
		fmt.Println(board[i])
	}
	return res
}

// 深度遍历
func find(m, n, i, j, k int, board [][]int) {
	if i >= m || i < 0 || j >= n || j < 0 {
		return
	}
	if isright(i, j, k) {
		board[i][j]++
		if board[i][j] == 1 { // 加了这个判断极大缩短了时间
			find(m, n, i, j+1, k, board)
			find(m, n, i+1, j, k, board)
		}

	}
}

// 位数之和是否满足条件
func isright(num1, num2, k int) bool {
	sum := 0
	for num1 != 0 {
		sum += num1 % 10
		num1 /= 10
	}
	for num2 != 0 {
		sum += num2 % 10
		num2 /= 10
	}
	return sum <= k
}

// 我实现的官方题解：层次遍历
/*
执行用时：4 ms, 在所有 Go 提交中击败了11.17%的用户
内存消耗：4.5 MB, 在所有 Go 提交中击败了5.10%的用户
*/
func movingCount(m int, n int, k int) int {
	queue := [][2]int{{0, 0}}
	// set := make(map[string]int)
	set := make(map[[2]int]int)
	for len(queue) != 0 {
		nums := queue[0]
		queue = queue[1:]
		i := nums[0]
		j := nums[1]
		//char := strconv.Itoa(i) + " " + strconv.Itoa(j)
		char := [2]int{i,j}
		if i < m && j < n && set[char] == 0 && isright(i, j, k) {
			set[char] = 1
			queue = append(queue, [2]int{i + 1, j})
			queue = append(queue, [2]int{i, j + 1})
		}
	}
	return len(set)
}

// 我实现的官方题解：递推
func movingCount3(m int, n int, k int) int {
	set := make(map[[2]int]int)
	set[[2]int{0, 0}] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (set[[2]int{i - 1, j}] == 1 || set[[2]int{i, j - 1}] == 1) && isright(i, j, k) {
				set[[2]int{i, j}] = 1
			}
		}
	}
	return len(set)
}

func main() {
	fmt.Println(movingCount(36, 11, 15))
}
