/*
给定一棵 n 个节点树。节点 1 为树的根节点，对于所有其他节点 i，它们的父节点编号为 floor(i/2) (i 除以 2 的整数部分)。在每个节点 i 上有 a[i] 个房间。此外树上所有边均是边长为 1 的无向边。
树上一共有 m 只松鼠，第 j 只松鼠的初始位置为 b[j]，它们需要通过树边各自找到一个独立的房间。请为所有松鼠规划一个移动方案，使得所有松鼠的总移动距离最短。

格式：


输入：
- 输入共有三行。
- 第一行包含两个正整数 n 和 m，表示树的节点数和松鼠的个数。
- 第二行包含 n 个自然数，其中第 i 个数表示节点 i 的房间数 a[i]。
- 第三行包含 m 个正整数，其中第 j 个数表示松鼠 j 的初始位置 b[j]。
输出：
- 输出一个数，表示 m 只松鼠各自找到独立房间的最短总移动距离。
示例：


输入：
     5 4
     0 0 4 1 1
     5 4 2 2
输出：4
解释：前两只松鼠不需要移动，后两只松鼠需要经节点 1 移动到节点 3
提示：

对于 30% 的数据，满足 n,m <=100。
对于 50% 的数据，满足 n,m <=1000。
对于所有数据，满足 n,m<=100000，0<=a[i]<=m, 1<=b[j]<=n。

*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var n, m int
	fmt.Scan(&n, &m)
	home := []int{0} // 每个数结点包含的房间数、0无实际意义
	locals := []int{}
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		home = append(home, num)
	}
	for i := 0; i < m; i++ {
		var local int
		fmt.Scan(&local)
		locals = append(locals, local)
	}
	//sort.Ints(locals)

	res := 0
	for _, v := range locals{
		minb, minstep :=  nearest(home, v, 0)
		res += minstep
		home[minb]--
	}

	fmt.Println(res)
}

func nearest(home []int, b int, step int) (int, int ) {
	n := len(home) - 1
	if b >n || b<1{
		return 0, math.MaxInt32
	}

	if home[b] != 0{
		return b,step
	}

	minStep := math.MaxInt32
	res := 0
	searched := 0
	for b>0{

		bdown, stepDown := findDown(home, b, step, searched)
		if stepDown< minStep{
			res = bdown
			minStep = stepDown
		}
		b /= 2
		searched = b
		step += 1
	}
	return res, minStep

}

// 从根往下找是否有合适的
func findDown(home []int, b , step, searched int) (int, int){
	n := len(home) - 1
	if b >n {
		return 0, math.MaxInt32
	}
	if home[b] != 0{
		return b,step
	}
	var bleft, left, right, bright int
	if b*2 != searched{
		bleft, left = findDown(home, b*2, step+1, b)
	}
	if b*2+1 != searched{
		bright, right = findDown(home, b*2+1, step+1, b)
	}

	if left<right{
		return bleft,left
	}else{
		return bright, right
	}

}

func min(a, b int )int{
	if a<b{
		return a
	}
	return b
}
