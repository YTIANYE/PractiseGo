package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func print(root *TreeNode) {
	fmt.Println("先序")
	print1(root)
	fmt.Println("中序")
	print2(root)
	fmt.Println("后序")
	print3(root)
}

func print1(root *TreeNode)  {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	print1(root.Left)
	print2(root.Right)
}

func print2(root *TreeNode)  {
	if root == nil {
		return
	}
	print1(root.Left)
	fmt.Println(root.Val)
	print2(root.Right)
}


func print3(root *TreeNode)  {
	if root == nil {
		return
	}
	print1(root.Left)
	print2(root.Right)
	fmt.Println(root.Val)
}
