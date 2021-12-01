package main

import (
	"fmt"
	"sort"
)

/*
在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字。
*/

//我的题解
// 执行用时：44 ms, 在所有 Go 提交中击败了15.88%的用户
// 内存消耗：8 MB, 在所有 Go 提交中击败了88.68%的用户

func findRepeatNumber1(nums []int) int{
	sort.Ints(nums)
	for i,num := range nums{
		if num == nums[i+1]{
			return num
		}
	}
	return -1
}

/*我实现的精选题解*/

// 执行用时：32 ms, 在所有 Go 提交中击败了89.62%的用户
// 内存消耗：8 MB, 在所有 Go 提交中击败了74.29%的用户

func findRepeatNumber(nums []int) int{
	for i:=0; i<len(nums); {
		num := nums[i]
		if num == i{
			i++
			continue
		}else if num == nums[num]{
			return num
		}else{
			nums[i], nums[num] = nums[num], nums[i]
		}
	}
	return -1
}

func main(){
	nums := []int{3, 4, 2, 0, 0, 1}
	fmt.Println(findRepeatNumber(nums))
}
