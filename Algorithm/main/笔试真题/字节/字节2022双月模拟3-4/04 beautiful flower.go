package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	p := make([][3]int, 200000)
	dp := make([][3]int, 200000)
	for i := 1; i <= n; i++ {
		Fscan(in, &p[i][0], &p[i][1], &p[i][2])
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				dp[i][j] = max(dp[i][j], dp[i-1][k]+p[i][j])
			}
		}
	}
	ans := 0
	for i := 0; i < 3; i++ {
		ans = max(ans, dp[n][i])
	}
	Fprintln(out, ans)
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
