/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：


输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]

提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000


*/

package main

import . "awesomeProject1/Algorithm/main/structure/LinkList"

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.4 MB, 在所有 Go 提交中击败了63.85%的用户

*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head.Next
	pre := head
	pre.Next = nil
	q := p
	for p != nil {
		q = p
		p = p.Next
		q.Next = pre
		pre = q
	}

	return pre

}
