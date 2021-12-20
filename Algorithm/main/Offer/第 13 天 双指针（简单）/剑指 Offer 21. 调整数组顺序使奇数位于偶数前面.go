/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。

示例：
输入：nums =[1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

提示：
0 <= nums.length <= 50000
0 <= nums[i] <= 10000

*/
package main

import "fmt"

/*我的题解*/
// 执行用时：20 ms, 在所有 Go 提交中击败了80.94%的用户
// 内存消耗：6.8 MB, 在所有 Go 提交中击败了93.76%的用户
func exchange(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	p, q := 0, len(nums)-1
	for p < q {
		// 我的写法
		// for p < len(nums) && isJishu(nums[p]) {
		// 	p++
		// }
		// for q>=0 && !isJishu(nums[q]) {
		// 	q--
		// }
		// if p < q {
		// 	nums[p], nums[q] = nums[q], nums[p]
		// }

		// 精选题解的写法
		// 执行用时：20 ms, 在所有 Go 提交中击败了80.94% 的用户
		// 内存消耗：6.8 MB, 在所有 Go 提交中击败了93.76% 的用户
		for p < q && isJishu(nums[p]) {
			p++
		}
		for p < q && !isJishu(nums[q]) {
			q--
		}
		nums[p], nums[q] = nums[q], nums[p]
	}
	return nums
}

func isJishu(num int) bool {
	if num%2 == 1 {
		return true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4}
	// nums := []int{1, 3, 5}
	// nums := []int{2, 4, 6}
	fmt.Println(exchange(nums))
}
