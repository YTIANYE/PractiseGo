/**

题目：
一个数n 转化成r进制数，如果转化结果是回文数，输出r
2<=r<=16
 */

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := []int{}
	for i:=2;i<=16;i++ {
		nums := toR(n, i)
		if isHui(nums) {
			res = append(res, i)
		}
	}
	//
	if len(res) == 0 {
		fmt.Println(-1)
	}else {
		for _, num := range res {
			fmt.Println(num)
		}
	}
}


func toR(n, i int) []int{
	nums := []int{}
	for n != 0 {
		num := n % i
		nums = append(nums, num)
		n /= i
	}
	return nums
}

/**
	6	0
	3	1
	1	1
 */


func isHui(nums []int) bool {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] != nums[j]{
			return false
		}
		i++
		j--
	}
	return true
}