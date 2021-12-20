package LinkList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreatList(nums []int) *ListNode {

	head := &ListNode{}
	p := head
	for _, val := range nums {
		node := &ListNode{Val: val}
		p.Next = node
		p = p.Next
	}
	return head.Next
}

func PrintList(head *ListNode) {
	p := head
	result := []int{}
	for p != nil {
		result = append(result, p.Val)
		p = p.Next
	}
	fmt.Println(result)
}

func DeleteNode(head *ListNode, val int) *ListNode {
	prehead := &ListNode{Next: head} // 原链表前加一个头结点
	p := prehead
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			break
		}
		p = p.Next
	}
	return prehead.Next
}

// func main(){
// 	nums := []int{1,2,3}
// 	list := CreatList(nums)
// 	PrintList(list)
// 	list = DeleteNode(list, 2)
// 	PrintList(list)
// }
