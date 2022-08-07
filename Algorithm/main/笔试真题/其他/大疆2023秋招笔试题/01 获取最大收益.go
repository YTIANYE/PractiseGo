/**
DJI-获取最大收益-后端开发工程师（研发部）A卷
时间限制： 5000MS
内存限制： 1048576KB
题目描述：
DJI 现有一款新型无人机，假设这架无人机做匀速直线飞行，飞行路途上有均匀存放价值不等的物品，飞机有个小爪子，可抓取途中物品, 当抓取到后，飞机抓取功能进入冷却阶段，冷却时间恰好飞过两个物品（即飞机冷却时间内无法抓取物品）

给定一个代表飞行途中存放物品价值的非负整数数组，请计算出飞机沿途能抓取的物品的最大价值。

示例 1：

输入：[1,4,3,2]

输出：6

解释：抓取 2 号物品 (金额 = 4) ，然后抓取 4 号物品 (金额 = 2)。

  抓取到的最高金额 = 4 + 2 = 6 。

示例 2：

输入：[2,7,1,3,5]

输出：12

解释：抓取 2 号物品 (金额 = 7), 抓取 5 号物品 (金额 = 5)。

  抓取到的最高金额 = 7 + 5 = 12 。



输入描述
输入为路线上的物品价值, 输入为一个数组，使用空格隔开

输出描述
最终能获取的最大价值


样例输入
2 7 1 3 5
样例输出
12

提示
1 <= nums.length <= 100

0 <= nums[i] <= 400
 */

package main

import (
	"fmt"
	"io"
)

func main() {
	nums := []int{}
	for {
		var a int
		if _, ok := fmt.Scan(&a); ok != io.EOF{
			nums = append(nums, a)
		}else {
			break
		}
	}
	//
	n := len(nums)
	//if n <= 3 {
	//	a := nums[0]
	//	for i:=1;i<n;i++ {
	//		a = max(a, nums[i])
	//	}
	//	fmt.Println(a)
	//}else {
	//	//
	//	dp := make([]int, n)
	//	dp[0] = nums[0]
	//	dp[1] = max(nums[0], nums[1])
	//	dp[2] = max(dp[1], nums[2])
	//	for i:=3;i<n;i++ {
	//		dp[i] = max(max(dp[i-1], dp[i-2]), dp[i-3] + nums[i])
	//	}
	//	fmt.Println(dp[n-1])
	//}
	if n <= 2 {
		a := nums[0]
		for i:=1;i<n;i++ {
			a = max(a, nums[i])
		}
		fmt.Println(a)
	}else {
		//
		dp := make([]int, n)
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])
		for i:=2;i<n;i++ {
			dp[i] = max(dp[i-1], dp[i-2] + nums[i])
		}
		fmt.Println(dp[n-1])
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
