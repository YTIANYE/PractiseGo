/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

 

示例 1:

输入: [7,5,6,4]
输出: 5

*/

package main

/*
7	5	6	4
			0
		1	0
	1	1	0
3
################
7	5	6	4
0
0	1
0	1	1
0	1	1	3



8	9	7	5	6	4
					0
				1
			1
		3
	4
4

4		0
64		1
564		1+1
7564	1*2+1*2+1

*/

// 暴力方法：超时
func reversePairs1(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}

	sum := 0

	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				sum++
			}
		}
	}
	return sum
}

// 归并排序：
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}

	// 递归
	mid := (left + right) / 2
	res := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)

	i, j := left, mid+1
	temp := []int{}
	// 计数
	for i <= mid && j <= right {
		if nums[i] <= nums[j] { // 注意符号，<=
			temp = append(temp, nums[i])
			res += j - (mid + 1) // 注意计数
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}
	// 处理剩余
	for i <= mid {
		temp = append(temp, nums[i])
		res += right - mid
		i++
	}
	for j <= right {
		temp = append(temp, nums[j])
		j++
	}
	// 覆盖
	for i := 0; i < len(temp); i++ {
		nums[left+i] = temp[i]
	}
	return res
}
