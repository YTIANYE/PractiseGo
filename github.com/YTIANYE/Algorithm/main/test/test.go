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

/*测试数组复制*/
func testcopyarr() {
	// 数组复制是值复制
	var ages [4]int = [4]int{1, 2, 3, 5}
	var b = ages
	b[0]++
	fmt.Println(ages) // [1 2 3 5]
	fmt.Println(b)    // [2 2 3 5]

	// 切片的复制
	// 切片的本质就是指向数组的指针
	slice := []int{1, 2, 3}
	slice2 := slice
	slice2[0]++
	fmt.Println(slice)  // [2 2 3]
	fmt.Println(slice2) // [2 2 3]

	// Go语言的内置函数 copy() 可以将一个数组切片复制到另一个数组切片中，
	// 如果加入的两个数组切片不一样大，就会按照其中较小的那个数组切片的元素个数进行复制。

	slice3 := []int{}
	copy(slice3, slice)
	fmt.Println(slice3) // []

	// slice3 和 slice4 地址不同
	slice4 := make([]int, len(slice))
	copy(slice4, slice)
	fmt.Println(slice4) // [2 2 3]

	slice4[0]++
	fmt.Println(slice)//[2 2 3]
	fmt.Println(slice4)//[3 2 3]
}

/*测试切片中添加nil*/
func testslicenil(){
	nums := []int{1,2,3}
	numsp :=[]*int{}
	for _,v := range nums{
		numsp = append(numsp, &v)
	}
	numsp = append(numsp, nil)
	for _,v := range nums{
		numsp = append(numsp, &v)
	}
	// 地址可以添加 nil
	// int不可以
	fmt.Println(numsp)//[0xc00000a088 0xc00000a088 0xc00000a088 <nil> 0xc00000a0a0 0xc00000a0a0 0xc00000a0a0]
}

func main() {
	testslicenil()
}
