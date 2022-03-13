/*

给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

 

示例 1：


输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：


输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100

*/

package main

// 我的题解：模拟
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了63.00%的用户
*/

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)    // 行
	n := len(matrix[0]) // 列
	if m == 0 || n == 0 {
		return []int{}
	}
	sum := m * n
	count := 0

	wallUp, wallRight, wallDown, wallLeft := -1, n, m, -1
	fangxiang := 0 // 0:"right", 1:"down", 2:"left", 3:"up"
	i, j := 0, 0
	res := []int{}

	for count < sum {
		res = append(res, matrix[i][j])
		count++
		if fangxiang == 0 {
			if j == wallRight-1 { // 注意改变方向和修改墙
				i++
				fangxiang = (fangxiang + 1) % 4
				wallUp++
			} else {
				j++
			}
		} else if fangxiang == 1 {
			if i == wallDown-1 {
				j--
				fangxiang = (fangxiang + 1) % 4
				wallRight--
			} else {
				i++
			}
		} else if fangxiang == 2 {
			if j == wallLeft+1 {
				i--
				fangxiang = (fangxiang + 1) % 4
				wallDown--
			} else {
				j--
			}
		} else { // fangxiang == 3
			if i == wallUp+1 {
				j++
				fangxiang = (fangxiang + 1) % 4
				wallLeft++
			} else {
				i--
			}
		}
	}

	return res
}
