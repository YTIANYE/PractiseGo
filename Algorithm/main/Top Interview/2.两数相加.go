/*
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0开头。

示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

提示：
每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零
*/

package main

import (
	. "awesomeProject1/Algorithm/main/structure/LinkList"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
执行用时：8 ms, 在所有 Go 提交中击败了84.27%的用户
内存消耗：4.4 MB, 在所有 Go 提交中击败了81.45%的用户
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num := 0 // 进位
	p, q := l1, l2
	res := ListNode{0, nil}
	r := &res
	for p != nil && q != nil {
		temp := p.Val + q.Val + num
		num = temp / 10

		node := ListNode{temp % 10, nil}
		r.Next = &node
		r = r.Next
		p = p.Next
		q = q.Next
	}

	for p != nil {
		temp := p.Val + num
		num = temp / 10
		node := ListNode{temp % 10, nil}
		r.Next = &node
		r = r.Next
		p = p.Next
	}

	for q != nil {
		temp := q.Val + num
		num = temp / 10
		node := ListNode{temp % 10, nil}
		r.Next = &node
		r = r.Next
		q = q.Next
	}

	// 注意处理最后的进位 num
	if num != 0 {
		node := ListNode{num, nil}
		r.Next = &node
	}

	return res.Next

}
