/**
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2



示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2


提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。


*/
package main

import "math"

/**
特例：

-2147483648
-1

*/

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了62.96%的用户
*/

// 我的题解：不保存进制（保存进制的修改版）
func divide(dividend int, divisor int) int {
	// 特殊例子
	if dividend == -(1<<31) && divisor == -1 {
		return 1<<31 - 1
	}
	// 取符号、除数被除数换成正数
	tag := 1
	if dividend < 0 {
		tag = -tag
		dividend = -dividend // 注意变换符号
	}
	if divisor < 0 {
		tag = -tag
		divisor = -divisor // 注意变换符号
	}
	//
	temp := divisor
	i := 1
	for temp < dividend - temp { // 注意32位数，不能越界  不是temp <= dividend
		temp <<= 1
		i <<= 1
	}
	//
	res := 0
	for dividend >= divisor {
		for temp > dividend {
			temp >>= 1
			i >>= 1
		}
		res += i
		dividend -= temp
	}
	// 符号
	if tag < 0 {
		res = -res
	}
	return res

}

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了13.80%的用户
*/

// 我的题解：保存进制
func divide1(dividend int, divisor int) int {

	// 特殊例子
	if dividend == -(1<<31) && divisor == -1 {
		return 1<<31 - 1
	}
	// 取符号、除数被除数换成正数
	tag := 1
	if dividend < 0 {
		tag = -tag
		dividend = -dividend // 注意变换符号
	}
	if divisor < 0 {
		tag = -tag
		divisor = -divisor // 注意变换符号
	}
	// 保存divisor的 1 2 4 8 ...倍的数
	nums := []int{divisor} // 2^0 = 1 ; 2^1 = 2 ; 2^2 = 4 ; 2^3 = 8 ;
	for temp := divisor<<1 ; temp < dividend - temp; temp <<= 1 { // 注意判断条件不要溢出
		nums = append(nums, temp)
	}
	// 累加
	res := 0
	i := len(nums) - 1
	for dividend >= divisor {
		for nums[i] > dividend && i >= 0 {
			i--
		}
		res += 1 << i
		dividend -= nums[i]
	}
	// 符号
	if tag < 0 {
		res = -res
	}
	return res
}

// 官方题解：二分查找
// 快速乘
// x 和 y 是负数，z 是正数
// 判断 z * y >= x 是否成立
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 { // 不能使用除法
		if z&1 > 0 {
			// 需要保证 result + add >= x
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 {
			// 需要保证 add + add >= x
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

func divide2(dividend, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1 // 注意溢出，并且不能使用除法
		if quickAdd(divisor, mid, dividend) {
			ans = mid
			if mid == math.MaxInt32 { // 注意溢出
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans
}

