/**
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

 

示例 1：

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
 

提示：

1 <= prices.length <= 105
0 <= prices[i] <= 104

*/

package main

import (
	"fmt"
	"math"
)

// 动态规划：我的题解
/**
执行用时：148 ms, 在所有 Go 提交中击败了11.08%的用户
内存消耗：7.8 MB, 在所有 Go 提交中击败了74.36%的用户
*/

func maxProfit(prices []int) int {

	// 动态规划
	buy := math.MaxInt32 // 记录当前最应该入手买入的点（即当前位置之前最小值，作为买入点）
	profit := 0          // 记录当前位置之前，最大的利润
	for _, price := range prices {
		profit = max121(profit, price-buy)
		buy = min121(buy, price)
	}

	return profit
}

func max121(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min121(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// nums := []int{7, 1, 5, 3, 6, 4}
	nums := []int{7, 6, 4, 3, 1}
	res := maxProfit(nums)
	fmt.Println(res)
}
