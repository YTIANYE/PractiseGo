/*
河道不同的位置高度也不相同，从河道起点到终点有n个位置，这些位置编号为1~n。每一个位
置的高度可以表示为hi(l、n是河道的两端，因此1左边、n右边的高度可以视为无穷大）。
水会从高处向低处流动，但原来的位置仍然有水。具体地来说，如果当前一个位置1是有水的，有某—个相邻j高度严格小于i（hj＜ hi），那公也会成为有水的，并且仍然是有水的。对于
相邻的格子也是如此。
现在小七想知道，通过一次注入水源最多可以使得多少个位置变成有水的
样例输入

5
5 1 2 1 5
样例输出

3
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var nums []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	fmt.Println(shui(nums))
}

func shui(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	res := 0
	i := 1
	pre := 0
	// for i < n {
	// 	if nums[i] < nums[i-1] {
	// 		i++
	// 	}else{
	// 		break
	// 	}
	// }
	// res = i - pre

	pre = i - 1
	for i < n {
		for i < n && nums[i] == nums[i-1] {
			i++
			pre = i - 1
			res = max(res, i-pre)
			continue
		}
		for i < n && nums[i] > nums[i-1] {
			i++
		}
		if i<n && nums[i] == nums[i-1]{
			i++
			pre = i-1
			res = max(res, i-pre)
			continue
		}
		for i< n && nums[i] < nums[i-1]{
			i++
		}
		res = max(res, i-pre)
		pre = i-1
	}
	return res

}

func max(a, b int) int{
	if a>b{
		return a
	}
	return b
}