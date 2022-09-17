/**
一个字符串只包含‘0’和‘1’
每次交换相邻的两个数，使得字符串中任意相邻的字符不相同，问最少需要经过多少步的交换？
*/

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	indexs := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '1' { // 统计每个 ‘1’ 出现的位置
			indexs = append(indexs, i)
		}
	}

	res1 := 0
	res2 := 0
	// 11100
	// 计算把 1 调整到合适的位置需要的步数
	if len(indexs)*2 > len(s) { // 1个数多
		for i := 0; i < len(indexs); i++ {
			res1 += abs(i*2, indexs[i])
		}
		fmt.Println(res1)
	} else if len(indexs)*2 < len(s) { // 1个数少
		for i := 0; i < len(indexs); i++ {
			res2 += abs(i*2+1, indexs[i])
		}
		fmt.Println(res2)
	} else { // 1 0 个数相等
		for i := 0; i < len(indexs); i++ {
			res1 += abs(i*2, indexs[i])
			res2 += abs(i*2+1, indexs[i])
		}
		fmt.Println(min(res1, res2))
	}
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
