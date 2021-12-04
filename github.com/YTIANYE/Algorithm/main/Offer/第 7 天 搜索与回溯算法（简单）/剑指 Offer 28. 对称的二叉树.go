/*
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树[1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,-1,3,-1,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3



示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：
输入：root = [1,2,2,-1,3,-1,3]
输出：false

限制：
0 <= 节点个数 <= 1000

*/
package main

import (
	. "awesomeProject1/github.com/YTIANYE/Algorithm/main/structure/TreeNode"
	"fmt"
)

/*我的题解1：一次层次遍历*/
// 执行用时：4 ms, 在所有 Go 提交中击败了20.85%的用户
// 内存消耗：3.1 MB, 在所有 Go 提交中击败了5.42%的用户

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) != 0 {
		var nums []int
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil { // 合理处理node为nil的情况
				nums = append(nums, -1)
			} else {
				nums = append(nums, node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		// 树根
		if len(nums) == 1 {
			continue
		}
		// nums不对称
		for i := 0; i < len(nums)/2; i++ {
			if nums[i] != nums[len(nums)-1-i] {
				return false
			}
		}
		// nums中全是-1结束
		tag := true
		for _, val := range nums {
			if val != -1 {
				tag = false
			}
		}
		if tag {
			break
		}

	}
	return true
}

/*我的题解2：两次先序遍历*/
// 执行用时：4 ms, 在所有 Go 提交中击败了20.85%的用户
// 内存消耗：3.2 MB, 在所有 Go 提交中击败了5.42%的用户

func order(root *TreeNode, nums *[]int, tag bool) {
	if root == nil {
		*nums = append(*nums, -1)
		return
	}
	*nums = append(*nums, root.Val)
	if tag { // 先左后右
		order(root.Left, nums, tag)
		order(root.Right, nums, tag)
	} else { // 先右后做
		order(root.Right, nums, tag)
		order(root.Left, nums, tag)
	}
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	if root.Right != nil && root.Left != nil {
		var nums1 []int
		var nums2 []int
		order(root.Left, &nums1, true)
		order(root.Right, &nums2, false)
		for i, val := range nums1 {
			if nums2[i] != val {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

/*我实现的精选题解：递归*/

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.9 MB, 在所有 Go 提交中击败了63.59%的用户
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return fun(root.Left, root.Right)
}

func fun(L *TreeNode, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	} else if L != nil && R != nil && L.Val == R.Val {
		return fun(L.Right, R.Left) && fun(L.Left, R.Right)
	} else {
		return false
	}
}

func main() {
	// nums := []int{6, 82, 82, -1, 53, 53, -1, -58, -1, -1, -58, -1, -85, -85, -1, -9, -1, -1, -9, -1, 48, 48, -1, 33, -1, -1, 33, 81, -1, -1, 81, 5, -1, -1, 5, 61, -1, -1, 61, -1, 9, 9, -1, 91, -1, -1, 91, 72, 7, 7, 72, 89, -1, 94, -1, -1, 94, -1, 89, -27, -1, -30, 36, 36, -30, -1, -27, 50, 36, -1, -80, 34, -1, -1, 34, -80, -1, 36, 50, 18, -1, -1, 91, 77, -1, -1, 95, 95, -1, -1, 77, 91, -1, -1, 18, -19, 65, -1, 94, -1, -53, -1, -29, -29, -1, -53, -1, 94, -1, 65, -19, -62, -15, -35, -1, -1, -19, 43, -1, -21, -1, -1, -21, -1, 43, -19, -1, -1, -35, -15, -62, 86, -1, -1, -70, -1, 19, -1, 55, -79, -1, -1, -96, -96, -1, -1, -79, 55, -1, 19, -1, -70, -1, -1, 86, 49, -1, 25, -1, -19, -1, -1, 8, 30, -1, 82, -47, -47, 82, -1, 30, 8, -1, -1, -19, -1, 25, -1, 49}
	// nums := []int{1, 2, 2, -1, 3, -1, 3}

	nums := []int{1, 2, 2, -1, 2, 2, -1, -1, -1, 2, -1, -1, 2, -1, -1}
	// nums := []int{1, 2, 2}
	root := TreeCreate(nums, 0)
	fmt.Println(isSymmetric(root))
}
