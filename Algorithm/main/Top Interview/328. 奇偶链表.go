/*

给定单链表的头节点head，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
第一个节点的索引被认为是 奇数 ， 第二个节点的索引为偶数 ，以此类推。
请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
你必须在O(1)的额外空间复杂度和O(n)的时间复杂度下解决这个问题。

示例 1:
输入: head = [1,2,3,4,5]
输出:[1,3,5,2,4]
示例 2:
输入: head = [2,1,3,5,6,4,7]
输出: [2,3,6,7,1,5,4]

提示:
n == 链表中的节点数
0 <= n <= 104
-106<= Node.val <= 106
*/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import (
	. "awesomeProject1/Algorithm/main/structure/LinkList"
)

// 我的题解：
/*
执行用时：4 ms, 在所有 Go 提交中击败了60.61%的用户
内存消耗：3.1 MB, 在所有 Go 提交中击败了85.53%的用户
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p, q := head, head.Next
	res := ListNode{0, head}
	pre := head.Next
	index := true // 奇数

	for p.Next != nil && q.Next != nil {
		if index {
			p.Next = q.Next
			p = p.Next
		} else {
			q.Next = p.Next
			q = q.Next
		}
		index = !index
	}
	q.Next = nil // 避免造成循环链表
	p.Next = pre

	return res.Next
}

// 官方题解：简洁版
func oddEvenList1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil { // 这样的判断条件比较省代码
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := CreatList(nums)
	res := oddEvenList(head)
	PrintList(res)
}
