package main

import (
	"awesomeProject1/Algorithm/main/structure/TreeNode"
	"fmt"
)

/*

从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
例如:
给定二叉树:[3,9,20,null,null,15,7],
3
/ \
9  20
/  \
15   7
返回其层次遍历结果：
[
[3],
[9,20],
[15,7]
]
提示：
节点总数 <= 1000
*/

/*我的题解*/

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.8 MB, 在所有 Go 提交中击败了23.39%的用户

func levelOrder2(root *TreeNode.TreeNode) [][]int {
	var result [][]int
	var queue []*TreeNode.TreeNode
	if root == nil {
		return nil
	}

	queue = append(queue, root)

	for len(queue) != 0 {
		current := make([]*TreeNode.TreeNode, len(queue)) // 每一层结点存储在一个队列中
		copy(current, queue)
		queue = []*TreeNode.TreeNode{}
		var r []int
		for len(current) != 0 { // 循环条件换做len(queue)可以不需要current队列
			node := current[0]
			current = current[1:]
			r = append(r, node.Val)
			if node.Left != nil { // 注意判空
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		result = append(result, r)
	}
	return result
}

func main() {
	nums := []int{0, 1, 2, 3, -1, 5, 6}
	root := TreeNode.TreeCreate(nums, 0)
	TreeNode.TreePrint(root)
	fmt.Println(" ")
	fmt.Println(levelOrder(root))
}
