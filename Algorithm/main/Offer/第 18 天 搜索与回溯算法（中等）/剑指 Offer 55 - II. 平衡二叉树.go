package main

import . "awesomeProject1/Algorithm/main/structure/TreeNode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 我的题解:自顶向下遍历
/*
执行用时：4 ms, 在所有 Go 提交中击败了95.95%的用户
内存消耗：5.5 MB, 在所有 Go 提交中击败了98.38%的用户
*/

func isBalanced1(root *TreeNode) bool {
	if root == nil{
		return true
	}
	temp := depth(root.Left) - depth(root.Right)
	if temp > 1 || temp < -1{
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int{
	if root == nil{
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1

}

func max(a, b int) int{
	if a>b{
		return a
	}else{
		return b
	}
}


// 我实现的官方题解:自底向上遍历
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：5.5 MB, 在所有 Go 提交中击败了99.03%的用户
*/
func isBalanced(root *TreeNode) bool{
	return height(root) >= 0
}

func height(root *TreeNode) int{
	if root == nil{
		return 0
	}

	hleft := height(root.Left)
	hright := height(root.Right)
	temp := hright-hleft
	if hright == -1 || hleft == -1 || temp > 1 || temp < -1{
		return -1
	}
	return max(hright, hleft) + 1
}