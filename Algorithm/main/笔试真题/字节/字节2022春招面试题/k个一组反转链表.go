/*
标题
单链表每隔k个元素做一次反转

题目描述
​
k 为正整数，小于等于链表的长度。 如果节点的数量不是 k 的倍数，那么最后剩下的节点应该保持原样。 ​

list: 1->2->3->4->5​

For k = 2, you should return: 2->1->4->3->5​

For k = 3, you should return: 3->2->1->4->5
*/

package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	// var a int
	// fmt.Scan(&a)
	// fmt.Printf("%s", "hello world")

	l5 := &ListNode{5, nil}
	l4 := &ListNode{4, l5}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	k := 2

	res := reverList(l1, k)

	for res != nil {
		fmt.Println(res.val)
		res = res.next
	}

}

func reverList(head *ListNode, k int) *ListNode {
	preHead := &ListNode{0, head}

	p := preHead.next
	pre := preHead
	for p != nil {
		q := pre // 注意错误：反转后的pre,变成了头，next， 不是p了
		for i := 0; i < k; i++ {
			if p != nil {
				pre = pre.next
				p = p.next
			} else {
				return preHead.next
			}
		}
		pre = reverse(q, p)
	}
	return preHead.next
}

// func reverse(prehead, end *ListNode)  {
// 	// fmt.Println(prehead.val, end.val)
// 	pre := prehead
// 	p := pre.next
// 	for p != end {
// 		q := p
// 		pre.next = p.next// 这里出现错误
// 		p = pre.next
// 		q.next = prehead.next
// 		prehead.next = q
// 	}
//
// }

func reverse(prehead, end *ListNode) *ListNode {
	lastone := prehead.next
	p := prehead.next
	prehead.next = end // 注意： 截断后头插
	for p != end {
		q := p
		p = p.next
		q.next = prehead.next
		prehead.next = q
	}
	return lastone
}
