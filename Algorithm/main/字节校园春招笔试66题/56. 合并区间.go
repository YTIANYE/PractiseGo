/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

 

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 

提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104

*/

package main

import "sort"

// 我的题解：排序+合并

/*
执行用时：120 ms, 在所有 Go 提交中击败了5.17%的用户
内存消耗：6.7 MB, 在所有 Go 提交中击败了79.24%的用户
*/

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 || n == 1 {
		return intervals
	}

	// 切片 按照一定的规则排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 0
	j := i + 1
	for j < len(intervals) {
		if intervals[j][0] <= intervals[i][1] {
			if intervals[j][1] > intervals[i][1] { // 注意右侧是不是第二个范围的数大
				intervals[i][1] = intervals[j][1]
			}
			intervals = append(intervals[:j], intervals[j+1:]...)
		} else {
			i++
			j++
		}
	}

	return intervals
}
