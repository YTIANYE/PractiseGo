/**
一共有n个士兵，序号为1...n
每个士兵可以防守或进攻，力量为其序号
一个长度为n的字符串s由0和1构成，0表示进攻，1表示防守
从s的某个位置pos切开,pos = 0...n
1...pos 左边的士兵进攻
pos+1...n 右边的士兵防守

求左边的进攻总值与右边的防御总值的绝对值最小是多少？
*/

package main

import "fmt"

// 通过样例 100%
func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	var l0, l1, r0, r1 int
	for i := range s {
		v := i + 1 // 力量值
		if s[i] == '0' {
			r0 += v
		} else {
			r1 += v
		}
	}

	// 计算
	res := abs(l0, r1)
	for pos := 1; pos <= n; pos++ {
		i := pos - 1
		if s[i] == '0' {
			l0 += pos
			r0 -= pos
		} else {
			l1 += pos
			r1 -= pos
		}
		res = min1(res, abs(l0, r1))
	}

	fmt.Println(res)

}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}
