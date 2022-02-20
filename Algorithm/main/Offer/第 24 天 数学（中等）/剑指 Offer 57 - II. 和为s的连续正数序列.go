/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：
输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：
输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

*/

package main

import (
	"fmt"
)

// 我的题解：分奇偶计算
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了77.24%的用户
*/

func findContinuousSequence1(target int) [][]int {
	res := [][]int{}
	for i := 2; i < target; i++ { // i 表示数组中有几个数
		var num int
		if i%2 == 1 { // 奇数个元素
			if target%i == 0 {
				num = target/i - (i-1)/2
			} else {
				continue
			}
		} else { // 偶数个元素
			if float64(target)/float64(i)-float64(target/i) == 0.5 {
				num = target/i + 1 - i/2
			} else {
				continue
			}
		}

		if num <= 0 {
			continue
		}
		nums := []int{}
		for j := 0; j < i; j++ {
			nums = append(nums, num+j)
		}
		res = append([][]int{nums}, res...)
	}

	return res
}

// 我实现的精选题解:滑动窗口
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了93.29%的用户
*/
func findContinuousSequence(target int) [][]int {
	left, right, sum := 1, 2, 3
	res := [][]int{}

	for left < right {
		if sum == target{
			nums := []int{}
			for i:=left;i<=right;i++{
				nums = append(nums, i)
			}
			res = append(res, nums)
		}
		if sum >= target{
			sum -= left
			left++
		}else{
			right++
			sum += right
		}
	}
	return res
}

func main() {
	target := 15
	fmt.Println(findContinuousSequence(target))
}
