/*
我们称一个长度为n的序列为正则序列，当且仅当该序列是一个由1~n组成的排列，即该序列由n个正整数组成，取值在[1,n]范围，且不存在重复的数，同时正则序列不要求排序

有一天小团得到了一个长度为n的任意序列s，他需要在有限次操作内，将这个序列变成一个正则序列，每次操作他可以任选序列中的一个数字，并将该数字加一或者减一。

请问他最少用多少次操作可以把这个序列变成正则序列？

数据范围：，
进阶：时间复杂度，空间复杂度

输入描述:
输入第一行仅包含一个正整数n，表示任意序列的长度。(1<=n<=20000)

输入第二行包含n个整数，表示给出的序列，每个数的绝对值都小于10000。


输出描述:
输出仅包含一个整数，表示最少的操作数量。


输入例子1:
5
-1 2 3 10 100

输出例子1:
103
*/

package main

import (
	"fmt"
	"sort"
)

// 我的题解：通过样例 100%
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n+1)
	curr := []int{}
	for i := 1; i <= n; i++ {
		nums[i]++
	}
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		if 0 <= num && num <= n && nums[num] == 1 {
			nums[num] = 0
		} else {
			curr = append(curr, num)
		}
	}

	res := 0
	sort.Ints(curr)
	j := 0
	for i := 1; i <= n; i++ {
		if nums[i] == 1 {
			res += abs(i, curr[j])
			j++
		}
	}

	fmt.Println(res)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
