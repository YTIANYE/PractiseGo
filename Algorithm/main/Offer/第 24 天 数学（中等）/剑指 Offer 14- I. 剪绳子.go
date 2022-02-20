/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

示例 1：

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1
示例2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36
提示：
2 <= n <= 58

*/

package main

import (
	"fmt"
	"math"
)

/*
事例
1	1
2	1
3	2
4	4
5	6
6	9
7	12
8	18
9	27
10	36
*/

// 我的题解：动态规划
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了6.05%的用户
*/

func cuttingRope1(n int) int {
	maxs := map[int]int{1: 1, 2: 1, 3: 2, 4: 4}
	if maxs[n] != 0 {
		return maxs[n]
	}
	maxs = cut(n, maxs)
	return maxs[n]
}

func cut(n int, maxs map[int]int) map[int]int {
	maxNum := 0
	for i := 1; i <= n/2; i++ {
		if maxs[i] == 0 {
			maxs = cut(i, maxs)
		}
		if maxs[n-i] == 0 {
			maxs = cut(n-i, maxs)
		}
		temp := max(maxs[i], i) * max(maxs[n-i], n-i)
		maxNum = max(temp, maxNum)
	}
	maxs[n] = maxNum
	return maxs
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 精选题解：动态规划
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了59.08%的用户
*/
func cuttingRope(n int) int {
	if n<=3{
		return max(1, n-1)
	}

	maxs := map[int]int{1: 1, 2: 2, 3: 3}

	for k := 4; k <= n; k++ {
		for i := 1; i <= k/2; i++ {
			maxs[k] = max(maxs[k], maxs[i] * maxs[k-i])
		}
	}
	return maxs[n]
}

// 我实现的精选题解：数学方法
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了41.40%的用户
*/
func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return int(math.Pow(3.0, float64(a)))
	} else if b == 1 {
		return int(math.Pow(3.0, float64(a-1)) * 4)
	} else { // 2
		return int(math.Pow(3.0, float64(a)) * 2)
	}
}

func main() {
	n := 20
	for i := 1; i < n; i++ {
		fmt.Println(i, cuttingRope(i))
	}
}
