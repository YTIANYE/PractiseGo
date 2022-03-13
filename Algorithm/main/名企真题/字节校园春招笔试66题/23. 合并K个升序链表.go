/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。


示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]

提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4

*/

package main

import (
	. "awesomeProject1/Algorithm/main/structure/LinkList"
)

// 我的题解：自底向上+归并合并
/*
执行用时：4 ms, 在所有 Go 提交中击败了99.22%的用户
内存消耗：5.4 MB, 在所有 Go 提交中击败了49.09%的用户
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0  {
		return nil
	}

	for len(lists) != 1 {
		var newLists []*ListNode
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				newLists = append(newLists, merge2Lists(lists[i], lists[i+1]))
			} else {
				newLists = append(newLists, lists[i])
			}
		}
		lists = newLists
	}
	return lists[0]
}

func merge2Lists(first, second *ListNode) *ListNode {
	if first == nil {
		return second
	}
	if second == nil {
		return first
	}

	preHead := &ListNode{-1, nil}
	p := preHead
	for first != nil && second != nil {
		if first.Val < second.Val {
			p.Next = first
			first = first.Next
		} else {
			p.Next = second
			second = second.Next
		}
		p = p.Next
		p.Next = nil
	}
	if first != nil {
		p.Next = first
	}
	if second != nil {
		p.Next = second
	}
	return preHead.Next
}

func main() {
	nums := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}
	lists := []*ListNode{}
	for _, num := range nums {
		head := CreatList(num)
		lists = append(lists, head)
	}
	res := mergeKLists(lists)
	PrintList(res)

}
