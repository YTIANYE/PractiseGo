/**
给定两个非空树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。
示例 1:
给定的树 s:
     3
    / \
   4   5
  / \
 1   2

给定的树 t：
   4
  / \
 1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

示例 2:
给定的树 s：
     3
    / \
   4   5
  / \
 1   2
    /
   0

给定的树 t：
   4
  / \
 1   2
返回 false。

*/

package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func have(s, t *Node) bool {
	if s == nil {
		return false
	}
	//
	if isSame(s, t) {
		return true
	}
	return have(s.Left, t) || have(s.Right, t)
}

func isSame(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil {
		return false
	} else if b == nil {
		return false
	}
	//
	// fmt.Println(a.Val, b.Val)
	if a.Val != b.Val {
		return false
	}
	return isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
}

func main() {
	node1 := &Node{1, nil, nil}
	node2 := &Node{2, nil, nil}
	node4 := &Node{4, node1, node2}
	node5 := &Node{5, nil, nil}
	node3 := &Node{3, node4, node5}

	//node0 := &Node{0, nil, nil}
	//node1 := &Node{1, nil, nil}
	//node2 := &Node{2, node0, nil}
	//node4 := &Node{4, node1, node2}
	//node5 := &Node{5, nil, nil}
	//node3 := &Node{3, node4, node5}

	node11 := &Node{1, nil, nil}
	node22 := &Node{2, nil, nil}
	node44 := &Node{4, node11, node22}

	fmt.Println(have(node3, node44))
}
