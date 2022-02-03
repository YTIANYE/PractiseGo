/*
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Input: preorder = [-1], inorder = [-1]
Output: [-1]
*/

package main

import . "awesomeProject1/Algorithm/main/structure/TreeNode"

/*
执行用时：4 ms, 在所有 Go 提交中击败了93.08%的用户
内存消耗：3.7 MB, 在所有 Go 提交中击败了98.39%的用户
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{Val: preorder[0]}
	index := 0
	for node.Val != inorder[index] && index < len(inorder) {
		index++
	}

	node.Left = buildTree(preorder[1:index+1], inorder[:index])
	node.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return node
}
