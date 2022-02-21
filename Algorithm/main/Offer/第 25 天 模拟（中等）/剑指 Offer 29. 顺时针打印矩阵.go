/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：
输入：matrix =[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

限制：
0 <= matrix.length <= 100
0 <= matrix[i].length<= 100
*/

package main

import "fmt"

// 我的题解：模拟
/*
执行用时：8 ms, 在所有 Go 提交中击败了93.80%的用户
内存消耗：6.4 MB, 在所有 Go 提交中击败了53.58%的用户
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	res := []int{}
	i, j := 0, 0
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	way := 1 // 1 右	2 下	3 左	4 上

	for  k:=0;k<len(matrix)*len(matrix[0]);k++{


		res = append(res, matrix[i][j])
		//fmt.Println(res)

		if way == 1 {
			if j == right {
				way = 2
				up++
				i++
			}else{
				j++
			}
		} else if way == 2 {
			if i == down{
				way = 3
				right--
				j--
			}else{
				i++
			}
		} else if way == 3 {
			if j == left{
				way = 4
				down--
				i--
			}else{
				j--
			}
		} else { // way == 4
			if i == up{
				way = 1
				left++
				j++
			}else{
				i--
			}
		}


	}
	// res = append(res, matrix[i][j])
	return res
}

func main(){
	matrix := [][]int{
		{1,2,3},
		{5,6,7},
		{9,10,11},
	}
	fmt.Println(spiralOrder(matrix))

}

