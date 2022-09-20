/**
Leetcode 72

*/
package main

import "fmt"

// 100%
func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(fun(s, t))
}

func fun(s, t string) int {
	n, m := len(s), len(t)
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	//
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	//
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}
	//
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			l := dp[i-1][j] + 1
			d := dp[i][j-1] + 1
			ld := dp[i-1][j-1]
			if s[i-1] != t[j-1] {
				ld++
			}
			dp[i][j] = min(ld, min(l, d))
		}
	}
	//
	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
