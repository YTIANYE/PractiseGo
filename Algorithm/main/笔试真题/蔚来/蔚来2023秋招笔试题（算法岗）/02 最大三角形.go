/**
给出n个点的坐标，选出三个构成一个三角形，问构成的三角形的边，最多包含多少个坐标点，
构不成三角形，返回0
 */


/**
4
2 2
5 5
2 5
3 4

9
 */

package main

import (
	"fmt"
	"math"
)

// 30%
func main() {
	var n int
	fmt.Scan(&n)

	xs := make([]int, n)
	ys := make([]int, n)
	for i:=0;i<n;i++ {
		fmt.Scan(&xs[i])
		fmt.Scan(&ys[i])
	}

	res := 0
	for i:=0;i<n;i++ {
		x1, y1 := xs[i], ys[i]
		for j:=i+1;j<n;j++ {
			x2, y2 := xs[j], ys[j]
			for k:=j+1;k<n;k++ {
				x3, y3 := xs[k], ys[k]
				if !isSanJiao(float64(x1), float64(y1), float64(x2), float64(y2), float64(x3), float64(y3)) {
					continue
				}
				num := fun2(x1,y1, x2,y2, x3, y3)
				res = max(res, num )
			}
		}
	}
	fmt.Println(res)
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isSanJiao(x1, y1, x2,y2, x3, y3 float64) bool {
	l1 := math.Pow(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2), 0.5)
	l2 := math.Pow(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2), 0.5)
	l3 := math.Pow(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2), 0.5)
	if l1 + l2 > l3 && l1+l3 > l2 && l2 + l3 > l1 {
		return true
	}
	return false
}

func fun2(x1, y1, x2,y2, x3, y3 int) int {


	res := 3
	res += dian(x1, y1, x2, y2)
	res += dian(x1, y1, x3,y3)
	res += dian(x2, y2, x3,y3)
	return res
}



func dian(x1, y1, x2,y2 int) int {  // 两点连线上的点，不算端点
	x := abs(x1-x2)
	y := abs(y1-y2)
	if x < y {
		x, y = y, x   // x大
	}

	if x== 0  {
		return y -1
	}
	if y == 0  {
		return x -1
	}
	if x % y == 0 {
		return y - 1
	}
	return 0
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

/**
  0 1 2 3 4 5
0
1
2     *     *
3         *
4
5           *
 */
