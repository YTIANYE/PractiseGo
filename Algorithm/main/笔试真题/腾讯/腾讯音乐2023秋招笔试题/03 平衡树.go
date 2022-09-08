/**
一个树的结点出度可以为0 或 2
给每个结点赋值，使得树中任意一个结点的左子树的权值和与右子树的权值和相同，
这棵树最少的权值和是多少
结果对1000000007取模
 */

/**
{0,0,0}
3
 */
package main

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param tree TreeNode类
 * @return int整型
 */
// 55%
func getTreeSum(tree *TreeNode) int {
	// write code here
	// 每个结点赋值
	setVal(tree)

	if tree == nil {
		return 0
	}
	left := getTreeSum(tree.Left)
	right := getTreeSum(tree.Right)
	if left == right {
		tree.Val = (tree.Val + left + right) % 1000000007
	} else {
		tree.Val = (tree.Val + max(left, right)*2) % 1000000007
	}
	return tree.Val
}

func setVal(tree *TreeNode) {
	if tree == nil {
		return
	}
	tree.Val = 1
	setVal(tree.Left)
	setVal(tree.Right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**

	1
  1	   3
1	1


*/
