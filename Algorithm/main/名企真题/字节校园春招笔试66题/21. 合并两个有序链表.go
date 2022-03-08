/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：


输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]
提示：

两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列

*/

package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import . "awesomeProject1/Algorithm/main/structure/LinkList"

// 我的题解：迭代
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.4 MB, 在所有 Go 提交中击败了87.12%的用户
*/

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil{
		return list2
	}
	if list2 == nil{
		return list1
	}

	res := &ListNode{0, nil }
	p := res
	for list1 != nil && list2 != nil{
		var q *ListNode
		if list1.Val < list2.Val{
			q = list1
			list1 = list1.Next
		}else{
			q = list2
			list2 = list2.Next
		}
		q.Next = nil
		p.Next = q
		p = p.Next
	}
	if list1 != nil{
		p.Next = list1
	}
	if list2 != nil{
		p.Next = list2
	}
	return res.Next
}

