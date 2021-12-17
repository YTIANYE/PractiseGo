/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。

示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]
示例 2：
输入：root = [1,2,3], targetSum = 5
输出：[]
示例 3：
输入：root = [1,2], targetSum = 0
输出：[]

提示：
树中节点总数在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000

*/
package main

import (
	. "awesomeProject1/github.com/YTIANYE/Algorithm/main/structure/TreeNode"
	"fmt"
)

// 我的题解：DFS
/*
执行用时：4 ms, 在所有 Go 提交中击败了89.80%的用户
内存消耗：4.5 MB, 在所有 Go 提交中击败了47.02%的用户
*/
func pathSum1(root *TreeNode, target int) [][]int {
	pathsum := [][]int{}
	if root == nil {
		return nil
	}
	pathFind(root, target, &pathsum)
	return pathsum
}

// DFS,依次从树中的每个结点出发
func pathFind(root *TreeNode, target int, pathsum *[][]int) {
	if root == nil {
		return
	}
	path := []int{}
	pathEach(root, target, path, pathsum)
	// 加上以下部分，就是树中的全部节点都可以作为起始节点
	// pathFind(root.Left, target, pathsum)
	// pathFind(root.Right, target, pathsum)
}

// 从当前节点出发的路径
func pathEach(root *TreeNode, target int, path []int, pathsum *[][]int) {
	if root == nil {
		return
	}

	sum := 0
	for _, val := range path { // for p := range path 中 p只是索引，不是value
		sum += val
	}

	// 形式一：不存在负数
	// if sum+root.Val == target && root.Left == nil && root.Right == nil { // 注意最后一个节点必须是叶子结点
	// 	path = append(path, root.Val)
	// 	p := make([]int, len(path)) // 切片间复制是引用复制，需要新建一个路径，避免后面对path的操作有影响
	// 	copy(p, path)
	// 	*pathsum = append(*pathsum, p)
	// 	return
	// } else if sum+root.Val < target {
	// 	path = append(path, root.Val)
	// 	pathEach(root.Left, target, path, pathsum)
	// 	pathEach(root.Right, target, path, pathsum)
	// }

	// 形式二：存在负数
	if sum+root.Val == target && root.Left == nil && root.Right == nil { // 注意最后一个节点必须是叶子结点
		path = append(path, root.Val)
		p := make([]int, len(path)) // 切片间复制是引用复制，需要新建一个路径，避免后面对path的操作有影响
		copy(p, path)
		*pathsum = append(*pathsum, p)
		return
	} else {
		path = append(path, root.Val)
		pathEach(root.Left, target, path, pathsum)
		pathEach(root.Right, target, path, pathsum)
	}
}

// 我的题解：DFS   不使用defer，（采用C、python语言的方式）回退指针

func pathSum2(root *TreeNode, target int) [][]int {
	path := []int{}
	ans := [][]int{}
	var dfs func(*TreeNode, int) // 递归调用自己，就需要声明，不能声明定义在一起
	dfs = func(root *TreeNode, left int) {
		if root == nil {
			return
		}
		left -= root.Val
		path = append(path, root.Val)

		if left == 0 && root.Left == nil && root.Right == nil {
			ans = append(ans, append([]int(nil), path...))
		}
		dfs(root.Left, left)
		dfs(root.Right, left)
		path = path[:len(path)-1]
	}
	dfs(root, target)
	return ans
}

// 我实现的官方题解：DFS
/*
执行用时：4 ms, 在所有 Go 提交中击败了89.61%的用户
内存消耗：4.4 MB, 在所有 Go 提交中击败了75.57%的用户
*/
func pathSum3(root *TreeNode, target int) [][]int {

	path := []int{}
	pathsum := [][]int{}
	var dfs func(*TreeNode, int) // 注意这种写法
	dfs = func(root *TreeNode, left int) {
		if root == nil {
			return
		}
		left -= root.Val
		path = append(path, root.Val)

		// defer函数会在return之后被调用，defer中的参数在声明时就已经确定了值了
		// defer输出的值，就是定义时的值。而不是defer真正执行时的变量值(很重要，搞不清楚的话就会产生于预期不一致的结果)
		defer func() { path = path[:len(path)-1] }()
		if left == 0 && root.Left == nil && root.Right == nil {
			pathsum = append(pathsum, append([]int{}, path...))
		}
		dfs(root.Left, left)
		dfs(root.Right, left)
		// defer 放在这个位置是不对的，因为defer定义时变量就固定了，放在这里path已经发生变化了,
		// 也就是存在之前执行了return，二而defer还没定义，defer的定义应该放在return之前
	}
	dfs(root, target)
	return pathsum
}

// 我的题解：BFS
/*
执行用时：232 ms, 在所有 Go 提交中击败了9.31%的用户
内存消耗：9.9 MB, 在所有 Go 提交中击败了5.12%的用户
*/
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	hashmap := make(map[*TreeNode]*TreeNode) // 保存儿子节点的父亲节点    index:son  value:father
	hashmap[root] = nil

	// 找路径
	findPath := func(node *TreeNode) []int {
		path := []int{}
		for node != nil {
			path = append([]int{node.Val}, path...) // 头插，保持顺序
			node = hashmap[node]
		}
		return path
	}

	pathsum := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		path := findPath(node)
		sum := 0
		for _, val := range path {
			sum += val
		}
		if sum == target && node.Left == nil && node.Right == nil {
			pathsum = append(pathsum, path)
		}
		if node.Left != nil {
			hashmap[node.Left] = node // 存入hash
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			hashmap[node.Right] = node // 存入hash
			queue = append(queue, node.Right)
		}
	}
	return pathsum
}

func main() {

	nums := []int{5,
		4, 8,
		11, -1, 13, 4,
		7, 2, -1, -1, -1, -1, 5, 1}
	target := 22

	// nums := []int{1, 2}
	// target := 1

	// nums := []int{-2, -1, -3}
	// target := -5

	root := TreeCreate(nums, 0)
	pathsum := pathSum(root, target)
	fmt.Println(pathsum)
}
