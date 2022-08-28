/**
裁缝
时间限制： 3000MS
内存限制： 589824KB
题目描述：
小团最近获得了美美团团国的裁缝资格证，成为了一个裁缝！现在小团有一个长度为n的大布料S（在这个国家布料其实是一个仅包含小写英文字母的字符串），小团可以将布料在任意位置剪切，例如布料abcd可以被裁剪为a与bcd 或ab与cd或abc与d ，不过，裁剪完之后是不能拼接起来的，因为小团还不是很擅长拼接布料。现在小团想知道能不能有一种裁剪方式能让他把布料恰好裁剪成客人要求的小布料。

形式化地，有一个串S，问是否能将其划分成m个不相交的连续子串，使得这些连续子串可以与要求的连续子串一 一对应。两个串相对应是指这两个串完全相等。例如"aab"="aab" 但 "aab"≠"baa"



输入描述
第一行两个空格隔开的正整数n和m，分别表示大布料S长度和客人要求的小布料数量。

接下来一行一个长度为n的仅包含小写英文字母的串S，表示大布料的组成。

接下来一行m个空格隔开的数x1,x2, ...,xm ，依次表示所要求的小布料长度。

接下来开始m行，每行一个长度为xi的仅包含小写英文字母的串si，表示第i个小布料的组成。

对于所有数据，



输出描述
如果存在这样的方案，输出方案总数。如果不存在任何方案，输出0。两个方案A、B不相同当且仅当方案A中存在一个相对于原始长布料的裁剪位置i，而方案B中并未在该位置i裁剪。

例如aaaaaa 裁剪方案aaa|aaa 与方案 aaa|aaa是相同的方案。而方案aa|aaaa与方案aaaa|aa是不同的方案，虽然划分出的结果都是aa与aaaa，但前者在第二个a处进行了裁剪，后者并没有在这里进行裁剪，所以视为不同的裁剪方案。




样例输入
6 2
aaaaaa
4 2
aaaa
aa
样例输出
2

提示
有两种方案，第一种是aaaa|aa，第二种是aa|aaaa，代表一次裁剪。
 */

/**
6 2
aaaaaa
4 2
aaaa
aa

2
*/

package main

import "fmt"


// 27%
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var S string
	fmt.Scan(&S)
	lnums := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&lnums[i])
	}
	strs := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&strs[i])
	}
	//
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			l := lnums[j]
			s := strs[j]
			left := i + 1 - l
			if left >= 0 && isSame1(S[left:i+1], s) { // i = 3  长度 3
				dp[i+1] += dp[left]
			}
		}
	}
	fmt.Println(dp[n])
}

func isSame1(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
