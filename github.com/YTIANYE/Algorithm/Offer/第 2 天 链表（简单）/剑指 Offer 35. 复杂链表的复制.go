package main

/*
请实现 copyRandomList 函数，复制一个复杂链表。
在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
还有一个 random 指针指向链表中的任意节点或者 null。

提示：
-10000 <= Node.val <= 10000
Node.random 为空（null）或指向链表中的节点。
节点数目不超过 1000 。
*/

// 输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
//
// 输入：head = [[1,1],[2,1]]
// 输出：[[1,1],[2,1]]
//
// 输入：head = [[3,null],[3,0],[3,null]]
// 输出：[[3,null],[3,0],[3,null]]
//
// 输入：head = []
// 输出：[]
// 解释：给定的链表为空（空指针），因此返回 null。

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：3.5 MB, 在所有 Go 提交中击败了62.36%的用户

//我的题解：hashmap
func copyRandomList(head *Node) *Node {
	prehead := &Node{}
	p := head
	q := prehead	//空的头指针
	nodesmap := make(map[*Node]*Node)
	for p != nil {
		node := &Node{Val: p.Val, Random: nil}
		q.Next = node
		nodesmap[p] = node
		p = p.Next
		q = q.Next
	}
	p = head
	q = prehead.Next
	for p != nil {
		q.Random = nodesmap[p.Random]
		p = p.Next
		q = q.Next
	}
	return prehead.Next
}
