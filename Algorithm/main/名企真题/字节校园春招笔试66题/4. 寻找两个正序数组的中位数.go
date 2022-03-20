/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

 

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 

 

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106

*/
package main

import "fmt"

// 我的题解：双指针：顺序查找
/*
执行用时：12 ms, 在所有 Go 提交中击败了81.88%的用户
内存消耗：5 MB, 在所有 Go 提交中击败了65.83%的用户
*/

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m == 0 && n == 0 {
		return 0
	}
	mid := (m + n - 1) / 2

	i := m - 1
	j := n - 1
	temp := 0
	for i >= 0 || j >= 0 {
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				temp = nums1[i]
				i--
			} else {
				temp = nums2[j]
				j--
			}
		} else if i >= 0 {
			temp = nums1[i]
			i--
		} else {
			temp = nums2[j]
			j--
		}
		mid--
		if mid == -1 {
			break
		}
	}
	if (m+n)%2 == 1 {
		return float64(temp)
	}
	temp2 := 0
	if i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			temp2 = nums1[i]
		} else {
			temp2 = nums2[j]
		}
	} else if i >= 0 {
		temp2 = nums1[i]
	} else if j >= 0 {
		temp2 = nums2[j]
	}
	return float64(temp+temp2) / 2
}

// 官方题解

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	}
	midIndex1, midIndex2 := totalLength/2-1, totalLength/2
	return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0

}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
