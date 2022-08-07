/**
题目：
数数有多少对 <div> </div>
 */

package main

import "fmt"

/**
<div>hellomixiaoyou</div>
1

<div>hello</div><div>mixiaoyou</div>
2

<div>hello</div>mi<div>xiao<div>you</div></div>
3

hellomixiaoyou
0

 */


// 我的题解：通过样例 100%
func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	left := 0 //"<div>" 个数
	count := 0
	for i:=0;i<n;i++ {
		if i + 5 <= n && s[i: i+5] == "<div>" {
			left++
		}else if i + 6 <= n && s[i:i+6] == "</div>" {
			if left > 0 {
				left--
				count++
			}
		}
	}
	fmt.Println(count)
}
