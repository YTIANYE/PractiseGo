package main

import (
	"fmt"
	"sort"
)

/*
从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

示例1:

输入: [1,2,3,4,5]
输出: True


示例2:

输入: [0,0,1,2,5]
输出: True


限制：
数组长度为 5 
数组的数取值为 [0, 13] .

*/

// 我的题解：暴力题解
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了83.27%的用户
*/
func isStraight1(nums []int) bool {
	sort.Ints(nums) // 默认从小到大

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		}
	}

	for i := count; i < len(nums)-1; i++ {
		temp := nums[i+1] - nums[i]
		if temp == 0 {
			return false
		}
		if temp != 1 {
			if temp-1-count > 0 {
				return false
			} else {
				count -= temp - 1
			}
		}

	}
	return true
}

// 精选题解
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了38.43%的用户
*/
func isStraight(nums []int) bool {
	sort.Ints(nums)

	joker := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			joker++
		}else if nums[i] == nums[i+1] {
			return false
		}
	}

	//  最大牌 - 最小牌 < 5 则可构成顺子
	return nums[4]-nums[joker] < 5
}

func main() {
	nums := []int{2, 2, 0, 0, 6}
	// nums := []int{11, 0, 9, 0, 0}
	fmt.Println(isStraight(nums))

}
