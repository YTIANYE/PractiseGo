/**
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



 

示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]
 

提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

*/

package main

// 我的题解：暴力题解
/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了32.23%的用户
 */

func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	hash := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	res := []string{""}
	for i := range digits {
		s := string(digits[i])
		strs := hash[s]
		temp := make([]string, len(res))
		copy(temp, res)
		res = []string{}
		for j := range strs {
			for t := range temp {
				res = append(res, temp[t]+strs[j])
			}
		}
	}
	return res

}
