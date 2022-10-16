/**
回文串
时间限制： 3000MS
内存限制： 589824KB
题目描述：
如果一个字符串正着读和反着读能够得到相同的结果，那么我们称其为回文串。小A前几年过生日的时候收到了一条回文串作为生日礼物。但是随着时间的推移这个串上某些位置的字符发生了跳变。现在给出修改每个位置上的字符所需的代价，请你计算将其变成回文串至少需要多少代价。



输入描述
第一行有一个正整数n（1<=n<=100000)，代表回文串的长度。

第二行有一个长度为n的仅由小写英文字母组成的字符串，代表变质的回文串。

第三行有n个正整数，分别代表修改每个字符所需的代价。大小在1到1000000之间。

输出描述
输出一个非负整数，代表将变质的回文串修复成回文串所需的最小代价。


样例输入
5
abcab
1 5 2 4 3
样例输出
5

提示
可将字符串修改为bbcbb，此时代价为1+4=5
 */


package main

import "fmt"

func main () {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	nums := make([]int, n)
	for i:= range nums {
		fmt.Scan(&nums[i])
	}
	//
	res := 0
	for i:=0;i<n/2;i++ {
		j := n-i-1
		if s[i] != s[j] {
			res += min2(nums[i], nums[j])
		}
	}
	fmt.Println(res )
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
