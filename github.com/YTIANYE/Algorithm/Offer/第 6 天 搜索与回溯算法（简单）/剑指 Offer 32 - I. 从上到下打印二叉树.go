package main

import "fmt"

/*
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：
[3,9,20,15,7]

提示：
节点总数 <= 1000
*/

/*层序遍历*/

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.6 MB, 在所有 Go 提交中击败了76.59%的用户

func levelOrder1(root *TreeNode) []int {
	result := []int{}
	queue := []*TreeNode{}
	if root == nil {
		return result
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		result = append(result, node.Val)
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

func main() {
	nums := []int{0, 1, 2, 3, -1, 5, 6}
	root := TreeCreate(nums, 0)
	TreePrint(root)
	fmt.Println(" ")
	fmt.Println(levelOrder1(root))
}
