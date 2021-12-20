package main

import (
	"awesomeProject1/Algorithm/main/structure/TreeNode"
	"fmt"
)

/*
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
*/

/*我的题解*/

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.8 MB, 在所有 Go 提交中击败了68.71%的用户

func levelOrder(root *TreeNode.TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode.TreeNode{root}
	level := true // 从左往右
	for len(queue) != 0 {
		var r []int
		l := len(queue) // 先取值
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			r = append(r, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		if !level {
			for i := 0; i < len(r)/2; i++ {
				r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
			}
		}
		result = append(result, r)
		level = !level
	}

	return result
}

func main() {
	nums := []int{0, 1, 2, 3, -1, 5, 6}
	root := TreeNode.TreeCreate(nums, 0)
	TreeNode.TreePrint(root)
	fmt.Println(" ")
	fmt.Println(levelOrder(root))
}
