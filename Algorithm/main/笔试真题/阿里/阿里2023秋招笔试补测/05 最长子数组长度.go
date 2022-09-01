package main

import "fmt"

func fun(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			res = max(res, changdu(nums[i:j+1]))
		}
	}
	return res
}

func changdu(nums []int) int {

	hash := make(map[int]bool)
	for i := range nums {
		hash[nums[i]] = true
	}
	if len(hash) == len(nums) {
		return len(hash)
	}
	return 0

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 3, 3, 3, 9}
	fmt.Println(fun(nums))
}
