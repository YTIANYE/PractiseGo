/*
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

 

示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 

提示：

1 <= numCourses <= 105
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对 互不相同

*/

package main

import "fmt"

// 我实现的精选题解：入度表（广度优先遍历）
/*
执行用时：4 ms, 在所有 Go 提交中击败了99.57%的用户
内存消耗：6.3 MB, 在所有 Go 提交中击败了37.98%的用户
*/
func canFinish1(numCourses int, prerequisites [][]int) bool {
	inDegrees := []int{}
	sons := make(map[int][]int) // key:父亲  values:儿子们
	for i := 0; i < numCourses; i++ {
		inDegrees = append(inDegrees, 0)
	}

	// 入度表、父子表
	for _, v := range prerequisites {
		cur, pre := v[0], v[1]
		inDegrees[cur]++
		sons[pre] = append(sons[pre], cur)
	}
	// 初始化queue
	queue := []int{}
	for i, v := range inDegrees {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	count := 0 // 出栈计数
	for len(queue) != 0 {
		num := queue[0]
		count++
		queue = queue[1:]
		for _, son := range sons[num] {
			inDegrees[son]--
			if inDegrees[son] == 0 {
				queue = append(queue, son)
			}
		}
	}
	return count == numCourses
}

// 我实现的精选题解：深度优先遍历
/*
执行用时：8 ms, 在所有 Go 提交中击败了93.78%的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了40.24%的用户
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	mySons := make(map[int][]int)
	flags := make([]int, numCourses) // 0 未被访问过  1 当前访问中  -1 被别人访问过

	for _, v := range prerequisites {
		pre, son := v[1], v[0]
		mySons[pre] = append(mySons[pre], son)
	}

	var dfs func(int) bool
	dfs = func(i int) bool {
		if flags[i] == 1 {
			return false
		} else if flags[i] == -1 {
			return true
		}
		// flags[i] == 0
		flags[i] = 1
		for _, son := range mySons[i] {
			if dfs(son) == false {
				return false
			}
		}
		flags[i] = -1
		return true
	}

	for i := 0; i < numCourses; i++ {
		if dfs(i) == false {
			return false
		}
	}
	return true
}

func main() {
	n := 3
	nums := [][]int{
		{1, 0},
		{1, 2},
		{0, 1},
	}

	fmt.Println(canFinish(n, nums))
}
