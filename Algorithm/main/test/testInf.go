package main

import (
	"fmt"
	"math"
)

/*
测试无穷
*/
func testInf() {
	// 浮点型
	fmt.Println(math.Inf(1))
	fmt.Println(math.Inf(-1))
	// 无符号整型uint
	const UINT_MIN uint = 0
	const UINT_MAX = ^uint(0)
	fmt.Println(UINT_MIN, UINT_MAX)
	// 有符号整型int
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	fmt.Println(INT_MIN, INT_MAX)
}
