/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

 

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]
 

限制：

0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000

*/
package main

import (
	"fmt"
)

// 我的题解：暴力方法，维持一个长度为k的数组
/*
执行用时：88 ms, 在所有 Go 提交中击败了13.40%的用户
内存消耗：6.6 MB, 在所有 Go 提交中击败了21.49%的用户
*/
func getLeastNumbers1(arr []int, k int) []int {
	if len(arr) < k || k <= 0 || len(arr) == 0 {
		return []int{}
	}

	res := []int{}
	res = append(res, arr[:k]...)
	maxindex := max(res)
	for i := k; i < len(arr); i++ {
		if arr[i] < res[maxindex] {
			res[maxindex] = arr[i]
			maxindex = max(res)
		}
	}

	return res
}

// 返回一个数组中最大值的索引
func max(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	index := 0
	for i := range arr {
		if arr[i] > arr[index] {
			index = i
		}
	}
	return index
}

// 我的题解：快速排序
/*
执行用时：20 ms, 在所有 Go 提交中击败了97.95%的用户
内存消耗：6.5 MB, 在所有 Go 提交中击败了93.67%的用户
*/

func getLeastNumbers(arr []int, k int) []int {
	if k <= 0 || len(arr) == 0 || k > len(arr) {
		return []int{}
	}
	quickSort(&arr, 0, len(arr)-1,k)
	return arr[:k]
}

func quickSort(nums *[]int, left, right, k int) {
	if left >= right {
		return
	}

	arr := *nums
	i, j := left, right
	for i < j {
		for i < j && arr[j] >= arr[left] {
			j--
		}
		for i < j && arr[i] <= arr[left] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[left] = arr[left], arr[i]

	if i < k {
		quickSort(nums, i+1, right, k)
	} else if i > k {
		quickSort(nums, left, i-1, k)
	} else {
		return
	}
}

func main() {

	// arr := []int{1, 2, 6, 0, 5}
	// k := 2

	arr := []int{0,1,2,1}
	k := 1

	fmt.Println(getLeastNumbers(arr, k))
}
