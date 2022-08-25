package main

import (
	"fmt"
	"io"
)

// 未通过
//func main() {
//	var w, y, n int // 5 5 0.2 3
//	var x float64
//	fmt.Scan(&w, &y, &x, &n)
//	//
//	li := int(float64(w) * x)
//	zai := w - li
//	sum := w * y
//	for i := 0; i < n; i++ {
//		sum = zai * (y+1) + li * 21
//		y = sum / w
//
//		//fmt.Println(y)
//	}
//	if sum%w != 0 {
//		y++
//	}
//	fmt.Println(y)
//}

// 应该是有一组数，不是一个数
func main() {
	var w, y, n int
	var x float64
	for {
		if _, ok := fmt.Scan(&w, &y, &x, &n); ok != io.EOF {
			//
			li := int(float64(w) * x)
			zai := w - li
			sum := w * y
			for i := 0; i < n; i++ {
				sum = zai*(y+1) + li*21
				y = sum / w

				//fmt.Println(y)
			}
			if sum%w != 0 {
				y++
			}
			fmt.Println(y)
		} else {
			break
		}
	}
}
