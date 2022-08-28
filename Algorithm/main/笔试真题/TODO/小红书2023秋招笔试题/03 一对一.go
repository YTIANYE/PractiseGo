/**
一对一
时间限制： 3000MS
内存限制： 589824KB
题目描述：
某公司正在组织新员工团建，其中一项活动是两两配对进行石头剪刀布。

因为新员工来自不同的学校和专业，他们许多人之间并不熟悉，但他们之间的朋友关系形成了一棵树。即若将朋友关系描述为一张无向图，则这张无向图中任意两点之间有且仅有一条路径。为了避免尴尬，组织者希望互为朋友的配对数量尽可能多。现在他希望你求出互为朋友的配对数量最多能有多少。



输入描述
第一行有一个偶数n(2<=n<=1000)，代表有n个新员工。

第二行有n-1个空格隔开的数，第i个数ai代表编号为i+1的新员工与编号为ai的员工互为朋友。

输入保证新员工之间的朋友关系形成了一棵树

输出描述
输出在所有可能的配对方案中，互为朋友的配对数量最多是多少。


样例输入
6
1 2 2 3 3
样例输出
2

提示
如样例中，一共有6个新员工，朋友关系有以下五个(1,2),(2,3),(2,4),(3,5),(3,6)。

可以证明无论如何匹配这6个人， 最多只能有两对是互为朋友的，因此输出2。
 */
package main

import (
	"fmt"
)

/**
1 2
2 3
2 4
3 5
3 6
*/
//func main() {
//	var n int
//	fmt.Scan(&n)
//	hash := make(map[int][]int, n)
//	for i := 0; i < n; i++ {
//		hash[i] = []int{}
//	}
//	for i := 1; i <= n; i++ {
//		var a int
//		fmt.Scan(&a)
//		b := i + 1
//		if a > b {
//			a, b = b, a
//		}
//		hash[a] = append(hash[a], b)
//	}
//	//
//	queue := []int{}
//	res := 0
//
//	var dfs func(num, maxNum int) int
//	dfs = func(num, maxNum int) int {
//		if len(hash[num]) == 0 {  // 没有朋友
//			return 0
//		}
//
//	}
//
//	for key := range hash {
//		maxNum := dfs(key, 0)
//		res = max(res, maxNum)
//	}
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 由一个点开始
func start(x int , guanxi [][]int) int {
	var gx [][]int
	copy(gx, guanxi)
	queue := qingchu(x, gx)
	for i := range queue {
		y := queue[i]
		newq := newqueue(i, queue)
		var newgx [][]int
		copy(newgx, gx)
		qingchu(y, newgx)
		for j := range newq{
			start(newq[j], newgx)
		}
	}


}

// queue中去掉i上的数
func newqueue(i int, queue []int) []int {
	newq := []int{}
	if i == 0 {
		newq = append(newq, queue[1:]...)
	}else if i == len(queue)-1 {
		newq = append(newq, queue[:len(queue)-1]...)
	}else {
		newq = append(queue[:i], queue[i+1:]...)
	}
	return newq
}

func qingchu(x int, gx [][]int) []int {
	m := len(gx[0]) // 列数
	res := []int{}
	for j := 0; j < m; j++ {
		if gx[x][j] == 1 {
			gx[x][j] = 0
			res = append(res, j)
		}
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	guanxi := make([][]int, n+1)
	for i := range guanxi {
		guanxi[i] = make([]int, n+1)
	}
	// 输入关系
	for i := 1; i <= n; i++ {
		var a int
		fmt.Scan(&a)
		b := i + 1
		if a > b {
			a, b = b, a
		}
		guanxi[a][b] = 1
	}
	// 由该关系开始
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if guanxi[i][j] == 1 {
				duishu := start(i, guanxi)
				res = max(res, duishu)
				break
			}
		}
	}
	fmt.Println(res)
}
