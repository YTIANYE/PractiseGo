/**
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。



示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同

*/

package main

// 我的题解：动态规划
/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了23.13%的用户
*/

func subsets0(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})

	for _, num := range nums {
		l := len(res)
		for i := 0; i < l; i++ {
			temp := append([]int{}, res[i]...)
			temp = append(temp, num)
			res = append(res, temp)
		}
	}
	return res
}

// 官方题解：二进制
func subsets1(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

// 我实现的官方题解：递归
/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了11.89%的用户
*/
func subsets2(nums []int) [][]int {
	res := [][]int{}

	var dfs func([]int, int)
	dfs = func(curr []int, index int) {
		if index == len(nums) {
			res = append(res, curr)
			return
		}
		dfs(append([]int{}, curr...), index+1) // curr是局部的
		dfs(append(append([]int{}, curr...), nums[index]), index+1)
	}
	dfs([]int{}, 0)
	return res
}

// 官方题解
func subsets(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1] // set是全局的。进来退出去
		dfs(cur + 1)
	}
	dfs(0)
	return
}
