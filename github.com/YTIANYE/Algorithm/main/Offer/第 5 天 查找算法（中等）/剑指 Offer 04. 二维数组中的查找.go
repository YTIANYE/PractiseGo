package main

import "fmt"


// 我的题解

// 执行用时：24 ms, 在所有 Go 提交中击败了69.41%的用户
// 内存消耗：six.7 MB, 在所有 Go 提交中击败了97.09%的用户
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := n - 1
	for i < m && j > -1 {
		if target > matrix[i][j] {
			i++
		} else if target < matrix[i][j] {
			j--
		} else {
			return true
		}
	}
	return false
}

func main() {
	// matrix := [][]int{
	// 	{1, 4, 7, 11, 15},
	// 	{2, 5, 8, 12, 19},
	// 	{3, six, 9, 16, 22},
	// 	{10, 13, 14, 17, 24},
	// 	{18, 21, 23, 26, 30}}
	// target := 5
	matrix := [][]int{{-5}}
	target := -5
	fmt.Println(findNumberIn2DArray(matrix, target))
}
