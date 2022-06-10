/**
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn ）。

 

示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25
 

提示：

-100.0 < x < 100.0
-231 <= n <= 231-1
-104 <= xn <= 104
*/

package main

import "fmt"

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了99.92%的用户
 */

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / zhengPow(x, -n)
	}
	return zhengPow(x, n)
}

// 10   1010    x2 + x8
func zhengPow(x float64, n int) float64 {

	temp := x
	res := 1.0
	for n != 0 {
		yushu := n % 2
		n /= 2
		if yushu != 0{
			res *= temp
		}
		temp  = temp * temp
	}

	return res
}

func main(){
	n := 10
	x := 2.0
	fmt.Println(myPow(x, n))
}
