/**
找出一个数组中缺失的最小的非负数，数组中大于这个数的全部改为这个数
 */
package main

import (
	"fmt"
	"sort"
)

//
//// 63%
//func main() {
//	var n int
//	fmt.Scan(&n)
//	nums := make([]int, n)
//	hash := make(map[int]int)  // key:val index
//	for i:=0;i<n;i++{
//		fmt.Scan(&nums[i])
//		hash[nums[i]] = i
//	}
//	//
//	res := make([]int, n)
//	sort.Ints(nums)
//	for i:=0;i<n;i++ {
//
//		// 删除nums[i]
//		temp := makeTemp(nums, i )
//		//
//		for num:= 0;num <n-1;num++ {
//			if num != temp[num] {
//				res[hash[nums[i]]] = num
//				break
//			}
//		}
//	}
//	//
//	fmt.Printf("%d", res[0])
//	for i:=1;i<len(res);i++ {
//		fmt.Printf(" %d", res[i])
//	}
//}
//
//func makeTemp(nums []int, i int) []int {
//	n := len(nums)
//	temp := []int{}
//	for j:=0;j<n;j++ {
//		if j == i {
//			continue
//		}
//		temp = append(temp, nums[j])
//	}
//	return temp
//}

// 100%
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	//
	yuan := make([]int, n)
	copy(yuan, nums)
	sort.Ints(nums)
	var que int
	tag := false
	for i := 0; i < n; i++ {
		if nums[i] != i {
			que = i
			tag = true
			break
		}
	}
	if tag == false {
		que = n
	}
	//
	res := []int{}
	for i := range yuan {
		if yuan[i] > que {
			res = append(res, que)
		} else {
			res = append(res, yuan[i])
		}
	}
	//
	fmt.Printf("%d", res[0])
	for i := 1; i < len(res); i++ {
		fmt.Printf(" %d", res[i])
	}
}

// 82%
