/**
1...n个树节点，[没有说明是二叉树]
每个树节点有一种颜色，比如 'A'
求每个树节点的及其子树拥有的颜色的种类数量是多少？
*/
package main

import "fmt"

type TreeNode struct {
	Val    int
	colors map[string]bool // 该节点及子树拥有的颜色的种类
	Son    []*TreeNode     // 孩子结点
}

// 100%
func main() {
	var n int
	fmt.Scan(&n)//结点数目
	hashNode := make(map[int]*TreeNode)
	for i := 1; i <= n; i++ {
		hashNode[i] = &TreeNode{Val: i}
		hashNode[i].Son = []*TreeNode{}            // 注意初始化
		hashNode[i].colors = make(map[string]bool) // 注意初始化
	}
	// 建树
	// 第i个数是第i+1的结点的父节点
	for i := 1; i < n; i++ {
		var root int
		fmt.Scan(&root)
		son := i + 1
		hashNode[root].Son = append(hashNode[root].Son, hashNode[son])
	}
	var s string
	fmt.Scan(&s)// 树节点的颜色
	for i := 1; i <= n; i++ {
		hashNode[i].colors[string(s[i-1])] = true
	}
	//
	res := make([]int, n+1)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		for _, son := range root.Son {
			dfs(son)
			for color := range son.colors {// 获取son的结点颜色的种类
				root.colors[color] = true
			}
		}
		res[root.Val] = len(root.colors)
	}
	dfs(hashNode[1])
	//打印
	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Printf("%d", res[i])
		} else {
			fmt.Printf(" %d", res[i])
		}
	}
}
