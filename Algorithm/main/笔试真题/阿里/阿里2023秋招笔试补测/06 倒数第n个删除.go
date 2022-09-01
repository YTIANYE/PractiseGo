package main

type Node1 struct {
	 Val int
	 Next *Node1
}

func remove(head *Node1, n int) *Node1 {
	pHead := &Node1{0, head }
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
