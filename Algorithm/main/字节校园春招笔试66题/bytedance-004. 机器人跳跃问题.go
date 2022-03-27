/*
机器人正在玩一个古老的基于 DOS 的游戏。游戏中有 N+1 座建筑——从 0 到 N 编号，从左到右排列。编号为 0 的建筑高度为 0 个单位，编号为 i 的建筑的高度为 H(i) 个单位。
起初， 机器人在编号为 0 的建筑处。每一步，它跳到下一个（右边）建筑。假设机器人在第 k 个建筑，且它现在的能量值是 E, 下一步它将跳到第个 k+1 建筑。它将会得到或者失去正比于与 H(k+1) 与 E 之差的能量。如果 H(k+1) > E 那么机器人就失去 H(k+1) - E 的能量值，否则它将得到 E - H(k+1) 的能量值。
游戏目标是到达第个 N 建筑，在这个过程中，能量值不能为负数个单位。现在的问题是机器人以多少能量值开始游戏，才可以保证成功完成游戏？
格式：
输入：
- 第一行输入，表示一共有 N 组数据.
- 第二个是 N 个空格分隔的整数，H1, H2, H3, ..., Hn 代表建筑物的高度
输出：
- 输出一个单独的数表示完成游戏所需的最少单位的初始能量
示例 1：
输入：
     5
     3 4 3 2 4
输出：4
示例 2：
输入：
     3
     4 4 4
输出：4
示例 3：
输入：
     3
     1 6 4
输出：3
提示：
1 <= N <= 10^5
1 <= H(i) <= 10^5

*/

package main

import (
	"fmt"
	"math"
)



// 数学方法：解方程组
/*
执行用时：4 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func robort(nums []int) int {
	n := len(nums) - 1
	res := float64(nums[0])
	for i := 1; i <= n; i++ {
		res += math.Pow(2, float64(n-i)) * float64(nums[i])
	}
	res /= math.Pow(2, float64(n))
	return int(math.Ceil(res))

}

func main() {
	var N int
	fmt.Scan(&N)
	nums := []int{}
	for ; N > 0; N-- {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	nums = append([]int{0}, nums...)
	fmt.Println(robort(nums))

}










// 数学方法：错误
/*
	a	b		c		d
0	-a	-a-b	b2-c	c2-d
	-a	-2a-b	2b2-c	2c2-d

*/
func robort1(nums []int) int {
	count := 0
	sum := 0
	for i := range nums {
		count = count*2 + 1
		sum = sum*2 - nums[i]
	}
	sum = -sum
	res := sum / count
	if res*count == sum {
		return res
	}
	return res + 1

}
