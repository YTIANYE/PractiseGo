/**
每个节点有一个颜色，可以是r g b
n个节点组成的一棵树，
输入n-1条边

问减去一条边后，剩下的每个部分，能不能都包含3种颜色，
如果可以，计数加一，问一共有多少种可能？
 */

package main

import "fmt"

/**
7
rgbrgbg
1 2
2 3
3 4
4 5
5 6
6 7
*/


// 31.24%
func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	tu := make([][]int, n)
	for i := range tu {
		tu[i] = make([]int, n)
	}
	bians := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		tu[a][b] = 1
		tu[b][a] = 1
		bians[i] = []int{a, b}
	}
	//
	var find func(int) bool
	find = func(num int) bool {
		tu1 := make([][]int, n)
		for i := range tu1 {
			tu1[i] = make([]int, n)
			copy(tu1[i], tu[i])
		}

		queue := []int{num}
		color := make(map[byte]bool)
		for len(queue) != 0 {
			i := queue[0]
			color[s[i]] = true
			if len(color) == 3 {
				return true
			}
			if len(queue) == 1 {
				queue = []int{}
			} else {
				queue = queue[1:]
			}
			//
			for j := 0; j < n; j++ {
				if tu1[i][j] == 1 {
					queue = append(queue, j)
					tu1[i][j] = 0
					tu1[j][i] = 0
				}
			}
		}
		return false
	}

	//
	var res int
	for _, bian := range bians {
		a, b := bian[0], bian[1]
		tu[a][b] = 0
		tu[b][a] = 0

		res1, res2 := find(a), find(b)
		if res1 && res2 {
			res++
		}
		tu[a][b] = 1
		tu[b][a] = 1
	}
	fmt.Println(res)
}

//
//type Node struct {
//	Val   byte // 颜色
//	Son   []*Node
//	Color map[byte]bool // 它自己及子树拥有的颜色
//}
//
//func main() {
//	var n int
//	fmt.Scan(&n)
//	var s string
//	fmt.Scan(&s)
//	hash := make(map[int]*Node)
//	for i := range s {
//		node := &Node{s[i], []*Node{}, make(map[byte]bool)}
//		node.Color[s[i]] = true
//		hash[i+1] = node
//	}
//	for i:=0;i<n-1;i++ {
//		var a, b int
//		fmt.Scan(&a, &b)
//		root := hash[a]
//		son := hash[b]
//		root.Son = append(root.Son, son)
//	}
//
//}
