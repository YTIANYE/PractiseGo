/**
题目：
一个数组a,每次删除数组中所有非质数的位上的数（位 从1开始，1...n）
当数组a的长度为1时，返回剩余的最后一个数
1 不是质数
*/

package main

import "math"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param a int整型一维数组
 * @return int整型
 */

// [3,1,1,4,5,6]
// 5

// 1 1 5
// 1 5
// 5

// 通过样例 100%
func getNumber(a []int) int {
	// write code here
	n := len(a)
	hash := zhishu(n)

	for len(a) != 1 {
		for i := len(a); i >= 1; i-- { // 注意倒序删除，避免删除某个位置造成长度变化、下标变化
			if !hash[i] { // 非质数位删除
				j := i - 1
				if j == len(a)-1 {
					a = a[:j]
				} else {
					a = append(a[:j], a[j+1:]...)
				}
			}
		}
	}
	return a[0]
}

func zhishu(n int) map[int]bool {
	hash := make(map[int]bool)
	for i := 2; i <= n; i++ {
		if iszhishu(i) {
			hash[i] = true
		}
	}
	return hash
}

func iszhishu(n int) bool {
	end := int(math.Pow(float64(n), 0.5))
	for i := 2; i <= end; i++ { // i <= n/2 的话只能通过76%的样例
		if n%i == 0 {
			return false
		}
	}
	return true
}
