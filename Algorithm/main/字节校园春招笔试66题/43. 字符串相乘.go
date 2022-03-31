/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。


示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
 

提示：
1 <= num1.length, num2.length <= 200
num1 和 num2 只能由数字组成。
num1 和 num2 都不包含任何前导零，除了数字0本身。

*/

package main

import "strconv"

// 我实现的官方题解：逐一相乘
/*
执行用时：4 ms, 在所有 Go 提交中击败了55.96%的用户
内存消耗：2.9 MB, 在所有 Go 提交中击败了71.76%的用户
*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	res := make([]int, n+m)

	// 0123
	//  012

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			c, d := cheng(num1[i], num2[j])
			res[i+j+1] += d
			jinwei := res[i+j+1] / 10
			res[i+j+1] %= 10

			res[i+j] += c + jinwei
			jinwei2 := res[i+j] / 10
			if jinwei2 != 0{
				res[i+j-1] += jinwei2
			}
			res[i+j] %= 10
		}
	}

	jieguo := ""
	for i := range res {
		jieguo += strconv.Itoa(res[i])
	}
	if jieguo[0] == '0'{
		jieguo = jieguo[1:]
	}
	return jieguo
}

func cheng(num1, num2 uint8) (int, int) {
	a, b := int(num1-'0'), int(num2-'0')
	res := a * b
	d := res % 10
	c := res / 10
	return c, d
}
