/*
在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
*/
package main

// 我实现的精选题解
func singleNumber(nums []int) int {
	twos, ones := 0, 0
	for _, num := range nums{
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}
	return ones
}
