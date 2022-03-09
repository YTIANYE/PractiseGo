/*
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

 

示例 1：
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
示例 2：


输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
 
提示：

树中节点数目范围是 [1, 3 * 104]
-1000 <= Node.val <= 1000

*/

package main

import (
	. "awesomeProject1/Algorithm/main/structure/TreeNode"
	"fmt"
)

// 我的题解:只能从根到叶子出发
/*
执行用时：200 ms, 在所有 Go 提交中击败了7.41%的用户
内存消耗：7.3 MB, 在所有 Go 提交中击败了91.80%的用户
*/
func maxPathSum(root *TreeNode) int {
	return dfsEach(root, -1000)
}

func dfsEach(root *TreeNode, MAX int) int {
	if root == nil {
		return MAX
	}
	temp := Each(root)
	MAX = max(temp, MAX)
	MAX = dfsEach(root.Left, MAX)
	MAX = dfsEach(root.Right, MAX)
	return MAX
}

// 当前的根节点作为枢纽中心
func Each(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := root.Val
	left := dfs(root.Left)
	right := dfs(root.Right)
	if left>0{
		res += left
	}
	if right >0{
		res += right
	}
	return res
}

// 从一个结点出发，得到的最大值
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	return max(max(left, right) + root.Val, root.Val)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// nums := []int{1,2,3}
	// nums := []int{-10, 9, 20, -1, -1, 15, 7}
	nums := []int{2, -1}
	root := TreeCreate(nums, 0)
	TreePrint(root)
	fmt.Println()
	fmt.Println(maxPathSum(root))
}
