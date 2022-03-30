/*
给你一个整数数组 nums，请你将该数组升序排列。

 

示例 1：

输入：nums = [5,2,3,1]
输出：[1,2,3,5]
示例 2：

输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]
 

提示：

1 <= nums.length <= 5 * 104
-5 * 104 <= nums[i] <= 5 * 104

*/

package main

import "fmt"

func main() {
	nums := []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	// return quickSort(nums)
	// return bubbleSort(nums)
	// return insertSort(nums)
	// return deapSort(nums)
	return mergeSort(nums)
}

// 我的题解：归并排序
/*
执行用时：56 ms, 在所有 Go 提交中击败了52.98%的用户
内存消耗：8.1 MB, 在所有 Go 提交中击败了34.74%的用户
*/
func mergeSort(nums []int) []int {
	var merge func(int, int)
	merge = func(left int, right int) {
		if left >= right {
			return
		}
		mid := (left + right) / 2
		merge(left, mid)
		merge(mid+1, right)
		temp := []int{}
		i := left
		j := mid + 1
		for i <= mid || j <= right {
			if i > mid || (j <= right && nums[j] < nums[i]) {
				temp = append(temp, nums[j])
				j++
			} else {
				temp = append(temp, nums[i])
				i++
			}
		}
		for i := left; i <= right; i++ {
			nums[i] = temp[i-left]
		}
	}
	merge(0, len(nums)-1)
	return nums
}

// 我的题解：堆排序
/*
执行用时：48 ms, 在所有 Go 提交中击败了70.04%的用户
内存消耗：7.4 MB, 在所有 Go 提交中击败了81.00%的用户
*/
func deapSort(nums []int) []int {
	n := len(nums)
	heap := buildHeap(nums)
	for i := n - 1; i >= 0; i-- {
		heap[i], heap[0] = heap[0], heap[i]
		maxHeapify(heap, 0, i)
	}
	return heap
}

// 创建
func buildHeap(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		nums = maxHeapify(nums, i, n)
	}
	return nums
}

// 从上向下调整
func maxHeapify(heap []int, root, heapLen int) []int {
	p := root
	for p*2+1 < heapLen {
		var next int
		l := p*2 + 1
		r := p*2 + 2
		if r >= heapLen || heap[l] > heap[r] {
			next = l
		} else {
			next = r
		}

		if heap[next] > heap[p] { // 大根堆
			heap[next], heap[p] = heap[p], heap[next]
			p = next
		} else {
			break
		}
	}
	return heap
}

// 我的题解：插入排序
func insertSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// 我的题解：冒泡排序
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}

// 我的题解：快速排序
/*
执行用时：716 ms, 在所有 Go 提交中击败了32.62%的用户
内存消耗：12.7 MB, 在所有 Go 提交中击败了27.28%的用户
*/
func quickSort(nums []int) []int {
	var quick func(int, int)
	quick = func(left int, right int) {
		if left >= right {
			return
		}
		first := nums[left]
		i := left
		j := right
		for i < j {
			for i < j && nums[j] > first { // 不是这里 >=
				j--
			}
			for i < j && nums[i] <= first { // 注意数组中存在重复的数，此处<=
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[left] = nums[left], nums[i]
		quick(left, i-1)
		quick(i+1, right)
	}
	quick(0, len(nums)-1)
	return nums
}
