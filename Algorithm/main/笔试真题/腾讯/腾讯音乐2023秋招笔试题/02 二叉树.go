/**
给出一个先序遍历和一个中序遍历的数组
返回可以排列出的二叉树
 */

/**
[1,1,2],[1,2,1]

[{1,1,#,#,2},{1,#,1,2}]

 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param preOrder int整型一维数组
 * @param inOrder int整型一维数组
 * @return TreeNode类一维数组
 */
// 3.33%
func getBinaryTrees(preOrder []int, inOrder []int) []*TreeNode {
	// write code here
	if len(preOrder) == 0 {
		node := &TreeNode{Val: -1}
		return []*TreeNode{ node}
	}

	val := preOrder[0]
	res := []*TreeNode{}
	for i := range inOrder {
		if inOrder[i] == val {
			lefts := getBinaryTrees(preOrder[1:i+1], inOrder[:i])
			rights := getBinaryTrees(preOrder[i+1:], inOrder[i+1:])
			//if len(lefts) == 0 {
			//	node := &TreeNode{Val: val}
			//	res = append(res, node)
			//} else {
			//	for j := range lefts {
			//		node := &TreeNode{Val: val}
			//		if len(lefts) != 0 && lefts[j].Val != -1 {
			//			node.Left = lefts[j]
			//		}
			//		if len(rights) != 0 && rights[j].Val != -1 {
			//			node.Right = rights[j]
			//		}
			//		res = append(res, node)
			//	}
			//}
			if len(lefts) == 0 && len(rights) == 0 {
					node := &TreeNode{Val: val}
					res = append(res, node)
			}else if len(lefts) == 0 {
				for k := range rights {
					node := &TreeNode{Val: val}
					if rights[k].Val != -1 {
						node.Right = rights[k]
					}
					res = append(res, node)
				}
			}else if len(rights) == 0 {
				for j:= range lefts {
					node := &TreeNode{Val: val}
					if lefts[j].Val != -1 {
						node.Left = lefts[j]
					}
					res = append(res, node)
				}
			}else {
				for j := range lefts {
					for k := range rights {
						node := &TreeNode{Val: val}
						if lefts[j].Val != -1 {
							node.Left = lefts[j]
						}
						if rights[k].Val != -1 {
							node.Right = rights[k]
						}
						res = append(res, node)
					}
				}
				// res应该去重
			}
		}
	}
	return res
}

func main() {
	preOrder, inOrder := []int{1, 1, 2}, []int{1, 2, 1}
	fmt.Println(getBinaryTrees(preOrder, inOrder))
}
