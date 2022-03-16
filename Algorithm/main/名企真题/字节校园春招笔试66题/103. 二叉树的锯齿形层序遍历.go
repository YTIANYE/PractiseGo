/*
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]

提示：

树中节点数目在范围 [0, 2000] 内
-100 <= Node.val <= 100

*/

package main

import . "awesomeProject1/Algorithm/main/structure/TreeNode"

// 我的题解：层次遍历
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.4 MB, 在所有 Go 提交中击败了64.40%的用户
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	flag := true // 左向右
	for len(queue) != 0 {
		var tempqueue []*TreeNode
		tempqueue = queue
		queue = []*TreeNode{}
		tempres := []int{}
		if flag {
			for i := range tempqueue{
				tempres = append(tempres, tempqueue[i].Val)
			}
		} else {
			for i := len(tempqueue) - 1; i >= 0; i-- {
				tempres = append(tempres, tempqueue[i].Val)
			}
		}
		res = append(res, tempres)
		flag = !flag

		for len(tempqueue) != 0 {
			node := tempqueue[0]
			tempqueue = tempqueue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
