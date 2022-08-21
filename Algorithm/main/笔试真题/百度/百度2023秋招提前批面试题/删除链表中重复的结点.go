/**
在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。 例如，链表 1->2->3->3->4->4->5  处理后为 1->2->5

数据范围：链表长度满足 0 \le n \le 1000 \0≤n≤1000  ，链表中的值满足 1 \le val \le 1000 \1≤val≤1000

进阶：空间复杂度 O(n)\O(n)  ，时间复杂度 O(n) \O(n)

例如输入{1,2,3,3,4,4,5}时，对应的输出为{1,2,5}


输入： {1,2,3,3,4,4,5}
输出： {1,2,5}

输入： {1,1,1,8}
输出： {8}

输入： {1,1}
输出： {}

输入： {}
输出： {}
 */

package main
// import . "nc_tools"

 type ListNode struct{
   Val int
   Next *ListNode
 }


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pHead ListNode类
 * @return ListNode类
 */

// 经过四五次修改才通过
func deleteDuplication( pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil || pHead.Next == nil { // 注意开始判断
		return pHead
	}
	head := &ListNode{Val: 0}
	head.Next = pHead

	p := pHead.Next
	pre := pHead
	ppre := head

	for p != nil {
		if p.Val == pre.Val {
			for p != nil && p.Val == pre.Val {
				p = p.Next
			}
			pre = p
			ppre.Next = pre
			if p == nil { // 注意最后一个判断
				break
			}
			p = pre.Next
		}else {
			p = p.Next
			pre = pre.Next
			ppre = ppre.Next
		}

	}
	return head.Next
}