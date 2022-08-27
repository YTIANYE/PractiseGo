/**
DNA匹配
时间限制： 3000MS
内存限制： 589824KB
题目描述：
有一种特殊的DNA，仅仅由核酸A和T组成，长度为n，顺次连接

科学家有一种新的手段，可以改变这种DNA。每一次，科学家可以交换该DNA上两个核酸的位置，也可以将某个特定位置的核酸修改为另一种核酸。

现在有一个DNA，科学家希望将其改造成另一种DNA，希望你计算最少的操作次数。



输入描述
输入包含两行，第一行为初始的DNA，第二行为目标DNA，保证长度相同。

输出描述
输出最少的操作次数


样例输入
ATTTAA
TTAATT
样例输出
3

提示
对于100%的数据，DNA长度小于等于100000
样例解释：
1.首先修改第一个位置的核酸（从A修改为T）
2.交换3和5位置的核酸
3.交换4和6位置的核酸
*/
package main

import "fmt"

// 通过 100%
func main() {
	var a, b string
	res := 0
	fmt.Scan(&a, &b)
	//
	i := 0
	for i < len(a) {
		if a[i] == b[i] {
			a = remove(a, i)
			b = remove(b, i)
		} else {
			i++
		}

	}

	//
	n := len(a)
	anumT, bnumT := numT(a), numT(b)
	budeng := abs(anumT - bnumT)
	res += budeng
	res += (n - budeng) / 2
	fmt.Println(res)

}

func numT(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'T' {
			res++
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func remove(a string, i int) string {
	if len(a) == 1 && i == 0 {
		return ""
	} else if i == 0 {
		return a[1:]
	} else if i == len(a)-1 {
		return a[:i]
	} else {
		return a[:i] + a[i+1:]
	}
}
