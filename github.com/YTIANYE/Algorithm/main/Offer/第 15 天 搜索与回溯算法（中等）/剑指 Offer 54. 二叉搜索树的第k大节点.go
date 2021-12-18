package main

import (
	. "awesomeProject1/github.com/YTIANYE/Algorithm/main/structure/TreeNode"
	"fmt"
)

// 我的题解:维护一个k大小的降序数组
/*
执行用时：12 ms, 在所有 Go 提交中击败了41.64%的用户
内存消耗：7.1 MB, 在所有 Go 提交中击败了5.42%的用户
*/
func kthLargest1(root *TreeNode, k int) int {
	numsMax := []int{} // key,大小次序  val 值	val降序
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if len(numsMax) == 0 {
			numsMax = append(numsMax, root.Val)
		} else {
			for i := len(numsMax) - 1; i >= 0; i-- {
				if numsMax[i] >= root.Val {
					numsMax = append(numsMax[:i+1], append([]int{root.Val}, numsMax[i+1:]...)...) // 插入这个数
					break                                                                         // 重要语句
				} else if i == 0 { // numsMax中第一个数就小于root.Val
					numsMax = append([]int{root.Val}, numsMax...)
				}
			}
			if len(numsMax) == k+1 {
				numsMax = numsMax[:k]
			}
		}
		// fmt.Println(numsMax)
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}

	}
	dfs(root)
	return numsMax[k-1]
}

// 我的题解：逆向中序遍历
/*执行用时：12 ms, 在所有 Go 提交中击败了41.64%的用户
内存消耗：6.4 MB, 在所有 Go 提交中击败了69.80%的用户*/
func kthLargest(root *TreeNode, k int) int {
	num := 0
	dfs(root, &k, &num)
	return num

}

func dfs(node *TreeNode, k *int, num *int) {
	if node.Right != nil {
		dfs(node.Right, k, num)
	}
	*k--
	if *k == 0 {
		*num = node.Val
		return // 后续遍历没有意义
	}
	if node.Left != nil {
		dfs(node.Left, k, num)
	}
}

func main() {
	nums := []int{5, 3, 6, 2, 4, -1, -1, 1}
	k := 3
	root := TreeCreate(nums, 0)
	// TreePrint(root)
	fmt.Println(kthLargest(root, k))
}
