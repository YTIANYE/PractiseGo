/**
出行
时间限制： 3000MS
内存限制： 589824KB
题目描述：
因为一些原因，小 A 决定去到 m 公里外的 Z 市。有 n 辆出租车，第 i 辆有起步价 ai  ，即若行驶距离在 bi 公里内，则统一收费 ai 元，若超出 bi 公里，则超出的部分每公里收费 ci 元。

小 A 想知道，如果他选择这 n 辆中的其中一辆出租车乘坐，最少需要花多少钱才能到达 Z 市。



输入描述


输出描述
输出一个整数x，表示最小花费。


样例输入
2 1
2 2
1 2
1 1
样例输出
2

提示
样例解释

有两辆车可供选择，小A无论选择哪一辆，价格都是起步价2元。
 */
package main

import (
	"fmt"
	"math"
)
// 100%
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	as, bs, cs := make([]int, n), make([]int, n), make([]int, n)
	for i:= range as {
		fmt.Scan(&as[i])
	}
	for i:= range bs {
		fmt.Scan(&bs[i])
	}
	for i := range cs {
		fmt.Scan(&cs[i])
	}
	//
	res := math.MaxInt
	for i:=0;i<n;i++ {
		temp := comp(as[i], bs[i], cs[i], m)
		res =  min(res, temp )
	}
	fmt.Println(res)
}

func comp(a, b , c int, m int) int {
	if m <= b {
		return a
	}
	return a + (m - b) * c
}

func min(a, b int ) int {
	if a < b {
		return a
	}
	return b
}