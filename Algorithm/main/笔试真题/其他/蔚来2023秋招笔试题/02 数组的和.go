/**
题目：
一个长度为n的数组，对其进行q次操作，每次操作是数组中的每个数加上一个数，
输出每次操作后数组每个元素的绝对值的和

输入n
输入数组元素
输入q
输入每次的操作数
 */

package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	nums := make([]int64, n)
	var sum int64
	for i := range nums {
		fmt.Scan(&nums[i])
		sum += nums[i]

	}
	var q int64
	fmt.Scan(&q)
	qx := make([]int64, q)
	for i := range qx {
		fmt.Scan(&qx[i])
	}
	////// 80%
	var i int64
	for i = 0; i < q; i++ { // 每次操作
		sum = 0
		num := qx[i]
		var j int64
		for j = 0; j < n; j++ {
			nums[j] = num + nums[j]
			sum += abs(nums[j])
		}
		fmt.Println(sum)
	}

}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func panduan(a, b int64) bool {
	if a >= 0 && b >= 0 {
		return true
	}
	if a < 0 && b < 0 {
		return true
	}
	return false
}
//
//func findIndex(nums []int64) int64 {
//	for i := 0; i < len(nums); i++ {
//		if nums[i] >= 0 {
//			return i
//		}
//	}
//	return len(nums) // 所有数都小于0
//}
//
