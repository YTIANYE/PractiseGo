/**
好的子数组的数量
时间限制： 3000MS
内存限制： 589824KB
题目描述：
现在一个数组被定义为好的，如果该数组中的最大值是最小值的k倍。

现在给你一个数组，你的任务是计算有多少个子数组是好的。



其中，子数组定义为原数组中一段连续的数组。

例如：[4, 3, 2, 7]有以下几个子数组：

[4], [4, 3], [4, 3, 2], [4, 3, 2, 7], [3], [3, 2], [3, 2, 7], [2], [2, 7], [7]



当k = 2时，答案为1，只有[4, 3, 2]是好数组，它的最大值是4，最小值是2，满足题意。



1<= n, k <= 1000

对于全体数据都保证数组中的数字在[1, 1e9]范围内



输入描述
第一行是一个正整数n, k，表示数组长度为n，好数组中最大值是最小值的k倍。

第二行是n个以空格分开的正整数。依次表示这个数组中的数字。

输出描述
一行一个非负整数，表示这个数组有几个子数组是好的。


样例输入
4 2
4 3 2 7
样例输出
1

提示
输入样例2

4 3

3 9 3 9

输出样例2

6

输入样例3

2 1

2 2

输出样例

3
 */
package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	//
	res := 0

	for i:=0;i<n;i++ {
		Max := nums[i]
		Min := nums[i]
		for j:=i;j<n;j++ {
			num := nums[j]
			Max = max(Max, num)
			Min = min(Min, num)
			if Max == Min * k {
				res++
			}
		}
	}
	fmt.Println(res)
}

func max(a, b int ) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}