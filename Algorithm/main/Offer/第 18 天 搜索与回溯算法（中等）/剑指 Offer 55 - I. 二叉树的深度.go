/*
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。



提示：

节点总数 <= 10000

*/

package main

import (
	. "awesomeProject1/Algorithm/main/structure/TreeNode"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 我的题解：先序遍历
/*
执行用时：4 ms, 在所有 Go 提交中击败了89.58%的用户
内存消耗：4.1 MB, 在所有 Go 提交中击败了97.61%的用户
*/

func maxDepth1(root *TreeNode) int {
	depth := 0
	return preOrder(root, depth)

}

func preOrder(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth += 1
	return max(preOrder(root.Left, depth), preOrder(root.Right, depth))

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 我实现的精选题解：后序遍历
/*
执行用时：4 ms, 在所有 Go 提交中击败了89.58%的用户
内存消耗：4 MB, 在所有 Go 提交中击败了99.72%的用户
*/

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 我的题解：层次遍历

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：4.1 MB, 在所有 Go 提交中击败了97.75%的用户
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0
	for len(queue) != 0 { // 用 != nil会一直循环，需要用 != 0
		tempqueue := []*TreeNode{}
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node.Left != nil {
				tempqueue = append(tempqueue, node.Left)
			}
			if node.Right != nil {
				tempqueue = append(tempqueue, node.Right)
			}
		}
		queue = tempqueue
		depth++
	}
	return depth
}

func main() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := TreeCreate(nums, 0)
	// TreePrint(root)

	fmt.Println(maxDepth(root))
}
