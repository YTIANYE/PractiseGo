/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
 

示例 1：


输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
示例 2：

输入：head = [5], left = 1, right = 1
输出：[5]
 

提示：

链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n
 

进阶： 你可以使用一趟扫描完成反转吗？

*/

package main

import . "awesomeProject1/Algorithm/main/structure/LinkList"

// 我的题解：
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了72.81%的用户
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	preHead := &ListNode{0, head}
	pre, p := preHead, head
	i:=1
	for ;i<left;i++{
		p = p.Next
		pre = pre.Next
	}
	q := p
	p = p.Next
	for ;i<right;i++{
		q.Next = p.Next
		p.Next = pre.Next
		pre.Next = p
		p = q.Next
	}
	return preHead.Next
}