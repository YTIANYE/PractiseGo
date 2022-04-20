/*
题目：给出一个数n，删除k个数，使得结果的值最大
*/

package main

import (
	"fmt"
	"strconv"
)

/**
思路：每次从前往后遍历，删除的数nums[i]，应该满足条件，nums[i] < nums[i+1]
*/

func main() {

	num := 73919485
	k := 3
	r := remove(num, k)
	fmt.Println(r)

}

func remove(num int, k int) int {
	res := strconv.Itoa(num)

	for i := 0; i < k; i++ {

		for j := 0; j < len(res)-1; j++ {
			if res[j] < res[j+1] {
				res = res[:j] + res[j+1:]
				break
			}
		}
	}

	r, _ := strconv.Atoi(res)
	return r

}
