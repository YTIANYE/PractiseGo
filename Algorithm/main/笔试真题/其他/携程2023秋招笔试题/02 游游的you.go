/**
a, b, c 个  'y' 'o' 'u'
组合成字符串后，每个“you”记 2 分，每个“oo”记 1 分，问最多可以得到少分
*/

package main

import "fmt"

//func main() {
//	var q int
//	fmt.Scan(&q)
//	var a, b, c , num , res int64
//	for i := 0; i < q; i++ {
//
//		fmt.Scan(&a, &b, &c)
//
//		num = min(min(a, b), c)
//		res = num<<1
//		b = b - num - 1
//		if b > 0 {
//			res += b
//		}
//		fmt.Println(res)
//	}
//}
//
//func min(a, b int64) int64 {
//	if a < b {
//		return a
//	}
//	return b
//}


// 80%
func main() {
	var q int
	fmt.Scan(&q)
	for ; q != 0; q-- {
		var a, b, c int
		fmt.Scan(&a, &b, &c)

		num := min(min(a, b), c)
		res := num*2
		b = b - num - 1
		if b > 0 {
			res += b
		}
		fmt.Println(res)

	}
}
