package main

import "fmt"

/*测试字符的类型*/
func testchar() {
	s := "abc"  // int32
	char := 'a' // int32
	var cha byte
	cha = 'b' // uint8
	for _, ch := range s {
		fmt.Printf("%T\n", ch) // 单个字符类型 int32
	}
	fmt.Printf("%T\n", char)
	fmt.Printf("%T\n", cha)
	fmt.Printf("%T\n", byte(char))
}

func testCopy(){
	s := "abc"
	ss := make([]uint8, len(s)) // 创建副本
	copy(ss, s)
	fmt.Printf(string(ss))
}

func main(){
	testCopy()
}

