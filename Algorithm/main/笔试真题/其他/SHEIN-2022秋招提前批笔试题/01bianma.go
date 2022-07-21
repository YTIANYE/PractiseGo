/**
题目：
1 -> A
2 -> B
......
26 -> Z
一串数字可以解析出的字符串的种类数是多少？

测试样例：
12
2
可以得到AB L共两种

226
3

0
0

06
0
 */

package main

import "strconv"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return int整型
 */

func numDecoding( s string ) int {
	// write code here
	hashmap := make(map[string]bool)
	for i:=1;i<=26;i++{
		hashmap[strconv.Itoa(i)] = true
	}


	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i:=1;i<=n;i++{
		if hashmap[string(s[i-1])] {
			dp[i] += dp[i-1]
		}
		if i != 1 && hashmap[s[i-2:i]] {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}