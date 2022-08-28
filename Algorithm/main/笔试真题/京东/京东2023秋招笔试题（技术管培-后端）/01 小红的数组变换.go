/**
给出四个数 a, b, i, j
每次操作是a*i+j
判断a最少经过几步操作可以大于b
 */

/**
3
1 4 2 1
4 1 2 1
1 5 1 2

2
0
3
 */
package main

import "fmt"

// 100%
func main() {
	var t int
	fmt.Scan(&t)
	for ; t != 0; t-- {
		var a, b, i, j int
		fmt.Scan(&a, &b, &i, &j)
		//
		if a > b {
			fmt.Println(0)
			continue
		}

		res := 0
		//
		if i == 1 {
			cha := b - a + 1
			if cha%j == 0 {
				res = cha / j
			} else {
				res = cha/j + 1
			}
			fmt.Println(res)
			continue
		}
		//

		for a <= b {
			a = a*i + j
			res++
		}
		fmt.Println(res)
	}
}

// 80%
//func main() {
//	var t int
//	fmt.Scan(&t)
//	for ;t != 0 ;t-- {
//		var a, b ,i,j int
//		fmt.Scan(&a, &b, &i, &j)
//		//
//		if a > b {
//			fmt.Println(0)
//			continue
//		}
//		//
//		res := 0
//		for a <=b {
//			a = a*i+j
//			res++
//		}
//		fmt.Println(res)
//	}
//}
