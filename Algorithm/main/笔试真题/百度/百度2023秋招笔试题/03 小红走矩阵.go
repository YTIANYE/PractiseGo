/**
注意：这题用动态规划不合适，动态规划适合只可以向右和下方向走，本题可能往回走

小红从左上角走到右下角，
每个格子有 r e d 三种字符，
小红不能走的三种情况是 r->d; d -> e; e->r;
问小红到达终点至少走几步，不能到返回 -1
 */
package main

import (
	"fmt"
	"math"
)

// 100%
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	tu := make([]string, n)
	for i:=0;i<n;i++ {
		fmt.Scan(&tu[i])
	}

	dp := make([][]int, n)
	for i:= range dp {
		dp[i] = make([]int, m)
	}
	for i:=0;i<n;i++ {
		for j:=0;j<m;j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	hash := make(map[string]bool)
	hash["re"] = true
	hash["rr"] = true
	hash["ee"] = true
	hash["ed"] = true
	hash["dr"] = true
	hash["dd"] = true
	//
	zou := func(x, y, i, j int) bool  {
		if i <0 || i >=n || j < 0 || j >= m {
			return false
		}
		key := string(tu[x][y]) + string(tu[i][j])
		if hash[key] == false {
			return false
		}

		if dp[x][y] + 1 < dp[i][j] {  // 注意是<号，否则只能通过30%
			dp[i][j] = dp[x][y] + 1
			return true
		}
		return false
	}

	// 广度优先的方式
	queue := [][]int{}
	queue = append(queue, []int{0,0})
	for len(queue) != 0 {
		tq := queue
		queue = [][]int{}
		for len(tq) != 0 {
			dian := tq[0]
			if len(tq) == 1 {
				tq = [][]int{}
			}else {
				tq = tq[1:]
			}
			x, y := dian[0], dian[1]

			next := [][]int{
				{x+1, y},
				{x-1, y},
				{x, y+1},
				{x, y-1},
			}
			for k := range next {
				i, j := next[k][0], next[k][1]
				if zou(x,y, i, j ) {
					tdian := []int{i,j}
					queue = append(queue, tdian)
				}
			}
		}
	}
	//
	if dp[n-1][m-1] == math.MaxInt32{
		fmt.Println(-1)
	}else {
		fmt.Println(dp[n-1][m-1])
	}
}
