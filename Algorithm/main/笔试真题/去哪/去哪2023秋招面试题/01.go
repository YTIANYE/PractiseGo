/**
编程题：我现在有商家的所有商品报价，现在要和别的商户竞争，我想排在所有竞争对手报价的第N名，请写一个程序，我输入N,和一个商品名，系统告诉我应该定多少价？
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	hash := make(map[string][]int)
	// 填充数据，假设价格均为整数

	// 输入
	var n int
	var s string
	fmt.Scan(&n, &s)

	// 计算报价
	var baojia func(n int, s string) int
	baojia = func(n int, s string) int {
		nums := hash[s]
		sort.Ints(nums)
		//
		if n > len(nums) {
			return -1
		}
		n--
		left := nums[n]
		return left - 1
	}
	// 结果
	fmt.Println(baojia(n, s))

}
