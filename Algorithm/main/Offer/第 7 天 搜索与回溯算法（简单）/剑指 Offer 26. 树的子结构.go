/*
输入两棵二叉树A和B，判断B是不是A的子结构。
(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:
		   3
		  / \
	 	4 	 5
	  / \
	1 	 2
给定的树 B：
	   4
	 /
	1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：
输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：
输入：A = [3,4,5,1,2], B = [4,1]
输出：true
限制：
0 <= 节点个数 <= 10000
*/
package main

import (
	"awesomeProject1/Algorithm/main/structure/TreeNode"
	"fmt"
)

/*我的题解*/
// 执行用时：12 ms, 在所有 Go 提交中击败了99.89%的用户
// 内存消耗：7 MB, 在所有 Go 提交中击败了60.43%的用户

func isSubStructure(A *TreeNode.TreeNode, B *TreeNode.TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if A == nil || B == nil {
		return false
	}
	if isSubTree(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B) {
		return true
	}

	// if isSubTree(A, B) {
	// 	return true
	// }
	// if A.Left != nil {
	// 	if isSubStructure(A.Left, B) {
	// 		return true
	// 	}
	// }
	// if A.Right != nil {
	// 	if isSubStructure(A.Right, B) {
	// 		return true
	// 	}
	// }
	return false
}

func isSubTree(a *TreeNode.TreeNode, b *TreeNode.TreeNode) bool {
	if b == nil { // b遍历完了结束
		return true
	} else if a == nil {
		return false
	} else if a.Val != b.Val {
		return false
	} else {
		return isSubTree(a.Left, b.Left) && isSubTree(a.Right, b.Right)
	}
}

func main() {
	// nums := []int{3,4,5,1,2,-1,-1}
	// nums := []int{1,4,3}
	nums := []int{3, 4, 5, 1, 2}
	nums2 := []int{4, 1, -1}
	A := TreeNode.TreeCreate(nums, 0)
	B := TreeNode.TreeCreate(nums2, 0)
	TreeNode.TreePrint(A)
	fmt.Println(" ")
	TreeNode.TreePrint(B)
	fmt.Println(" ")
	fmt.Println(isSubStructure(A, B))

}
