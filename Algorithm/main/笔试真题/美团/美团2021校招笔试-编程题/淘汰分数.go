/*
某比赛已经进入了淘汰赛阶段,已知共有n名选手参与了此阶段比赛，他们的得分分别是a_1,a_2….a_n,小美作为比赛的裁判希望设定一个分数线m，使得所有分数大于m的选手晋级，其他人淘汰。

但是为了保护粉丝脆弱的心脏，小美希望晋级和淘汰的人数均在[x,y]之间。

显然这个m有可能是不存在的，也有可能存在多个m，如果不存在，请你输出-1，如果存在多个，请你输出符合条件的最低的分数线。

数据范围：，
进阶：时间复杂度，空间复杂度

输入描述:
输入第一行仅包含三个正整数n,x,y，分别表示参赛的人数和晋级淘汰人数区间。(1<=n<=50000,1<=x,y<=n)

输入第二行包含n个整数，中间用空格隔开，表示从1号选手到n号选手的成绩。(1<=|a_i|<=1000)


输出描述:
输出仅包含一个整数，如果不存在这样的m，则输出-1，否则输出符合条件的最小的值。


输入例子1:
6 2 3
1 2 3 4 5 6

输出例子1:
3
*/

package main

import (
	"fmt"
)

// 我的题解：通过样例100%

func main() {

	var n, x, y int
	fmt.Scan(&n, &x, &y)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums[i] = num
	}

	// 排序
	var quickSort func(left, right int)
	quickSort = func(left, right int) {
		if left >= right {
			return
		}
		i, j := left, right
		for i < j {
			for i < j && nums[j] >= nums[left] { // 注意存在相等值时候的快排，其中一个要加等号
				j--
			}
			for i < j && nums[i] < nums[left] {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[left], nums[i] = nums[i], nums[left]
		quickSort(left, i-1)
		quickSort(i+1, right)
	}
	quickSort(0, n-1)
	// fmt.Println(nums)

	// sort.Ints(nums)
	// 计算:注意nums中可能存在相同的值
	res := -1
	for i := 0; i < n; i++ {
		right := i
		for right < n && nums[right] == nums[i] {
			right++
		}
		num1, num2 := right, n-right
		// 0 1 2 3 4
		if num1 >= x && num1 <= y && num2 >= x && num2 <= y {
			res = nums[i]
			break
		}
	}
	fmt.Println(res)

}
