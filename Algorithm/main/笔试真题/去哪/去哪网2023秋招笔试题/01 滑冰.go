/**
题目： 01背包问题

滑冰运动员有体力energy
他可以选择多组动作，每个动作有相应消耗的体力值和获得的分数
问他最多可以获得多少分？
 */

package main

import "sort"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 运动员可得到的最高分
 * @param energy int整型 运动员体力值
 * @param actions int整型二维数组 二维数组i为动作号 actions[i][0]为动作i+1消耗体力,actions[i][1]为动作i+1得分
 * @return int整型
 */
// 20分的12%
func maxScore( energy int ,  actions [][]int ) int {
	// write code here
	n := len(actions)
	nums :=make( []float64, n)
	for i, action:= range actions {
		nums[i] = float64(action[1]) / float64(action[0])
	}
	hash := make(map[float64]int)
	sort.Float64s(nums)
	res := 0
	for i:=n-1;i>=0 && energy > 0 ;i-- {
		index := hash[nums[i]]
		if energy >= actions[index][0] {
			res += actions[index][1]
			energy -= actions[index][0]
		}
	}
	return res
}