package main

import "fmt"

func mySort(nums []int) []int {
	n := len(nums)
	for i:=0;i<n;i++ {
		for j:=0;j<n-1-i;j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums

}


func main() {
	nums := []int{2,5,1}
	res := mySort(nums)
	fmt.Println(res)
}
