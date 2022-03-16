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

// 测试 copy
func testCopy(){
	s := "abc"
	ss := make([]uint8, len(s)) // 创建副本
	copy(ss, s)
	fmt.Printf(string(ss))
}

// 测试比较数组大小
func testComp(){
	a := "123"
	b := "1"
	c := "0"
	d := "234"

	comp := func(str1, str2 string ) {
		if str1 > str2{
			fmt.Println("str1 > str2")
		}else if str1 < str2 {
			fmt.Println("str1 < str2")
		}else{
			fmt.Println("str1 = str2")
		}
	}
	comp(a,b)
	comp(a,c)
	comp(a,d)
}

func main(){
	testComp()
}

