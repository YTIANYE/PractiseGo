/**
已有n个不同颜色的积木排列在一起，
小团团想要改变该序列，得到自身想要的序列
改变后的积木个数是m个
改变方式两种，左边添加或者右边拿走
问经过最少多少步可以得到想要的积木序列


输入：
n
已知队列
m
想要的队列
输出：
最少需要的步数
*/

package main

import "fmt"

// 我的题解：通过样例18%
func main() {
	var n int
	fmt.Scan(&n)

	jimu := []int{}
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		jimu = append(jimu, num)
	}

	var m int
	fmt.Scan(&m)
	res := []int{}
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		res = append(res, num)
	}

	// 计算
	// 34798
	// 1234

	index := 0
	for i := 1; i <= len(res); i++ { // i 长度
		start := len(res) - i
		tempres := res[start:]
		tempjimu := jimu[:i]

		// fmt.Println("#", tempres, tempjimu)

		tag := true // 匹配
		for j := 0; j < i; j++ {
			if tempres[j] != tempjimu[j] {
				tag = false
			}
		}
		if tag {
			index = max(index, i)
		} else {
			continue
		}
	}
	result := len(jimu) - index + len(res) - index
	fmt.Println(result)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
