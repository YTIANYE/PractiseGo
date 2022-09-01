package main

type Node struct {
	 Val int
	 Next *Node
}

func remove(head *Node, n int) *Node {
	pHead := &Node{0, head }
	pre := pHead
	p := pre.Next
	i := 0
	for i <n {//n = 2   pre a a p
		i++
		p = p.Next
	}
	for p != nil {
		p = p.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return pHead.Next
}
