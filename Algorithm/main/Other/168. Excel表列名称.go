/**
给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。

例如：

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...


示例 1：

输入：columnNumber = 1
输出："A"
示例 2：

输入：columnNumber = 28
输出："AB"
示例 3：

输入：columnNumber = 701
输出："ZY"
示例 4：

输入：columnNumber = 2147483647
输出："FXSHRXW"


提示：

1 <= columnNumber <= 231 - 1
*/
package main

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.8 MB, 在所有 Go 提交中击败了44.38%的用户
 */
// 我的题解：哈希
func convertToTitle(columnNumber int) string {

	res := ""
	hash := make(map[int]string)
	for i := 1; i <= 26; i++ {
		hash[i] = string('A' + byte(i-1))
	}
	//
	for columnNumber != 0 {
		temp := columnNumber % 26
		if temp == 0 {
			temp = 26
		}
		columnNumber -= temp
		columnNumber /= 26
		res = hash[temp] + res
	}
	return res
}
