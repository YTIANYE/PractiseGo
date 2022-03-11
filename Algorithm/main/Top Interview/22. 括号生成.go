/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
 

提示：

1 <= n <= 8


*/
package main

import "fmt"

/*
()

()() (())

()()()
(())() ()(()) (()())
((()))

()()()()
(())()() ()(())() ()()(()) (())(()) ()(()()) (()())()
()((())) ((()))() (()()()) ((())()) (()(())) ((()())) (((())))

*/

// 我的题解：不断的在上一层的结果中插入"()"，然后去重

func generateParenthesis(n int) []string {
	res := make(map[string]int)
	res["()"] = 1

	for i:=1;i<n;i++{
		tempres := res
		res =  make(map[string]int)
		for str := range tempres{
			l := len(str)
			for j:=0;j<=l;j++{
				temp := str[:j] + "()" + str[j:]
				res[temp]++
			}
		}
	}
	result := []string{}
	for r := range res{
		result = append(result, r)
	}
	return result
}

func main(){
	n := 2
	fmt.Println(generateParenthesis(n))
}