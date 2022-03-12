/*

给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0

提示：

-231 <= x <= 231 - 1

*/

package main

import "math"

// 我的题解：
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了86.08%的用户
*/
func reverse3(x int) int {
	sigh := 1 // 符号位
	if x < 0 {
		sigh = -1
		x = -x
	}
	maxInt := math.MaxInt32

	res := 0
	for x != 0 {
		yushu := x % 10
		x /= 10
		if res > maxInt/10 {
			return 0
		} else if res == int(maxInt/10) && ((sigh == 1 && yushu > 8) || (sigh == -1 && yushu > 7)) {
			return 0
		}
		res = res*10 + yushu
	}

	return res * sigh

}
