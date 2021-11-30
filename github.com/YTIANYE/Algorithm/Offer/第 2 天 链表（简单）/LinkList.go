package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func CreatList(nums []int) *ListNode {

	head := &ListNode{}
	p := head
	for _,val := range nums{
		node := &ListNode{Val: val}
		p.Next = node
		p = p.Next
	}
	return head.Next
}

func PrintList(head *ListNode){
	p := head
	for p != nil{
		fmt.Println(p.Val)
		p = p.Next
	}
}

// func main(){
// 	nums := []int{1,2,3}
// 	list := CreatList(nums)
// 	PrintList(list)
// }

