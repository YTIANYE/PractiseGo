package main

import "fmt"

func main() {
	const a = 1 // const a := 1 是错误的方式
	b := 10.0
	fmt.Println(a / b) // 不会报错，a是常量，即使不是小数

	// c := 1
	// fmt.Println(c / b) // 会报错，类型不一
}
