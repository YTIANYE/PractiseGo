/*
给定一个二叉树的根节点 root ，返回它的 中序 遍历。

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100
*/

package main

import (
	. "awesomeProject1/Algorithm/main/structure/TreeNode"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了99.97%的用户
*/

func inorderTraversal1(root *TreeNode) (res []int) {

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil { // 写在这里一次判断比较简单
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}


// 我实现的官方题解：Morris 中序遍历
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了86.69%的用户
*/

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	x := root

	for x != nil {
		if x.Left == nil {
			res = append(res, x.Val)
			x = x.Right
		} else {
			predecessor := x.Left
			for predecessor.Right != nil && predecessor.Right != x {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = x
				x = x.Left
			} else { // predecessor.Right == x
				res = append(res, x.Val)
				// 恢复原样
				predecessor.Right = nil
				x = x.Right
			}
		}

	}
	return res
}
