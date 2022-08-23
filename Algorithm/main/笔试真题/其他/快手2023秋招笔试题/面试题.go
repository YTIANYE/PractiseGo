package main

import "fmt"

// 判断输出

func FunSlice(s []int, t int) {
	s[0]++
	fmt.Println("print3: ", s, &s[0]) // [2 2 3] 0xc000016168
	s = append(s, t)                  // 因为长度不够，所以开辟了一段新的存储空间
	fmt.Println("print4: ", s, &s[0]) // [2 2 3 4] 0xc00000c390
	s[0]++
	fmt.Println("print5: ", s, &s[0]) // [3 2 3 4] 0xc00000c390
}

func main() {
	x := []int{1, 2, 3}
	fmt.Println("print1: ", x, &x[0]) // [1 2 3] 0xc000016168

	FunSlice(x, 4)
	fmt.Println("print2: ", x, &x[0]) // [2 2 3] 0xc000016168
}
