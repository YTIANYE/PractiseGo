/**
给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。



注意：

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231−1]。本题中，如果除法结果溢出，则返回 231 − 1


示例 1：

输入：a = 15, b = 2
输出：7
解释：15/2 = truncate(7.5) = 7
示例 2：

输入：a = 7, b = -3
输出：-2
解释：7/-3 = truncate(-2.33333..) = -2
示例 3：

输入：a = 0, b = 1
输出：0
示例 4：

输入：a = 1, b = 1
输出：1


提示:

-231 <= a, b <= 231 - 1
b != 0


*/

package main

import (
	"fmt"
	"math"
)

/**
执行用时：4 ms, 在所有 Go 提交中击败了46.06%的用户
内存消耗：2.4 MB, 在所有 Go 提交中击败了6.11%的用户
*/

// 我的题解:保留进制数
func divide1(a int, b int) int {
	// 唯一的溢出情况
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	// 全部换成负数
	fuhao := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		fuhao = -1 // 结果为负数
	}
	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}
	// 存储b的0 1 2 4 8 。。。次方数
	nums := make(map[int]int)
	nums[0] = -1
	nums[1] = b
	for i := 1; nums[i] >= a-nums[i]; i += i {
		nums[i+i] = nums[i] + nums[i]
	}
	// 计算b进制数下的 a
	//list := []int{}
	//var jisuan func(num int)
	//jisuan = func(num int) {
	//	if num > b  {
	//		return
	//	}
	//	i:= 1
	//	for nums[i+i] != 0 && nums[i+i] >= num{
	//		i+=i
	//	}
	//	list = append(list, i)
	//	jisuan(num - nums[i])
	//}
	//jisuan(a)
	//
	//// 计算结果
	//res := 0
	//for _, num := range list {
	//	res += num
	//}

	// 简化过程
	num := a
	res := 0
	for num <= b {
		i := 1
		for nums[i+i] != 0 && nums[i+i] >= num {
			i += i
		}
		res += i
		num -= nums[i]
	}

	return res * fuhao
}

// 精选题解：不保留进制数[节省空间]
/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了66.67%的用户
 */
func divide(a int, b int) int {
	// 防止溢出
	if a == math.MinInt32  {
		if b == -1 {
			return math.MaxInt32
		}
		if b == 1 {
			return math.MinInt32
		}
	}
	// 符号转化
	fuhao := 1
	if a < 0 {
		fuhao = -fuhao
	}
	if b < 0 {
		fuhao = -fuhao
	}
	// 放置转正后溢出
	res := 0
	if a == math.MinInt32 {
		if b > 0 {
			a += b
		}else {
			a -= b
		}
		res++
	}
	a, b = abs(a), abs(b)  // 注意在a b 值未发生改变之前，记录符号
	// 计算
	for i := 31; i >= 0; i-- {
		if (a >> i) >= b {
			a -= b << i
			res += 1 << i
		}
	}
	return res * fuhao
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// a, b := 15, 2
	// a, b := -2147483648, -1
	// a, b := -2147483648, 1
	// a, b := 1, 1
	a, b := math.MinInt32, math.MinInt32
	fmt.Println(divide(a, b))
}
