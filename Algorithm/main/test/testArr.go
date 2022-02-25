package main

import (
	"fmt"
)


/*测试数组复制*/
func testcopyarr() {
	// 1 数组复制是值复制
	var ages = [4]int{1, 2, 3, 5}
	var b = ages
	b[0]++
	fmt.Println("ages", ages) // [1 2 3 5]
	fmt.Println("b", b)       // [2 2 3 5]

	// 2 切片的复制是引用复制
	// 切片的本质就是指向数组的指针
	slice := []int{1, 2, 3}
	slice2 := slice
	slice2[0]++
	fmt.Println("slice", slice)   // [2 2 3]
	fmt.Println("slice2", slice2) // [2 2 3]

	// 3 切片copy	Go语言的内置函数 copy() 可以将一个数组切片复制到另一个数组切片中，

	// 如果加入的两个数组切片不一样大，就会按照其中较小的那个数组切片的元素个数进行复制。
	slice3 := []int{}
	copy(slice3, slice)
	fmt.Println("slice3", slice3) // []		slice3的容量较小

	// slice3 和 slice4 地址不同
	slice4 := make([]int, len(slice))
	copy(slice4, slice)
	fmt.Println("slice4", slice4) // [2 2 3]

	slice4[0]++
	fmt.Println("slice", slice)   // [2 2 3]
	fmt.Println("slice4", slice4) // [3 2 3]

	// 4 用append复制一个切片时
	slice5 := append([]int{}, slice...)
	fmt.Println("slice5", slice5)

	var slice6 []int
	slice6 = append(slice6, slice...)
	fmt.Println("slice6", slice6)
}

/*测试数组 头插法 */
func testHeadInsert() {
	nums := []int{1, 2, 3}
	res := []int{}
	for _, num := range nums {
		res = append([]int{num}, res...)
	}
	fmt.Println("原数组", nums)
	fmt.Println("反转后的数组", res)
}

/*测试切片中添加nil*/
func testslicenil() {
	nums := []int{1, 2, 3}
	numsp := []*int{}
	for _, v := range nums {
		numsp = append(numsp, &v)
	}
	numsp = append(numsp, nil)
	for _, v := range nums {
		numsp = append(numsp, &v)
	}
	// 地址可以添加 nil
	// int不可以
	fmt.Println(numsp) // [0xc00000a088 0xc00000a088 0xc00000a088 <nil> 0xc00000a0a0 0xc00000a0a0 0xc00000a0a0]
}



func main() {
	// testcopyarr()
	testHeadInsert()
}
