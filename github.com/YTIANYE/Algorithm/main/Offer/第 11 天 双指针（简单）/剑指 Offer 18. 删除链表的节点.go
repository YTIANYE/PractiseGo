/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。

注意：此题对比原题有改动

示例 1:

输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为5的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:

输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为1的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

说明：

题目保证链表中节点的值互不相同
若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点
*/
package main

import . "awesomeProject1/github.com/YTIANYE/Algorithm/main/structure/LinkList"

/*我的题解：附加头结点*/

/*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.8 MB, 在所有 Go 提交中击败了59.63%的用户*/

func deleteNode(head *ListNode, val int) *ListNode {

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

func main() {
	nums := []int{4, 5, 1, 9}
	head := CreatList(nums)
	PrintList(head)
	head = deleteNode(head, 5)
	PrintList(head)
}
