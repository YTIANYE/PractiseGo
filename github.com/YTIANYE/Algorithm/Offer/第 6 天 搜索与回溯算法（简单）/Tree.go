package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*打印一棵树 先序遍历*/

func TreePrint(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	TreePrint(root.Left)
	TreePrint(root.Right)
}

/*由数组创建树*/

func TreeCreate(nums []int, i int) *TreeNode {
	if len(nums) == 0 || i >= len(nums) || nums[i] == -1 {
		return nil
	}
	node := &TreeNode{Val: nums[i]}
	node.Left = TreeCreate(nums, i*2+1)
	node.Right = TreeCreate(nums, i*2+2)
	return node
}

// func main() {
// 	nums := []int{0, 1, 2, -1, -1, 5, 6}
// 	root := TreeCreate(nums, 0)
// 	TreePrint(root)
// }
