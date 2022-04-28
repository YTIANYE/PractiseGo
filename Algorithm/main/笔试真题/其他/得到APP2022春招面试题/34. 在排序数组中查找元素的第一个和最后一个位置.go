/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回[-1, -1]。

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]

*/

package main

import "fmt"

// 非最佳算法
func find(nums []int, target int) (int, int) {
	start, end := -1, -1
	if len(nums) == 0 {
		return start, end
	}
	//
	var findendstart func(int)
	findendstart = func(i int) {
		for start = i; start >= 0; start-- {
			if nums[start] != nums[i] {
				break
			}
		}
		for end = i; end < len(nums); end++ {
			if nums[end] != nums[i] {
				break
			}
		}
		start++
		end--
	}
	//
	var biaFind func(left, right int) bool
	biaFind = func(left, right int) bool {
		if left > right || left <0 || right >=len(nums){
			return false
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			findendstart(mid)
			return true
		}
		return biaFind(left, mid-1) || biaFind(mid+1, right)
	}

	biaFind(0, len(nums))
	return start, end
}

func main() {
	nums := []int{}
	target := 6
	fmt.Println(find(nums, target))
}

//
// // 分页
//
// select * limit 10
// 	from student ;
