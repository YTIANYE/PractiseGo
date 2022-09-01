package main


type Node struct {
	Val int
	Next *Node
}

func rev(prehead, head *Node, k int) (*Node, *Node) {
	p := head
	q := p
	for i:=0;i<k;i++ {
		if p == nil {
			return prehead.Next, p
		}else {
			p = p.Next
		}
	}
	//
	prehead.Next = p
	for p!= q {
		node := q
		q = q.Next
		node.Next = prehead.Next
		prehead.Next = node
	}
	//
	pre := prehead
	for i:= 0 ;i<k;i++ {
		pre= pre.Next
	}
	return pre, p
}

func revers(head *Node , k int) *Node {
	pHead := &Node{0, head}
	pre := pHead
	p := head
	for p != nil {
		pre, p = rev(pre, p, k )
	}
	return pHead.Next
}