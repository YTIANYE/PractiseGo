package main

import "awesomeProject1/github.com/YTIANYE/Algorithm/main/structure/LinkList"

/*我的题解*/

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.5 MB, 在所有 Go 提交中击败了100.00%的用户

func reverseList1(head *LinkList.ListNode) *LinkList.ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	h := &LinkList.ListNode{Next: head}
	p := head.Next
	for p != nil{
		head.Next = p.Next
		p.Next = h.Next
		h.Next = p
		p = head.Next
	}
	return h.Next
}

//递归
func reverseList(head *LinkList.ListNode) *LinkList.ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}



func main(){
	nums := []int{1,2,3}
	head := LinkList.CreatList(nums)
	LinkList.PrintList(head)
	head = reverseList(head)
	LinkList.PrintList(head)
}