/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

 

示例 1:



输入: [1,2,3,null,5,null,4]
输出: [1,3,4]
示例 2:

输入: [1,null,3]
输出: [1,3]
示例 3:

输入: []
输出: []
 

提示:

二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100 


*/

package main

import (
	. "awesomeProject1/Algorithm/main/structure/TreeNode"
)

// 我的题解：层次遍历

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了99.82%的用户
*/

func rightSideView(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}

	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) != 0{

		tqueue := queue
		queue = []*TreeNode{}
		res = append(res, tqueue[len(tqueue)-1].Val)
		for len(tqueue) != 0{
			node := tqueue[0]
			tqueue = tqueue[1:]
			if node.Left != nil{
				queue = append(queue, node.Left)
			}
			if node.Right != nil{
				queue = append(queue, node.Right)
			}
		}
	}

	return res

}
