/**
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

 

示例 1：


输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]
示例 2：


输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 

提示：

m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1
 

进阶：

一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个仅使用常量空间的解决方案吗？

*/

package main

// 我的题解：哈希
/**
执行用时：12 ms, 在所有 Go 提交中击败了71.55%的用户
内存消耗：6.1 MB, 在所有 Go 提交中击败了89.11%的用户
*/

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m := len(matrix)    // 行
	n := len(matrix[0]) // 列
	hash1 := make(map[int]bool)
	hash2 := make(map[int]bool)

	// 记录应该 置零 的行和列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				hash1[i] = true
				hash2[j] = true
			}
		}
	}
	// 置 0
	for i := range hash1 {
		if hash1[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := range hash2 {
		if hash2[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

// 官方题解：置零的方式更简单一些
func setZeroes1(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix { // 遍历方式更简洁
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}
