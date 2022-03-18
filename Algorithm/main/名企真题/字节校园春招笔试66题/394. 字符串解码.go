/*

给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

 

示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
 

提示：

1 <= s.length <= 30
s 由小写英文字母、数字和方括号 '[]' 组成
s 保证是一个 有效 的输入。
s 中所有整数的取值范围为 [1, 300] 


*/

package main

import (
	"fmt"
	"strconv"
)


// 我的题解：辅助栈
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：6 MB, 在所有 Go 提交中击败了8.11%的用户
*/

func decodeString1(s string) string {

	stack := []uint8{}
	for i := range s{
		if s[i] == ']'{
			// 获取【】中的内容
			tempstring := ""
			for stack[len(stack)-1] != '['{
				tempstring = string(stack[len(stack)-1]) + tempstring
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			// 获取【】前面的数字
			num := ""
			for len(stack) >0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9'{
				num = string(stack[len(stack)-1]) + num
				stack = stack[:len(stack)-1]
			}
			numInt, _ := strconv.Atoi(num)
			// 获取 数字*【】 的结果
			tempres := ""
			for ;numInt != 0;numInt--{
				tempres += tempstring
			}
			// 新的字符串装入栈中
			for i := range tempres{
				stack = append(stack, tempres[i])
			}
		}else{
			stack = append(stack, s[i])
		}
	}

	res := ""
	for i := range stack{
		res += string(stack[i])
	}
	return res
}

// 我实现的官方题解：辅助栈
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了31.47%的用户
*/

func decodeString(s string ) string{
	stackMulti := []int{}
	stackRes := []string{}
	res := ""
	multi := 0

	for i := range s{
		if s[i] == '['{
			stackMulti = append(stackMulti, multi)
			stackRes = append(stackRes, res)
			multi = 0
			res = ""
		}else if s[i] == ']'{
			curMulti := stackMulti[len(stackMulti)-1]
			lastRes := stackRes[len(stackRes)-1]
			stackMulti = stackMulti[:len(stackMulti)-1]
			stackRes = stackRes[:len(stackRes)-1]
			tempres := res
			res = lastRes
			for ;curMulti !=0;curMulti--{
				res += tempres
			}
		}else if s[i] >= '0' && s[i] <= '9'{
			multi = multi * 10 + int(s[i] - '0')
		}else{
			res += string(s[i])
		}
	}
	return res
}

func main(){
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))
}