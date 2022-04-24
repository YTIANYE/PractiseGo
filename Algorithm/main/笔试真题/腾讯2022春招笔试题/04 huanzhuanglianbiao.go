/*
题目：一个环形链表，顺序遍历出若干个子链表，存储在a中
根据a还原出原来的环形链表，并将链表从某处切开，返回切开后的链表
要求返回的链表顺序遍历的结果或者逆序遍历的结果
的字典序是所有切开链表中最小的
字典序：如环形链表1,2,3,1...中，最小的为1,2,3
*/

package main

import . "awesomeProject1/Algorithm/main/structure/LinkList"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param a ListNode类一维数组 指向每段碎片的开头
 * @return ListNode类
 */

func solve(a []*ListNode) *ListNode {
	// write code here

	var head *ListNode
	head = a[0] // 环形链表的一个头
	hash := make(map[int]*ListNode)

	t := a[0]
	a = a[1:]
	for t != nil {
		hash[t.Val] = t
		t = t.Next
	}

	for len(a) != 0 { // 存在段落
		for i := 0; i < len(a); i++ { // 遍历每一个段落
			q := a[i]               // 每一段的头结点
			if hash[q.Val] != nil { // q存在于链表head中，可以链接
				p := hash[q.Val]
				for p.Next != nil && q.Next != nil { // 不是||，会报错，当时就是这里错了没A出来
					p = p.Next
					q = q.Next
				}
				if p.Next == nil { // a[i]中有结点可以续接上
					p.Next = q.Next
					q = q.Next
					for q != nil { // 续接上a[i]多出来的结点
						if q.Next != nil && q.Next.Val == head.Val { // 注意环链表head的连接
							q.Next = head
							break
						}
						hash[q.Val] = q // 注意维护hash ,思考很久才发现这个问题
						q = q.Next
					}
				}
				// 从 a 中去除该序列
				if i == len(a)-1 {
					a = a[:i]
				} else {
					a = append(a[:i], a[i+1:]...)
				}
				break
			}
		}
	}

	// 记录环形链表中的最小值
	MIN := head.Val
	for p := head.Next; p != head; p = p.Next {
		MIN = min(MIN, p.Val)
	}

	// 断开
	head = hash[MIN]
	p := head
	for p.Next != head {
		p = p.Next
	}
	p.Next = nil
	return head
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 构建a
	var a []*ListNode
	lists := [][]int{
		{1, 2, 3},
		{2, 3, 4, 5, 6},
		{6, 7, 8},
		{8, 9, 10, 1, 2},
		{3, 4, 5, 6, 7},
	}
	for _, list := range lists {
		h := CreatList(list)
		a = append(a, h)
	}

	// 计算结果
	res := solve(a)
	PrintList(res)
}
