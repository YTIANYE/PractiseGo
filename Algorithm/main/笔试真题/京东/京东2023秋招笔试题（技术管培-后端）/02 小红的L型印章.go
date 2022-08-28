/**
小红有L型的印章，即
10
11
其中1为涂色部分
小红可以将印章旋转用于涂色

给出一个图形，判断小红是否可以用印章图出这个图形
 */

package main

import (
	"fmt"
	"strconv"
)
/**
2
3 3
010
010
010
3 4
1100
1101
0011

1
3 4
1100
1101
0011
 */

// 50%
func main() {
	var t int
	fmt.Scan(&t)
	for ; t != 0; t-- {
		if yunxing(){
			fmt.Println("Yes")
		}else {
			fmt.Println("No")
		}
	}
}

func yunxing() bool {
	var n, m int
	fmt.Scan(&n, &m)
	//
	if n < 2 || m < 2 {
		return false
	}
	// 图
	tu := make([][]int, n)
	for i := range tu {
		tu[i] = make([]int, m)
	}
	for i := range tu {
		var s string
		fmt.Scan(&s)
		for j := 0; j < m; j++ {
			tu[i][j], _ = strconv.Atoi(string(s[j]))
		}
	}
	// 标记
	tag := make([][]int, n)
	for i := range tag {
		tag[i] = make([]int, m)
	}
	// 是否涂满
	isRight := func() bool {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if tu[i][j] != tag[i][j] {
					return false
				}
			}
		}
		return true
	}
	//a b
	//c d
	// 涂色
	tuse := func(a, b,c, d int) bool {
		if a == 1 && b == 1 && c == 1 && d == 1 {
			return true
		}
		if a == 0 && b == 1 && c == 1 && d == 1 {
			return true
		}
		if a == 1 && b == 0 && c == 1 && d == 1 {
			return true
		}
		if a == 1 && b == 1 && c == 0 && d == 1 {
			return true
		}
		if a == 1 && b == 1 && c == 1 && d == 0 {
			return true
		}
		return false
	}

	// 开始涂色
	for i:=0;i<n-1;i++ {
		for j:=0;j<m-1;j++ {
			a := tu[i][j]
			b := tu[i][j+1]
			c := tu[i+1][j]
			d := tu[i+1][j+1]
			if tuse(a,b,c,d) {
				if a == 1 {
					tag[i][j] = 1
				}
				if b == 1 {
					tag[i][j+1] = 1
				}
				if c == 1 {
					tag[i+1][j] = 1
				}
				if d == 1 {
					tag[i+1][j+1] = 1
				}
			}
		}
	}
	return isRight()
}
