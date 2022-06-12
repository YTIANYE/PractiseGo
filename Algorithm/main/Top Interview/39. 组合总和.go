/**
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。



示例 1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
示例 2：

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
示例 3：

输入: candidates = [2], target = 1
输出: []


提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都 互不相同
1 <= target <= 500
*/

package main

import "sort"

// 我的题解:深度遍历，暴力列举

/**
执行用时：4 ms, 在所有 Go 提交中击败了41.82%的用户
内存消耗：3.4 MB, 在所有 Go 提交中击败了17.17%的用户
 */

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}

	var dfs func([]int, []int, int)
	dfs = func(candidate , nums []int, sum int) {
		for i, num := range candidate {
			if num+sum < target {
				dfs(candidate[i:], append(nums, num), sum+num)
			} else if num+sum > target {
				return
			} else {
				nums = append(nums, num)
				temp := make([]int, len(nums))
				copy(temp, nums)// 注意直接temp := nums,会发生错误
				nums = nums[:len(nums)-1]
				res = append(res, temp)
				return
			}
		}
	}
	dfs(candidates, []int{}, 0)

	return res
}
