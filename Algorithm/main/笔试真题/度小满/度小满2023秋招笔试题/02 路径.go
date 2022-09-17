/**
路径
时间限制： 3000MS
内存限制： 589824KB
题目描述：
现在有一个n行m列的网格图。每个格子上有一个数字，为 +k 或者 -k。

初始你在位于左上方的第一行第一列，你的目标是走到位于右下方的第n行第m列。

现在限制你每一步只能向下或者向右走。显然，你有很多种实现目标的方法。

现在，我们定义路径的权值和为你选定的路径上所有格子上的数字的加和。

我们的问题是，你能否选出一条合适的路径，使得你获得恰好为x的权值和。

  1 <= n, m <= 10

  对全体数据保证：1 <= T <= 5, 0 <= k <= 10, 0 <= x <= 1000000



输入描述
第一行一个正整数T，表示总共有T组数据。

对于每组数据：

    第一行给出四个自然数n, m, k, x，以空格分开，表示给出的网格图是n行m列的，其中仅有+k和-k两种数字，所希望你获得的权值和为x。

    接下来将依次给出n行，每行m个整数，表示网格中的数字。

输出描述
对每一组数据输出一行yes或no表示答案，共输出T行。


样例输入
3
2 4 1 2
-1 -1 1 -1
-1 -1 -1 -1
1 2 1 1
1 -1
1 3 9 9
9 -9 9
样例输出
no
no
yes

提示
该样例中共三组数据。

第一组中，通过不同的路径可以组合出-3或者-5。两者均和2不同。因此是no。

第二组中，仅存在一种路径，其答案为0。和所要求的1不同，因此是no。

第三组中，仅存在一种路径，其答案为9。和所要求的9相同。因此是yes。
 */
package main

import "fmt"

// 100%
func main() {
	var T int
	fmt.Scan(&T)
	for ; T != 0; T-- {
		var n, m, k, x int
		fmt.Scan(&n, &m, &k, &x)
		tu := make([][]int, n)
		for i := range tu {
			tu[i] = make([]int, m)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Scan(&tu[i][j])
			}
		}
		//
		walk := func() bool {
			fen := make([][]map[int]bool, n)
			for i := range fen {
				fen[i] = make([]map[int]bool, m)
				for j := range fen[i] {
					fen[i][j] = make(map[int]bool)
				}
			}
			//
			fen[0][0][tu[0][0]] = true
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					hash := fen[i][j]
					if j-1 >= 0 {
						for num := range fen[i][j-1] {
							hash[num+tu[i][j]] = true
						}
					}
					if i-1 >= 0 {
						for num := range fen[i-1][j] {
							hash[num+tu[i][j]] = true
						}
					}
				}
			}
			//
			for num := range fen[n-1][m-1] {
				if num == x {
					return true
				}
			}
			return false
		}
		//
		if walk() {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
