/**
设有编号为1，2，……，n的n(n>0)个人围成一个圈，从第1个人开始报数，报到m时停止报数，
报m的人出圈，再从他的下一个人起重新报数，报到m时停止报数，
报m的出圈，……，如此下去，直到所有人全部出圈为止。当任意给定n和m后，设计算法求n个人出圈的次序

输入：5 3
输出： 31524

输入: 1 2
输出：1

*/

package main

import "fmt"

func main() {
	res := chuquan(1, 2)
	fmt.Println(res)
}

// 我的题解： 模拟
func chuquan2(n, m int) []int {
	res := []int{}
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	//
	index := 0
	start := 0
	for index <= m-1 {
		if index == m-1 {
			j := (start + index) % len(nums)
			res = append(res, nums[j])
			nums = remove(nums, j)
			if len(nums) == 0 {
				break
			}
			start = j % len(nums) // 注意 容易出错
			index = -1
		}
		index++
	}
	return res
}

func remove(nums []int, i int) []int {
	if i == 0 {
		return nums[1:]
	} else if i == len(nums)-1 {
		return nums[:len(nums)-1]
	} else {
		nums = append(nums[:i], nums[i+1:]...)
		return nums
	}
}

// 其他题解：模拟、标记
func chuquan(n, m int) []int {
	res := []int{}
	tag := make([]int, n)

	count := 0
	index := -1
	for count < n {
		i := 0
		for i < m {
			index = (index + 1) % n
			if tag[index] == 0 {
				i++
			}
		}
		tag[index] = 1
		res = append(res, index+1)
		count++
	}
	return res
}
