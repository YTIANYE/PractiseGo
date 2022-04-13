package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)
	for ; T > 0; T-- {
		// 输入
		var N int
		fmt.Scan(&N)
		var canzhuo string
		fmt.Scan(&canzhuo)
		var M int
		fmt.Scan(&M)
		var ren string
		fmt.Scan(&ren)

		// 座位统计
		res := []int{}
		zhuo0 := []int{}
		zhuo1 := []int{}
		zhuo2 := []int{}
		for i := range canzhuo {
			temp := int(canzhuo[i] - '0')
			if temp == 0 {
				zhuo0 = append(zhuo0, i)
			} else if temp == 1 {
				zhuo1 = append(zhuo1, i)
			} else {
				zhuo2 = append(zhuo2, i)
			}
		}
		// 选择座位
		sitzhuo1 := func() {
			index := zhuo1[0]
			zhuo1 = zhuo1[1:]
			zhuo2 = append(zhuo2, index)
			res = append(res, index)
		}
		sitzhuo0 := func() {
			index := zhuo0[0]
			zhuo0 = zhuo0[1:]
			zhuo1 = append([]int{index}, zhuo1...)
			sort.Ints(zhuo1)
			res = append(res, index)
		}

		for i := range ren {
			if ren[i] == 'M' {
				if len(zhuo1) != 0 {
					sitzhuo1()
				} else {
					sitzhuo0()
				}
			} else {
				if len(zhuo0) != 0 {
					sitzhuo0()
				} else {
					sitzhuo1()
				}
			}
		}

		// 结果
		for i := range res {
			fmt.Println(res[i]+1)
		}
	}

}
