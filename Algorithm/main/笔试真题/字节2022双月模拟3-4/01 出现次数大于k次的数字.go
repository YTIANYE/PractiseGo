/**
统计出现次数大于k次的数字，并由小到大排列
*/

package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	hash := make(map[int]int)
	for {
		var num int
		if _, ok := fmt.Scan(&num); ok != io.EOF {
			hash[num]++
		} else {
			break
		}
	}

	nums := []int{}
	for num, v := range hash {
		if v >= k {
			nums = append(nums, num)
		}
	}
	sort.Ints(nums)

	// 打印
	for i, num := range nums {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}

}
