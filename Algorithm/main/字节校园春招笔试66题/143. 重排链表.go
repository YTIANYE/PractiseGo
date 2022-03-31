/**
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

 

示例 1：



输入：head = [1,2,3,4]
输出：[1,4,2,3]
示例 2：



输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]
 

提示：

链表的长度范围为 [1, 5 * 104]
1 <= node.val <= 1000


*/

package main

import . "awesomeProject1/Algorithm/main/structure/LinkList"


// 两次遍历——我的题解
/*
执行用时：8 ms, 在所有 Go 提交中击败了85.99%的用户
内存消耗：5.1 MB, 在所有 Go 提交中击败了81.62%的用户
*/

func reorderList(head *ListNode) {
	p := head
	q := head
	// 1 2 3 4
	for p != nil && p.Next != nil {
		p = p.Next.Next
		q = q.Next
	}
	mid := q
	q = q.Next
	mid.Next = nil // 断开列表
	head2 := &ListNode{0, nil}
	for q != nil {
		temp := q.Next
		q.Next = head2.Next
		head2.Next = q
		q = temp
	}
	m := head
	for head2.Next != nil {
		n := head2.Next
		head2.Next = n.Next
		n.Next = m.Next
		m.Next = n
		m = m.Next.Next
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	head := CreatList(nums)
	reorderList(head)
	PrintList(head)
}
