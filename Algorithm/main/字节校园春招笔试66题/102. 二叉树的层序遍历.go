/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

 

示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]
 

提示：

树中节点数目在范围 [0, 2000] 内
-1000 <= Node.val <= 1000

*/

package main

import . "awesomeProject1/Algorithm/main/structure/TreeNode"

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.6 MB, 在所有 Go 提交中击败了92.12%的用户
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	for len(queue)!= 0{

		tempres := []int{}
		tempqueue := queue
		queue = []*TreeNode{}

		for len(tempqueue)!= 0{
			node := tempqueue[0]
			tempqueue = tempqueue[1:]
			tempres = append(tempres, node.Val)
			if node.Left != nil{
				queue = append(queue, node.Left)
			}
			if node.Right != nil{
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tempres)
	}
	return res

}