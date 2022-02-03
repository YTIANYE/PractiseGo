/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]

示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。

*/

package main

import (
	. "awesomeProject1/Algorithm/main/structure/TreeNode"
	"fmt"
)

// 我实现的官方题解：递归
/*
执行用时：8 ms, 在所有 Go 提交中击败了95.75%的用户
内存消耗：7.4 MB, 在所有 Go 提交中击败了97.97%的用户
*/

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}

}

// 我的题解：

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode{

	parent := make(map[*TreeNode]*TreeNode)
	visit := make(map[*TreeNode]bool)

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil{
			return
		}


		if root.Left != nil{
			parent[root.Left] = root
			dfs(root.Left)
		}
		if root.Right != nil{
			parent[root.Right] = root
			dfs(root.Right)
		}
	}
	dfs(root)

	for p != nil{
		visit[p] = true
		p = parent[p]
	}
	for q != nil{
		if visit[q]== true{
			return q
		}
		q = parent[q]
	}
	return nil
}


func main() {
	nums := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	root := TreeCreate(nums, 0)
	// TreePrint(root)
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Left.Right.Right).Val)
}
