package main

import "fmt"

type TestStruct struct{}

func NilOrNot(v interface{}) {
	fmt.Println(v)                      // <nil>   v 是一个指向nil的地址
	fmt.Println(v == nil)               // false
	fmt.Println(v.(*TestStruct) == nil) // true 判断空接口中的值，需要使用类型断言
}

func main() {
	var s *TestStruct
	fmt.Println(s)        // <nil>
	fmt.Println(s == nil) // true
	NilOrNot(s)           // false
}
