/**

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

 

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]
 

提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

*/

package main

import "fmt"

// 我的题解：暴力法

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.8 MB, 在所有 Go 提交中击败了9.14%的用户
*/

func permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)

	var dfs func([]int, []int)
	dfs = func(curr []int, last []int) {
		if len(curr) == n{
			res = append(res, curr)
			return
		}
		for i := range last{
			tempcurr := []int{}
			templast := []int{}
			tempcurr = append(append(tempcurr, curr...), last[i])
			templast = append(append(templast, last[:i]...), last[i+1:]...)
			dfs(tempcurr, templast)
		}
	}
	dfs([]int{}, nums)


	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
