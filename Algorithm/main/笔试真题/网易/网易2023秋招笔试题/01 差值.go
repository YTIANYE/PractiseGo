/**
计算数组中两个数的最小差值
 */

package main

import (
	"fmt"
	"sort"
)

// 100%
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:= range nums {
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)
	res := abs(nums[0] - nums[1])
	for i:= 0 ;i<n-1;i++ {
		res = min(res, abs(nums[i]- nums[i+1]))
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b{
		return a
	}
	return b
}

func abs(a int )  int {
	if a > 0 {
		return  a
	}
	return -a
}