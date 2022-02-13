/*
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

示例 1：

输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
示例 2：

输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]


限制：
2 <= nums.length <= 10000

*/
package main

import "fmt"

// 我实现的官方题解
/*
执行用时：16 ms, 在所有 Go 提交中击败了84.40%的用户
内存消耗：6.1 MB, 在所有 Go 提交中击败了99.77%的用户
*/
func singleNumbers(nums []int) []int {

	// 找到两个数的异或结果
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		num = num ^ nums[i]
	}
	// 异或结果为1的 位
	temp := 1
	for temp&num == 0 {
		temp = temp << 1
	}

	// 分组
	num1 := 0
	num2 := 0
	for _, n := range nums {
		if n&temp == 0 {
			num1 ^= n
		} else {
			num2 ^= n
		}
	}
	return []int{num1, num2}
}

func main() {
	nums := []int{1, 2, 10, 4, 1, 4, 3, 3}
	fmt.Println(singleNumbers(nums))
}
