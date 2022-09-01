package main

import "fmt"

type Node struct {
	Val int
	Next *Node
}

func rever(pHead *Node) *Node{
	if pHead == nil {
		return pHead
	}

	preHead := &Node{0, pHead}
	p := pHead.Next
	pre := pHead
	for p != nil {
		pre.Next = p.Next
		p.Next = preHead.Next
		preHead.Next = p
		p = pre.Next
	}
	return preHead.Next
}

func main() {
	c := &Node{3, nil}
	// b := &Node{2, c}
	//a := &Node{1, b}
	res := rever(c)
	fmt.Println(res)
}
