/**
漂亮花
https://atcoder.jp/contests/dp/tasks/dp_q

从n朵花中按顺序挑选，高度需要依次递增，求挑出的所有花的奖励值最大是多少？


输入
n 朵花，
每朵花的高度
每朵花的奖励值
输出
最大奖励值
 */

package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param N int整型 花的数量
 * @param h int整型一维数组 花的高度
 * @param a int整型一维数组 花的漂亮值
 * @return long长整型
 */

// 我的题解：动态规划
// 9,[4,2,5,8,3,6,1,7,9],[6,8,8,4,6,3,5,7,5]
// 9,[ ,2,5, , ,6, ,7,9],[ ,8,8, , ,3, ,7,5]

func salesman( N int ,  h []int ,  a []int ) int64 {
	dp := make([]int64, N)
	for i:=0;i<N;i++{
		dp[i] = int64(a[i])
		for j:=0;j<i;j++{
			if h[i] > h[j] {
				dp[i] = max(dp[i], dp[j] + int64(a[i]))
			}
		}
	}
	return dp[N-1]
}

func max(a, b int64 ) int64 {
	if a > b {
		return a
	}
	return b
}


func main() {
	N := 9
	h := []int{4,2,5,8,3,6,1,7,9}
	a := []int{6,8,8,4,6,3,5,7,5}
	fmt.Println(salesman(N, h, a)) // 31
}
