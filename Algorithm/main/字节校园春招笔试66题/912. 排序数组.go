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

func sortArray(nums []int) []int {
	// return quickSort(nums)
	// return bubbleSort(nums)
	return insertSort(nums)
}


// 我的题解：插入排序
func insertSort(nums []int) []int{
	for i:=0;i<len(nums)-1;i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i] > nums[j]{
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// 我的题解：冒泡排序
func bubbleSort(nums []int) []int{
	for i:=0;i<len(nums)-1;i++{
		for j:=1;j<len(nums)-i;j++{
			if nums[j-1] > nums[j]{
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

