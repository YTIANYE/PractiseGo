/*
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。

 

示例 1：


输入：root = [1,2,3]
输出：25
解释：
从根到叶子节点路径 1->2 代表数字 12
从根到叶子节点路径 1->3 代表数字 13
因此，数字总和 = 12 + 13 = 25
示例 2：


输入：root = [4,9,0,5,1]
输出：1026
解释：
从根到叶子节点路径 4->9->5 代表数字 495
从根到叶子节点路径 4->9->1 代表数字 491
从根到叶子节点路径 4->0 代表数字 40
因此，数字总和 = 495 + 491 + 40 = 1026
 

提示：

树中节点的数目在范围 [1, 1000] 内
0 <= Node.val <= 9
树的深度不超过 10
*/
package main

import (
	. "awesomeProject1/Algorithm/main/structure/TreeNode"
	"strconv"
)


// 我的题解：dfs
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了44.91%的用户
*/

func sumNumbers(root *TreeNode) int {
	var res int
	var dfs func (*TreeNode, string)
	dfs = func(node *TreeNode, s string) {
		if node == nil{
			return
		}

		s += strconv.Itoa(node.Val)
		if node.Right == nil && node.Left == nil{
			num, _ := strconv.Atoi(s)
			res += num
		}

		dfs(node.Left, s)
		dfs(node.Right, s)
	}
	dfs(root, "")
	return res
}