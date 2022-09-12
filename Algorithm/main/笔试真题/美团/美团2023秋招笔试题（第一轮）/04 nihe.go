/**
拟合
时间限制： 3000MS
内存限制： 589824KB
题目描述：
 小团生日收到妈妈送的两个一模一样的数列作为礼物！他很开心的把玩，不过不小心没拿稳将数列摔坏了！
现在他手上的两个数列分别为A和B，长度分别为n和m。小团很想再次让这两个数列变得一样。
他现在能做两种操作，操作一是将一个选定数列中的某一个数a改成数b，这会花费|b-a|的时间，
操作二是选择一个数列中某个数a，将它从数列中丢掉，花费|a|的时间。
小团想知道，他最少能以多少时间将这两个数列变得再次相同！



输入描述
第一行两个空格隔开的正整数n和m，分别表示数列A和B的长度。

接下来一行n个整数，分别为A1 A2...An

接下来一行m个整数，分别为B1 B2...Bm

对于所有数据，1≤n,m≤2000， |Ai|,|Bi|≤10000

输出描述
输出一行一个整数，表示最少花费时间，来使得两个数列相同。


样例输入
1 1
-9821
7742
样例输出
17563

提示
可以选择两次第二种操作，消除数列A的第一个数和数列B的第一个数，需要花费9821+7742=17563的时间

也可以选择一次第一种操作，将数列A的第一个数改成数列B的第一个数，也是需要花费9821+7742=17563的时间

所以答案为17563

规则
*/
package main

import "fmt"

// 未验证
func main() {
	// 输入
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n) // 注意输入n之后在初始化
	b := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	//
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + abs(a[i-1])
	}
	for j := 1; j <= m; j++ {
		dp[0][j] = dp[0][j-1] + abs(b[j-1])
	}

	//
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = min(min(
				dp[i-1][j]+abs(a[i-1]),
				dp[i][j-1]+abs(b[j-1])),
				dp[i-1][j-1]+abs(a[i-1]-b[j-1]))
		}
	}
	//
	res := dp[n][m]
	fmt.Println(res)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
