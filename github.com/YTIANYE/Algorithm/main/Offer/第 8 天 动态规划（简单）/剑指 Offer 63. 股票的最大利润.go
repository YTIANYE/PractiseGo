/*
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

示例 1:
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

示例 2:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

限制：
0 <= 数组长度 <= 10^5

*/
package main

// 执行用时：4 ms, 在所有 Go 提交中击败了90.69%的用户
// 内存消耗：3 MB, 在所有 Go 提交中击败了100.00%的用户

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	profit := 0
	minnum := prices[0]
	maxnum := 0
	for _, val := range prices {
		if val < minnum { // 发现新的最小值
			minnum = val
			maxnum = 0
		} else if val > maxnum {
			maxnum = val
			if maxnum-minnum > profit {
				profit = maxnum - minnum
			}
		}
	}
	return profit
}
