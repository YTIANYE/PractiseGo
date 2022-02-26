/*
把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。

你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

示例 1:
输入: 1
输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
示例2:
输入: 2
输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]

限制：
1 <= n <= 11
*/
package main

import (
	"fmt"
	"math"
)

// 我的题解：动态规划
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了25.29%的用户
*/
func dicesProbability(n int) []float64 {

	hashMap := map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1}

	for i := 1; i < n; i++ { // i:色子的个数-1
		newHash := make(map[int]int)
		for k := range hashMap { // k：此时有的数
			for j := 1; j <= 6; j++ {
				index := j + k
				newHash[index] += hashMap[k] // 关键的一步
			}
		}
		hashMap = newHash
	}

	res := []float64{}
	for i := n; i <= 6*n; i++ {
		res = append(res, float64(hashMap[i])/math.Pow(6.0, float64(n)))
		// fmt.Println(hashMap[i])
	}
	return res
}

func main() {
	n := 3
	fmt.Println(dicesProbability(n))
}
