/**
题目：长度为n的数组，有第一个到最后一个经过的最小累计代价是多少？
每一步只能走到下一个或者下下个位置上，每一步的代价为两个位置的绝对值
 */

package main



import (
"fmt"
)

func main(){
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:=0;i<n;i++{
		fmt.Scan(&nums[i])
	}

	// 计算
	dp := make([]int, n)
	dp[0] = 0
	for i:=1;i<n;i++{
		dp[i] = abs(nums[i], nums[i-1]) + dp[i-1]
		if i>=2{
			temp := abs(nums[i], nums[i-2]) + dp[i-2]
			dp[i] = min(dp[i], temp)
		}
	}
	fmt.Println(dp[n-1])
}

func abs(a, b int ) int{
	if a > b {
		return a-b
	}
	return b-a
}

func min(a, b int ) int{
	if a > b {
		return b
	}
	return a
}