/**
蔚来面试题：
给定一个单链表的头节点 head,实现一个调整单链表的函数，使得每K个节点之间为一组进行逆序
并且从链表的尾部开始组起，头部剩余节点数量不够一组的不需要逆序。（不能使用队列或者栈作为辅助）
 */

package main

import "fmt"

type Node struct{
	Val int
	Next *Node
}

func reverse(head *Node , k int ) *Node {
	pHead := &Node{0, head}
	count := 0
	for p:=pHead.Next;p!= nil;p = p.Next {
		count++
	}
	m := count % k // 2
	pre := pHead
	q := pre.Next

	for i:=0;i<m;i++ { // pre a b q   m = 2
		q = q.Next
		pre = pre.Next
	}
	// pre q a b x
	for q != nil {
		for pre.Next != q {
			pre = pre.Next
		}
		//
		for i:=0;i<k;i++ { // k = 3
			q = q.Next
		}
		pre, q = rev(pre, q)
	}
	return pHead.Next
}

func rev(pre *Node, q *Node) (*Node, *Node) {
	ppre := pre.Next
	p := ppre.Next
	/**
	  // pre a b c q
	  // pre b a c q
	*/
	for p != q {
		ppre.Next = p.Next
		p.Next = pre.Next
		pre.Next = p
		p = ppre.Next
	}
	return pre, q
}

func main() {
	k := 3
	pHead := &Node{0, nil}
	p := pHead

	for i:=1;i<=8;i++ {
		node := &Node{i, nil}  // 括号错了，  node := &Node(i, nil)
		p.Next = node
		p = p.Next
	}

	// 打印
	head := pHead.Next
	p = head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
	// 翻转
	newHead := reverse(head, k)
	// 打印
	p = newHead
	fmt.Println("第二次打印")
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}