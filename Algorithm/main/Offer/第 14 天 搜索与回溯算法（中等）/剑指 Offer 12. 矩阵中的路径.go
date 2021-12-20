/*
给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。

示例 1：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：
输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

提示：

1 <= board.length <= 200
1 <= board[i].length <= 200
board 和 word 仅由大小写英文字母组成
*/
package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m := len(board) // 行
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				var indexs [][]int
				// indexs = append(indexs, []int{i, j})
				if findword(board, i, j, 0, word, indexs) {
					return true
				}
			}
		}
	}
	return false
}

func findword(board [][]byte, i int, j int, k int, word string, indexs [][]int) bool {
	if k == len(word) {
		return true
	}

	m := len(board) // 行
	n := len(board[0])
	if i < 0 || i >= m || j < 0 || j >= n { // 注意 i j 都是索引， 不可以与m n相等
		return false
	}

	if board[i][j] == word[k] && !isinindexs(i, j, indexs) {
		indexs = append(indexs, []int{i, j})
		return findword(board, i, j+1, k+1, word, indexs) ||
			findword(board, i, j-1, k+1, word, indexs) ||
			findword(board, i+1, j, k+1, word, indexs) ||
			findword(board, i-1, j, k+1, word, indexs)
	} else {
		return false
	}

}

// 判断当前board[i][j]是否已经走过
// 【精选题解中遍历过的字符换成“ ”，继续遍历，之后会恢复“ ”,可以节省存储空间 】
func isinindexs(i, j int, indexs [][]int) bool {
	for _, index := range indexs {
		if index[0] == i && index[1] == j {
			return true
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'a', 'b'},
		{'c', 'd'}}
	word := "abcd"
	//
	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// word := "ABCCED"
	fmt.Println(exist(board, word))
}
