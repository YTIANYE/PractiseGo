/**
两个链表分别为两个数字的二进制
a 由头到尾 表示一个数的二进制
b 由尾到头 表示一个数的二进制

求这两个数异或后的数是多少？ 用链表二级制的形式表示
*/
package main

//import "fmt"
//import . "nc_tools"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a ListNode类
 * @param b ListNode类
 * @return ListNode类
 */
// 90.91%
func xorList(a *ListNode, b *ListNode) *ListNode {
	// write code here
	a = reverse(a)
	preHead := &ListNode{0, nil}
	t := preHead
	p, q := a, b
	for p != nil && q != nil {
		val := p.Val ^ q.Val
		node := &ListNode{val, nil}
		//t.Next = node
		//t = t.Next
		node.Next = t.Next
		t.Next = node

		p = p.Next
		q = q.Next
	}
	for p != nil {
		val := p.Val
		node := &ListNode{val, nil}
		node.Next = t.Next
		t.Next = node
		p = p.Next
	}
	for q != nil {
		val := q.Val
		node := &ListNode{val, nil}
		node.Next = t.Next
		t.Next = node
		q = q.Next
	}
	// 去除前导0
	t = preHead.Next
	for t != nil && t.Val == 0 {
		preHead.Next = t.Next
		t = preHead.Next
	}

	return preHead.Next

}

func reverse(b *ListNode) *ListNode {
	preHead := &ListNode{0, b}
	pre := b
	p := b.Next
	for p != nil {
		pre.Next = p.Next
		p.Next = preHead.Next
		preHead.Next = p
		p = pre.Next
	}
	return preHead.Next
}
