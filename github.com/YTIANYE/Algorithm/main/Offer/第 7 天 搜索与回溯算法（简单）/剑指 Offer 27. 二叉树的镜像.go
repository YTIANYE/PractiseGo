/*
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

	 	 	4
	 	/ 	 \
	 2 	 	 7
	/ \ 	 / \
1 	 3 		6 	 9
镜像输出：

			4
	 	/ 	 \
	 7 	 	 2
	/ \ 	 / \
9 	 6 		3 	1


示例 1：

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

限制：
0 <= 节点个数 <= 1000

*/
package main

import . "awesomeProject1/github.com/YTIANYE/Algorithm/main/structure/TreeNode"

/*我的题解*/

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了13.48%的用户

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	node := &TreeNode{Val: root.Val}

	node.Right = mirrorTree(root.Left)

	node.Left = mirrorTree(root.Right)

	return node
}
