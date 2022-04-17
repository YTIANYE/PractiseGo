/**
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

 

示例 1：



输入：root = [1,2,3,4,5,6,7]
输出：[1,#,2,3,#,4,5,6,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。
示例 2:

输入：root = []
输出：[]
 

提示：

树中节点的数量在 [0, 212 - 1] 范围内
-1000 <= node.val <= 1000
 

进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

*/

package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 我的题解：层次遍历
/**
执行用时：4 ms, 在所有 Go 提交中击败了90.04%的用户
内存消耗：6.5 MB, 在所有 Go 提交中击败了18.85%的用户

*/

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i < n-1; i++ {
			queue[i].Next = queue[i+1]
		}
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
