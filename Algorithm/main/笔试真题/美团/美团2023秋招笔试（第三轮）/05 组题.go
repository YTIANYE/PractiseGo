/**
A B C 是长度为n的数组
每个数组中各选一个数 a, b, c
要求
a < b < a*2
b < c < b*2
问有多少种选择方法？
 */

package main

import (
	"fmt"
	"sort"
)
// 81%
func main() {
	var n int
	fmt.Scan(&n)
	A := make([]int, n)
	B:= make([]int, n)
	C := make([]int, n)
	for i:=0;i<n;i++ {
		fmt.Scan(&A[i])
	}
	for i:=0;i<n;i++ {
		fmt.Scan(&B[i])
	}
	for i:=0;i<n;i++ {
		fmt.Scan(&C[i])
	}
	sort.Ints(A)
	sort.Ints(B)
	sort.Ints(C)
	//
	res := 0
	for i:=0;i<n;i++ {
		a := A[i]
		for j:=0;j<n;j++ {
			if B[j] <= a {
				continue
			}
			if B[j] > a * 2 {
				break
			}
			b := B[j]
			//
			for k:=0;k<n;k++ {
				if C[k] <= b {
					continue
				}
				if C[k] > b * 2 {
					break
				}
				res ++
				//c := C[k]
				//fmt.Println(res, a, b, c)
			}
		}
	}
	fmt.Println(res)
}
