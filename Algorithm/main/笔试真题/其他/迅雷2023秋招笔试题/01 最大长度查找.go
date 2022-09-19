/**
一个数组中寻找子数组
子数组满足条件是：
回文数组
子数组中数字的位置在原数组中相对不变，且连续

求满足这样条件的最长的子数组的长度是多少？

例子：
2
3
51 52 51
4
51 52 52 51

3
4
*/
package main

import "fmt"

//85.71%
func main() {
	var T int
	fmt.Scan(&T)
	for ; T != 0; T-- {
		var n int
		fmt.Scan(&n)
		nums := make([]int, n)
		for i := range nums {
			fmt.Scan(&nums[i])
		}
		res := fun1(nums, n)
		fmt.Println(res)
	}
}

func fun1(nums []int, n int) int {
	res := 1

	fun11 := func(i, j int) {
		l, r := i-1, j+1
		for l >= 0 && r < n {
			if nums[l] == nums[r] && nums[l] <= nums[l+1] {
				l--
				r++
				continue
			}
			break

		}
		temp := r - l - 1
		if temp > res {
			res = temp
		}
	}

	for i := 0; i < n; i++ {
		fun11(i, i)
		if i+1 < n && nums[i] == nums[i+1] {
			//if res < 2 {
			//	res = 2
			//}
			fun11(i, i+1)
		}
	}

	return res
}
