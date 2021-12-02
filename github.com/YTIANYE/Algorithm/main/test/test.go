package main

import (
	"fmt"
	"math"
)

/*
测试输入输出
*/
func testScan() {
	var num int
	fmt.Scan(&num)
	fmt.Println(num)
}

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

/*测试字符的类型*/
func testchar() {
	s := "abc"
	for _, ch := range s {
		fmt.Printf("%T\n", ch) // 单个字符类型 int32
	}
}

func main() {
	testchar()
}
