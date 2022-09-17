/**
给出两个长度为n的数组，
若两个数组对应位置的两个子数组中，
对应位置上，后面的数减去前面的数相等，
则称这两个子数组为相似子数组，
求相似子数组的最大的长度是多少？
 */
package main

import "fmt"

/**
5
1 2 3 1 3
-1 0 3 -1 1

2



*/

// 20%
//func main() {
//	var n int
//	fmt.Scan(&n)
//	A := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&A[i])
//	}
//	B := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&B[i])
//	}
//
//	res := 0
//	l, r := 0, 0 // 左闭右开
//	for ; r < n-1; r++ {
//		if A[r+1]-A[r] != B[r+1]-B[r] {
//			//res = max1(res, r-l)
//			if r - l > res {
//				res = r-l
//			}
//			l = r + 1
//		}
//	}
//	// res = max1(res, r-l)
//	if r - l > res {
//		res = r-l
//	}
//	fmt.Println(res + 1)
//}
//
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// 20%
func main() {
	var n int
	fmt.Scan(&n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	B := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&B[i])
	}
	//
	a := make([]int, n-1)
	b := make([]int, n-1)
	for i:= range a {
		a[i] = A[i+1] - A[i]
	}
	for i:= range b {
		b[i] = B[i+1] - B[i]
	}
	//

	//
	//a := []int{0, 0, 1, 1, 0, 0, 0, 1, 1, 1}
	//b := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	//
	//
	res := 0
	l, r := 0, 0
	for r < n-1 {
		if a[r] == b[r] {
			r++
		} else {
			res = max1(res, r-l)
			r++
			l = r
		}
	}
	res = max1(res, r-l)
	fmt.Println(res + 1)

}
