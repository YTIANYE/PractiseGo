/*
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000
*/
package main

import . "awesomeProject1/github.com/YTIANYE/Algorithm/main/structure/LinkList"

/*我的题解：不改变原链表*/
/*执行用时：8 ms, 在所有 Go 提交中击败了30.82%的用户
内存消耗：4.5 MB, 在所有 Go 提交中击败了5.15%的用户*/

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	l := list
	p, q := l1, l2
	for q != nil && p != nil {
		if p.Val < q.Val {
			l.Next = &ListNode{Val: p.Val}
			p = p.Next
		} else {
			l.Next = &ListNode{Val: q.Val}
			q = q.Next
		}
		l = l.Next
	}
	for p != nil {
		l.Next = &ListNode{Val: p.Val}
		p = p.Next
		l = l.Next
	}
	for q != nil {
		l.Next = &ListNode{Val: q.Val}
		q = q.Next
		l = l.Next
	}
	return list.Next

}

// 我的题解：改变原链表
/*执行用时：4 ms, 在所有 Go 提交中击败了93.95%的用户
内存消耗：4.1 MB, 在所有 Go 提交中击败了100.00%的用户*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	list := &ListNode{}
	l := list
	// 合并
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}
	// 收尾
	if l1 != nil {
		l.Next = l1
	}
	if l2 != nil {
		l.Next = l2
	}
	return list.Next
}
