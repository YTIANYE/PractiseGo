/**
一个数组中，每个数与其相邻数的差的绝对值的最大值定义为平滑，
改变一个数的大小后，使数组的平滑尽可能最小，求最小的平滑是多少？
 */

package main

import "fmt"

// 52.94%
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i:=0;i<n;i++ {
		fmt.Scan(&nums[i])
	}
	//

	comp := func(ns []int) (int, int ){
		ping := 0
		index := 0
		for i:=0;i+1<n;i++ {
			j := i+1
			cha := abs(ns[i] - ns[j])
			if cha > ping {
				ping = cha
				index = i
			}
		}
		return ping, index
	}

	//
	var res int
	ping , index := comp(nums)


	if index >= 1 {
		nums[index] =  abs(nums[index+1] + nums[index-1])/2
	}

	ping2, _ := comp(nums)

	nums[index] = nums[index+1]
	ping3, _ := comp(nums)

	res = (min(min(ping, ping2), ping3))
	fmt.Println(res)
}

func abs(a int ) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}