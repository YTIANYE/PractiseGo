/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
 

示例 1：

输入：nums = [1,2,0]
输出：3
示例 2：

输入：nums = [3,4,-1,1]
输出：2
示例 3：

输入：nums = [7,8,9,11,12]
输出：1
 

提示：

1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1

*/

package main

import (
	"fmt"
	"sort"
)

// 我的题解：
/*
执行用时：132 ms, 在所有 Go 提交中击败了31.11%的用户
内存消耗：24.5 MB, 在所有 Go 提交中击败了58.85%的用户
*/

func firstMissingPositive(nums []int) int {
	hashMap := make(map[int]int) // min max

	for _, num := range nums{
		if len(hashMap) == 0{
			hashMap[num] = num
			continue
		}
		flags := true
		for k, v := range hashMap{
			if k-1 == num{
				hashMap[num] = v
				delete(hashMap, k)
				flags = false
				break
			}else if v+1 == num {
				hashMap[k]++
				flags = false
				break
			}else if k == num || v == num{
				flags = false
				break
			}
		}
		if flags{
			hashMap[num] = num
		}

	}

	keys := []int{}
	for k := range hashMap{
		keys = append(keys, k)
	}
	sort.Ints(keys)

	res := 1
	for _, minNum := range keys{
		maxNum := hashMap[minNum]
		if res < minNum{
			return res
		}else if res <= maxNum{// 注意符号
			res = maxNum+1
		}
	}
	return res

}

// 官方题解：索引打标记
func firstMissingPositive1(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			fmt.Println(num-1)
			nums[num - 1] = -abs(nums[num - 1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}



func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 官方题解：交换数据

func firstMissingPositive2(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return n + 1
}

func main(){
	// nums := []int{1,2,0}
	nums := []int{1,2,2,1,3,1,0,4,0}
	fmt.Println(firstMissingPositive(nums))
}