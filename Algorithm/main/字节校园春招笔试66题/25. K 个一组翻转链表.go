/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：

你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 

示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
示例 2：


输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]
示例 3：

输入：head = [1,2,3,4,5], k = 1
输出：[1,2,3,4,5]
示例 4：

输入：head = [1], k = 1
输出：[1]
提示：

列表中节点的数量在范围 sz 内
1 <= sz <= 5000
0 <= Node.val <= 1000
1 <= k <= sz

*/

package main

import . "awesomeProject1/Algorithm/main/structure/LinkList"

// 我的题解
/*
执行用时：4 ms, 在所有 Go 提交中击败了90.42%的用户
内存消耗：3.4 MB, 在所有 Go 提交中击败了100.00%的用户
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	preHead := &ListNode{-1, head}
	pre := preHead
	p := head
	for p != nil {
		pre, p = reverse(pre, p, k)
	}
	return preHead.Next
}

func reverse(preHead, head *ListNode, k int) (*ListNode, *ListNode) {
	p, q := head, head

	for i := 0; i < k; i++ {
		if p == nil {
			return preHead.Next, p
		}
		p = p.Next
	}
	preHead.Next = p // 截断链表
	// 头插入过程
	for q != p {
		node := q
		q = q.Next
		node.Next = preHead.Next
		preHead.Next = node
	}
	// 查找 p的前一个结点
	pre := preHead
	for i := 0; i < k; i++ {
		pre = pre.Next
	}
	return pre, p
}

func main() {
	// nums := []int{1,2,3,4,5}
	// k:= 2
	nums := []int{1, 2}
	k := 2
	list := CreatList(nums)
	newList := reverseKGroup(list, k)
	PrintList(newList)

}
