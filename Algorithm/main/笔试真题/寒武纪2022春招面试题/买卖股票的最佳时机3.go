/*
买卖股票的最好时机（3）
假设你有一个数组，其中第i个元素是某只股票在第i天的价格。
设计一个算法来求最大的利润。你最多可以进行两次交易。
注意：
你不能同时进行多个交易（即：你必须在再次购买之前出售之前买的股票）

实例：
输入：1,4,2
输出：3

实例：
输入：2,4,1
输出：2



*/
package main

import "fmt"

// 官方题解：动态规划
/**
执行用时：96 ms, 在所有 Go 提交中击败了84.24%的用户
内存消耗：8.5 MB, 在所有 Go 提交中击败了98.65%的用户
*/
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 4, 2, 1, 5}
	fmt.Println(maxProfit(nums))
}
