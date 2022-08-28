/**
长城：一个数的相邻数相等且与这个数本身不等
如：
1 2 1 2 1 2 1
n个数排成一排，将这排数改为长城，至少需要多少步？
每一步可以将一个数做一次更改

例子：
输入：
6
1 1 1 4 5 6
输出
3
*/

package main

import "fmt"

// 100%
func main() {
	var n int
	fmt.Scan(&n)
	var num int
	hashji := make(map[int]int)
	hashou := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if i%2 == 0 {
			hashou[num]++
		} else {
			hashji[num]++
		}
	}
	jimaxnum, jimax := many(hashji)
	oumaxnum, oumax := many(hashou)
	//
	var res int
	if jimaxnum != oumaxnum {
		res = n - jimax - oumax
	} else {
		delete(hashji, jimaxnum)
		delete(hashou, oumaxnum)
		_, jimax2 := many(hashji)
		_, oumax2 := many(hashou)
		max2 := max(jimax+oumax2, oumax+jimax2)
		res = n - max2
	}

	fmt.Println(res)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func many(hash map[int]int) (int, int) {
	res := 0
	num := 0
	for k, v := range hash {
		if res < v {
			num = k
			res = v
		}
	}
	return num, res
}
