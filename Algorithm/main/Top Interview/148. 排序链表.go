/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。


示例 1：

输入：head = [4,2,1,3]
输出：[1,2,3,4]
示例 2：


输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
示例 3：

输入：head = []
输出：[]

*/
package main

import (
	. "awesomeProject1/Algorithm/main/structure/LinkList"
)

// 1. 我的题解：暴力超时

func sortList1(head *ListNode) *ListNode {
	preHead := &ListNode{-1, nil}
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	for p != nil {
		q := p
		p = p.Next
		pre := preHead
		for pre.Next != nil && pre.Next.Val < q.Val {
			pre = pre.Next
		}
		q.Next = pre.Next
		pre.Next = q
	}
	return preHead.Next
}

// 2. 我的题解:哈希表
/*
执行用时：44 ms, 在所有 Go 提交中击败了13.45%的用户
内存消耗：12.8 MB, 在所有 Go 提交中击败了5.08%的用户
*/
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	hashMap := make(map[int][]*ListNode)
	nums := []int{}
	p := head
	for p != nil {
		hashMap[p.Val] = append(hashMap[p.Val], p)
		p = p.Next
	}
	for k := range hashMap {
		nums = append(nums, k)
	}

	// sort.Ints(nums)
	nums = sortFast(nums, 0, len(nums)-1)

	preHead := &ListNode{Val: -1}
	pre := preHead
	for _, num := range nums {
		for _, p := range hashMap[num] {
			pre.Next = p
			pre = pre.Next
			pre.Next = nil
		}
	}
	return preHead.Next
}

// 数组的快速排序
func sortFast(nums []int, left, right int) []int {
	if left >= right {
		return nums
	}

	i, j := left+1, right
	for i <= j {
		for j >= i && nums[j] > nums[left] {
			j--
		}
		for j >= i && nums[i] < nums[left] {
			i++
		}
		if j >= i {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left], nums[j] = nums[j], nums[left]

	nums = sortFast(nums, left, j-1)
	nums = sortFast(nums, j+1, right)
	return nums
}

// 3. 我实现的官方题解：归并排序、自底向上
/*
执行用时：28 ms, 在所有 Go 提交中击败了62.42%的用户
内存消耗：7.3 MB, 在所有 Go 提交中击败了32.17%的用户
*/
func sortList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}

	preHead := &ListNode{-1, head}
	for limit := 1; limit < count; limit *= 2 {
		pre := preHead
		p := preHead.Next
		for p != nil {
			left := p
			for i := 1; i < limit && p != nil && p.Next != nil; i++ {
				p = p.Next
			}
			right := p.Next
			p.Next = nil
			p = right
			for i := 1; i < limit && p != nil && p.Next != nil; i++ {
				p = p.Next
			}

			var last *ListNode
			if p != nil {
				last = p.Next
				p.Next = nil
			}
			p = last

			pre.Next = merge(left, right)
			for pre.Next != nil {
				pre = pre.Next
			}
		}
	}
	return preHead.Next
}

// 链表合并
func merge(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	preHead := &ListNode{-1, nil}
	p := preHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}

	if left != nil {
		p.Next = left
	}
	if right != nil {
		p.Next = right
	}
	return preHead.Next
}

// 4. 我实现的官方题解：归并排序、自顶向下
/*
执行用时：24 ms, 在所有 Go 提交中击败了91.44%的用户
内存消耗：7.1 MB, 在所有 Go 提交中击败了80.77%的用户
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	preHead := &ListNode{-1, head}
	fast := preHead
	slow := preHead
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	second := slow.Next
	slow.Next = nil
	left := sortList(preHead.Next)
	right := sortList(second)
	return merge(left, right)
}

func main() {
	nums := []int{4, 2, 1, 9, 6, 8, 5, 0, 3}
	head := CreatList(nums)
	res := sortList(head)
	PrintList(res)
}
