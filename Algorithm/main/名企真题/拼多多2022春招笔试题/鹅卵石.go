/*
统计每一种颜色的鹅卵石出现的位置是不是按照等差数列排列
如果是，打印出每种颜色和对应的差，
打印时颜色由小到大
颜色只出现一次的也是等差数列，差为0

输入：
鹅卵石颜色的排列
输出：
符合要求的颜色的数目
每个颜色 和 差

例子：
输入：
1 2 1
输出：
2
1 2
2 0

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 记录颜色以及对应颜色出现的位置
	var n int
	fmt.Scan(&n)
	yanseMap := make(map[int][]int)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		_, ok := yanseMap[num]
		if ok {
			yanseMap[num] = append(yanseMap[num], i)
		} else {
			yanseMap[num] = []int{i}
		}
	}

	// 判断每种颜色的位置是不是等差数列
	res := [][]int{}
	for k, val := range yanseMap {
		ok, cha := isDencha(val)
		if ok {
			res = append(res, []int{k, cha})
		}
	}

	// 颜色由小到大排列
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})

	// 结果
	fmt.Println(len(res))
	for _, val := range res {
		fmt.Println(val[0], val[1])
	}

}

func isDencha(nums []int) (bool, int) { // 如果是，返回  是 差, 如果不是返回 不是， 0
	if len(nums) == 1 {
		return true, 0
	}
	temp := nums[1] - nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != temp {
			return false, 0
		}
	}
	return true, temp
}
