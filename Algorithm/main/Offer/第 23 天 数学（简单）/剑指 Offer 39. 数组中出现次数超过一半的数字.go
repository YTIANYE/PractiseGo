/*
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例1:
输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

限制：
1 <= 数组长度 <= 50000
*/
package main

import "fmt"

// 我的题解：哈希表
/*
执行用时：16 ms, 在所有 Go 提交中击败了76.65%的用户
内存消耗：5.9 MB, 在所有 Go 提交中击败了76.65%的用户
*/
func majorityElement1(nums []int) int {
	length := len(nums)
	mapNums := make(map[int]int)
	for _, num := range nums {
		mapNums[num]++
		if mapNums[num] > length/2 {
			fmt.Println(mapNums)
			return num
		}
	}
	return -1

}

// 我实现的官方题解：摩尔投票 算法
// 不同则抵消，占一半以上的必然留在最后，适合找占一半以上的数，而不是众数，例子[1, 1, 2, 2, 2, 3, 3]
/*
执行用时：16 ms, 在所有 Go 提交中击败了76.65%的用户
内存消耗：5.9 MB, 在所有 Go 提交中击败了83.53%的用户
*/
func majorityElement(nums []int) int {

	res := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			res = num
			count++
		} else if res != num {
			count--
		} else {
			count++
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement(nums))
}
