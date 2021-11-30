package main

import (
	"fmt"
)

/*我的题解*/
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：3.1 MB, 在所有 Go 提交中击败了90.96%的用户

func reversePrint(head *ListNode) []int {
	var nums []int

	p := head
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}

	for i:=0; i<len(nums)/2; i++{
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	return nums

	// 也可以用栈
}

func main() {
	nums := []int{1,2,3}
	list := CreatList(nums)
	PrintList(list)
	fmt.Println(reversePrint(list))
}
