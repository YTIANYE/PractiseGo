/*
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

 

示例 1：

输入：x = 4
输出：2
示例 2：

输入：x = 8
输出：2
解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
 

提示：

0 <= x <= 231 - 1

*/

package main

// 我的题解：暴力题解
/**
执行用时：24 ms, 在所有 Go 提交中击败了5.72%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了35.77%的用户
 */


func mySqrt1(x int) int {
	res := 0
	for {
		if res*res <= x && (res+1)*(res+1) > x {
			return res
		}
		res += 1

	}
	return 0
}


// 官方题解：二分查找

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r - l) / 2
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

