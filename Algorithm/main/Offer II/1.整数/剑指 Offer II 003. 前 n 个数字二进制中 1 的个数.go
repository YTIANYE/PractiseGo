/**
给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。



示例 1:

输入: n = 2
输出: [0,1,1]
解释:
0 --> 0
1 --> 1
2 --> 10
示例 2:

输入: n = 5
输出: [0,1,1,2,1,2]
解释:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101


说明 :

0 <= n <= 105


进阶:
给出时间复杂度为 O(n*sizeof(integer)) 的解答非常容易。但你可以在线性时间 O(n) 内用一趟扫描做到吗？
要求算法的空间复杂度为 O(n) 。
你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount ）来执行此操作。

*/
package main

import "fmt"

// 我的题解：动态规划（类似最高有效位）
/**
执行用时：4 ms, 在所有 Go 提交中击败了85.77%的用户
内存消耗：4.4 MB, 在所有 Go 提交中击败了63.86%的用户
*/

/**
0	0
1	1
2	10
3	11
4	100
5	101
6	110
7	111
8	1000
*/

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	list := make([]int, n+1)
	list[0], list[1] = 0, 1
	l := 2
	j := 0
	i := 2
	for i <= n {
		if j == l {
			l = i
			j = 0
		}
		list[i] = 1 + list[j]
		j++
		i++
	}
	return list
}

// 官方题解1：Brian Kernighan
func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits1(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

// 官方题解2 ： 动态规划：最高有效位
func countBits2(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}

// 官方题解3：动态规划：最低有效位
func countBits3(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

// 官方题解4：动态规划——最低设置位
func countBits4(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}




func main() {
	fmt.Println(countBits(5))
}
