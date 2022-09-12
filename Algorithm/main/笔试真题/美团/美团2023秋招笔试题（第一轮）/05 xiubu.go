/**
修补
时间限制： 3000MS
内存限制： 589824KB
题目描述：
小团的玩具火箭有点磨损了，上面有很多地方翘起来了，小团想要用强力胶进行修补，但在强力胶凝结之前，需要找点东西压住。
幸好小团有很多这样的东西。小团有m种配重材料，第i种材料重ai单位重量（因为小团有太多了，可以认为每种都有任意多个）。
火箭上有n个地方翘起来了，需要至少bi单位重量的东西来压住，而且只能用一个配重材料来压，
(多了的话不好压，多个配重材料容易散开，所以小团不想用多个来折腾)。
小团想一次就把所有翘起来的地方全都修补好，请问他需要使用的配重材料重量之和最少是多少？

输入描述
第一行两个正整数n和m，分别代表需要修补的地方个数以及材料种类数。
接下来一行n个数b1,b2,...,bn，含义如题。
接下来一行m个数 a1,a2,...,am，含义如题。
对于40%的数据，n,m≤100
对于另外30%的数据，n,m≤2000
对于所有数据，1≤n,m≤50000，1≤ai,bi≤104

输出描述
输出小团最少需要使用的配重材料重量之和。如果没有任何办法满足，输出-1
样例输入
1 1
5
4
样例输出
-1
提示
样例1解释
需要5单位重量，只有4单位重量的材料，压不住，输出-1。

输入样例2
3 3
4 1 3
4 2 1
输出样例2
9
样例解释2
第一个地方需要重量为4的，第二个地方可以用重量为1的，第三个地方只能选择重量为4的才能压住。所以总重量需求为9。可以证明没有更优方案。
*/
package main

import (
	"fmt"
	"sort"
)
// 未验证
func main() {
	otherAnswer()
}

func myAnswer() {
	var n, m int
	fmt.Scan(&n, &m)
	b := make([]int, n)
	a := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i]) // 至少需要b
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i]) // 已有材料a
	}

	// 找到有序数组中，>=t的第一个数
	var biaFind func(t int, l, r int) int
	biaFind = func(t, l, r int) int {
		if l >= r {
			return a[l]
		}
		mid := (l + r) / 2
		if a[mid] < t {
			return biaFind(t, mid+1, r)
		} else {
			return biaFind(t, l, mid) // 查找第一个大于等于的方法
		}
	}

	//
	sort.Ints(a)
	var res int
	tag := true
	for i := 0; i < n; i++ {
		if b[i] > a[m-1] {
			tag = false
			break
		}
		res += biaFind(b[i], 0, m-1)
	}
	if tag {
		fmt.Println(res)
	} else {
		fmt.Println(-1)
	}
}

func otherAnswer() {
	var n, m int
	fmt.Scan(&n, &m)
	b := make([]int, n)
	a := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i]) // 至少需要b
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i]) // 已有材料a
	}
	//
	sort.Ints(a)
	sort.Ints(b)
	temp := 0
	res := 0
	tag := true
	for i := 0; i < n; i++ {
		j := temp //m只需要从上一次的开始遍历即可，从头遍历的话会超时
		for ; j < m; j++ {
			if b[i] <= a[j] {
				res += a[j]
				temp = j
				continue
			}
		}
		if j == m {
			tag = false
			break
		}
	}
	// 结果
	if tag {
		fmt.Println(res)
	} else {
		fmt.Println(-1)
	}
}
