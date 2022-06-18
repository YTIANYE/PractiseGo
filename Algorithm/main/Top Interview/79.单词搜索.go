/**
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



示例 1：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true
示例 3：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false


提示：

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成


进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
*/
package main

import "fmt"

// 我的题解：递归+剪枝
/**
执行用时：88 ms, 在所有 Go 提交中击败了51.16%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了19.46%的用户
*/

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	tagmap := make([][]bool, m)
	for i := range tagmap {
		tagmap[i] = make([]bool, n)
	}

	var dfs func(int, int, int) bool
	dfs = func(x, y, index int) bool {
		if index == len(word) {
			return true
		}
		if x >= m || x < 0 || y >= n || y < 0 || tagmap[x][y] == true || board[x][y] != word[index] {
			return false
		}

		tagmap[x][y] = true
		if dfs(x, y+1, index+1) || dfs(x, y-1, index+1) || dfs(x+1, y, index+1) || dfs(x-1, y, index+1) {
			return true
		}
		tagmap[x][y] = false
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false

}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))

}
