/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。

 

示例 1：

输入：num1 = "11", num2 = "123"
输出："134"
示例 2：

输入：num1 = "456", num2 = "77"
输出："533"
示例 3：

输入：num1 = "0", num2 = "0"
输出："0"

提示：

1 <= num1.length, num2.length <= 104
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零

*/

package main

import "strconv"

/*
执行用时：8 ms, 在所有 Go 提交中击败了48.78%的用户
内存消耗：6.8 MB, 在所有 Go 提交中击败了60.57%的用户
*/

func addStrings(num1 string, num2 string) string {
	res := ""
	var jinwei uint8
	jinwei = 0

	i := len(num1)-1
	j := len(num2)-1
	for i >= 0 || j >= 0 {
		temp := jinwei
		if i >= 0 {
			temp += num1[i] - '0'
		}
		if j >= 0 {
			temp += num2[j] - '0'
		}
		yushu := temp % 10
		jinwei = temp / 10
		res = strconv.Itoa(int(yushu)) + res
		i--
		j--
	}

	if jinwei >0{
		res = strconv.Itoa(int(jinwei)) + res
	}
	return res
}
