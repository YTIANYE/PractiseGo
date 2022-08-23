/**
丑数：质因数只包含2,3,5的正整数

n= 10
输出12

1,2,3,4,5,6,8,9,10,12

1是丑数
n不超过1690
*/

package main

import fmt "fmt"

// import "sort"

func main() {
	n := 10
	fmt.Println(choushu(n))
}

func choushu(n int) int {
	//k:= 0
	nums := []int{1}
	if n == 1 {
		return 1
	}
	//
	for i := 0; i < n; i++ {
		x, y, z := jisuan(nums[i])
		nums = add(nums, x)
		nums = add(nums, y)
		nums = add(nums, z)
		//  fmt.Println(nums)
	}
	return nums[n-1]
}

func jisuan(num int) (int, int, int) {
	return num * 2, num * 3, num * 5
}

func add(nums []int, num int) []int {
	for i := 1; i < len(nums); i++ {
		if num == nums[i] { // 注意去重
			return nums
		}
		if num < nums[i] {
			nums = append(nums[:i], append([]int{num}, nums[i:]...)...)
			return nums
		}
	}
	nums = append(nums, num)
	return nums

}
