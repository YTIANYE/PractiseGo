package main

import (
	"fmt"
	"math"
	"strconv"
)

func minNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	// 冒泡排序
	/*
		执行用时：8 ms, 在所有 Go 提交中击败了14.38%的用户
		内存消耗：3.1 MB, 在所有 Go 提交中击败了27.09%的用户
	*/
	/*	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if comp(nums[j], nums[j+1]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}*/

	// 快排
	/* 执行用时：4 ms, 在所有 Go 提交中击败了65.72% 的用户
	内存消耗：2.9 MB, 在所有 Go 提交中击败了37.46% 的用户*/
	var quickSort func(int, int)
	quickSort = func(l, r int) {
		if l >= r {
			return
		}

		i, j := l, r
		for i < j {
			for i < j && comp(nums[j], nums[l]) {
				j--
			}
			for i < j && !comp(nums[i], nums[l]) {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[l] = nums[l], nums[i]
		quickSort(l, i-1)
		quickSort(i+1, r)
	}

	quickSort(0, len(nums)-1)

	// 返回结果
	str := ""
	for _, num := range nums {
		str += strconv.Itoa(num)
	}
	return str
}

func comp(a, b int) bool {
	num1 := float64(a)*math.Pow(10, float64(len(strconv.Itoa(b)))) + float64(b)
	num2 := float64(b)*math.Pow(10, float64(len(strconv.Itoa(a)))) + float64(a)
	if num1 > num2 {
		return true // a > b  需要交换
	}
	return false
}

func main() {
	nums := []int{3,30,34,5,9}//3033459
	// nums := []int{10,2}
	// nums := []int{1,1,1}
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// nums := []int{4704,6306,9385,7536,3462,4798,5422,5529,8070,6241,9094,7846,663,6221,216,6758,8353,3650,3836,8183,3516,5909,6744,1548,5712,2281,3664,7100,6698,7321,4980,8937,3163,5784,3298,9890,1090,7605,1380,1147,1495,3699,9448,5208,9456,3846,3567,6856,2000,3575,7205,2697,5972,7471,1763,1143,1417,6038,2313,6554,9026,8107,9827,7982,9685,3905,8939,1048,282,7423,6327,2970,4453,5460,3399,9533,914,3932,192,3084,6806,273,4283,2060,5682,2,2362,4812,7032,810,2465,6511,213,2362,3021,2745,3636,6265,1518,8398}
	// nums := []int{121,12}
	fmt.Println(minNumber(nums))

}
