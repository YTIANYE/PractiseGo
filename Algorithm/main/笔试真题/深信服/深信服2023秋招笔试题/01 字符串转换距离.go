/**
LeetCode 72.编辑距离

一个字符串转化成另外一个字符串
可以进行的操作是
删除一个
添加一个
修改一个

问至少需要多少次操作？
 */


package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param a string字符串 字符串1
 * @param b string字符串 字符串2
 * @return int整型
 */

// 100%
func minDistance( a string ,  b string ) int {
	// write code here
	n , m := len(a), len(b)
	if m == 0 {
		return n
	}
	if n == 0{
		return m
	}
	//
	dp := make([][]int, n+1)
	for i:= range dp {
		dp[i] = make([]int, m+1)
	}
	//
	for i:=0;i<=n;i++ {
		dp[i][0] = i
	}
	for i := 0; i<=m;i++ {
		dp[0][i] = i
	}
	//
	for i:= 1;i<=n;i++ {
		for j:=1;j<=m;j++ {
			zuo := dp[i-1][j] + 1
			xia := dp[i][j-1] + 1
			zuoxia := dp[i-1][j-1]
			if a[i-1] != b[j-1] {
				zuoxia += 1
			}
			dp[i][j] = min1(zuo, min1(zuoxia, xia))
		}
	}

	//
	return dp[n][m]
}

func min1 (a, b int) int {
	if a < b {
		return a
	}
	return b
}