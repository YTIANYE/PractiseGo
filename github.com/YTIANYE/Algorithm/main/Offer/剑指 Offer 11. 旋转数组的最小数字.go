package main

import "fmt"

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：
输入：[3,4,5,1,2]
输出：1

示例 2：
输入：[2,2,2,0,1]
输出：0
*/

// 我的题解
// 执行用时：4 ms, 在所有 Go 提交中击败了75.68%的用户
// 内存消耗：3.1 MB, 在所有 Go 提交中击败了59.47%的用户

func minArray1(numbers []int) int {
	min_num := numbers[0]
	for i := len(numbers)-1 ; i >= 0; i-- {
		min_num = min(min_num, numbers[i])
	}

	return min_num
}

func min(num int, i int) int {
	if num < i{
		return num
	}else{
		return i
	}
}

// 我实现的官方题解
// 执行用时：4 ms, 在所有 Go 提交中击败了75.68%的用户
// 内存消耗：3.1 MB, 在所有 Go 提交中击败了59.47%的用户
func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right{
		mid := left + (right - left) / 2
		if numbers[mid] < numbers[right] {
			right = mid
		}else if numbers[mid] > numbers[right]{
			left = mid + 1
		}else {
			right -= 1
		}
	}
	return numbers[left]
}

func main(){
	numbers := []int{3,4,5,1,2}
	// numbers := []int{2,2,2,0,1}
	fmt.Println(minArray(numbers))
}