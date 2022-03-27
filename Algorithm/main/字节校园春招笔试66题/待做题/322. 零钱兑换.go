/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。


示例1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0

提示：

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

*/

package main

import (
	"fmt"
	"math"
	"sort"
)

// 我的题解：正向 dfs + 剪枝 超时
func coinChange1(coins []int, amount int) int {
	sort.Ints(coins)
	res := math.MaxInt32

	var dfs func([]int, int, int)
	dfs = func(nums []int, release int, count int) {
		if release == 0 {
			res = min(res, count)
		} else if release > 0 {
			for i := len(nums) - 1; i >= 0; i-- {
				dfs(nums, release-nums[i], count+1)
				nums = nums[:len(nums)-1]
			}
		}
	}

	dfs(coins, amount, 0)
	if res != math.MaxInt32 {
		return res
	} else {
		return -1
	}
}

// 官方题解：反向 dfs 超时
func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	sort.Reverse(sort.IntSlice(coins))
	// 增加一个记录表
	hashCoin := make(map[int]int)

	var dfs func(int) int
	dfs = func(release int) int {
		if release == 0 {
			return 0
		} else if release < 0 {
			return -1
		} else if res, ok := hashCoin[release]; ok && hashCoin[release] != -1 {
			return res
		}

		minCount := math.MaxInt32
		for _, coin := range coins {

			res := dfs(release - coin)
			if res >= 0 && res < minCount {
				minCount = res + 1
			}
		}

		if minCount < math.MaxInt32 {
			hashCoin[release] = minCount
			return minCount
		}
		hashCoin[release] = -1
		return -1
	}

	return dfs(amount)
}

// 官方题解：动态规划
/*
执行用时：8 ms, 在所有 Go 提交中击败了91.82%的用户
内存消耗：6.7 MB, 在所有 Go 提交中击败了16.08%的用户
*/
func coinChange(coins []int, amount int) int {
	dp := []int{}
	for i := 0; i <= amount; i++ {
		dp = append(dp, math.MaxInt32)
	}
	dp[0] = 0

	for _, coin := range coins {
		for x := coin; x <= amount; x++ {
			dp[x] = min(dp[x], dp[x-coin]+1)
		}
	}

	if dp[amount] != math.MaxInt32 {
		return dp[amount]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{186, 419, 83, 408}

	amount := 6249
	fmt.Println(coinChange(coins, amount))
}
