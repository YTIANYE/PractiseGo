/**
题目描述：

K语言的每个代码块由仅包含一个关键字xxx的一行开头，并以仅包含end xxx的一行结束。其中xxx是一个长度不超过100的仅由小写英文字母组成的字符串。与其他语言类似，K语言中的代码块仅允许互相包含，不允许仅有部分重叠。现在给出几段仅包含代码块开头行和结尾行的K语言程序，请你判断它们是否符合K语言的语法。



输入描述
第一行有一个正整数T(1<=T<=100)，代表有T段K语言程序需要你进行判断。

每段K语言程序前面均有一个正整数n(1<=n<=100)代表这段K语言程序包含多少行。所有n之和不超过500。

每段K语言程序均由小写英文字母和空格组成，且关键字不包含空格。每行的长度不超过100。

输出描述
对于每段K语言程序，若其符合语法则输出Yes，否则输出No。


样例输入
2
10
if
while
loop
end loop
end while
for
end for
end if
switch
end switch
4
if
while
end if
end while
样例输出
Yes
No
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
// 通过样例 27%


func main() {
	var t int
	fmt.Scan(&t)
	input := bufio.NewScanner(os.Stdin)
	for ; t != 0; t-- {
		input.Scan()
		str := strings.Split(input.Text(), " ")
		n, _ := strconv.Atoi(str[0])
		// fmt.Println(n)
		stack := []string{}
		res := "Yes"
		for l := 0; l < n; l++ {
			input.Scan()
			strs := strings.Split(input.Text(), " ") // 一个空格还是几个空格
			var first string
			for i := range strs { // 前导空格
				if strs[i] != "" && strs[i] != " " {
					first = strs[i]
					break
				}
			}
			// fmt.Println("first:", first)
			if first == "end" {
				var s string
				for i:=1;i<len(strs);i++ {
					if strs[i] != "" && strs[i] != " " {
						s = strs[i]
					}
				}
				// fmt.Println("s", s)
				if !isright(s) || len(stack) == 0 || stack[len(stack)-1] != s {
					res = "No"
				} else if stack[len(stack)-1] == s{
					stack = stack[:len(stack)-1]
				}
			} else if isright(first){
				stack = append(stack, strs[0])
			} else{
				res = "No"
			}
			// fmt.Println(stack)
		}
		if len(stack) != 0 {
			res = "No"
		}
		fmt.Println(res)
	}
}

func isright(s string) bool{
	if len(s) >100{
		return false
	}
	for i := range s{
		if !('a' <= s[i] && s[i] <= 'z'){
			return false
		}
	}
	return true
}
