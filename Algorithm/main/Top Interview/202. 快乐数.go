/**
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

 

示例 1：

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
示例 2：

输入：n = 2
输出：false
 

提示：

1 <= n <= 231 - 1

*/

package main

import "math"

// 我的题解：哈希
/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了25.58%的用户
*/
func isHappy(n int) bool {
	hash := make(map[int]bool)
	hash[n] = true
	for n != 1 {
		n = happy(n)
		if hash[n] { // 循环了
			return false
		}
		hash[n] = true
	}
	return true
}

func happy(n int) int {
	res := 0.0
	for n != 0 {
		res += math.Pow(float64(n%10), 2)
		n /= 10
	}
	return int(res)
}
