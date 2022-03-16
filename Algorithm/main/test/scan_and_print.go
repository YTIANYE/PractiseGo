package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fun1()
	// fun2()
	// fun3()

}

// 获取 []string

func fun1() {
	input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	str := strings.Split(input.Text(), " ")
	// 	// temp, _ := strconv.Atoi(strings.Split(input.Text(), " ")[0])
	// 	fmt.Println(str)
	// }
	input.Scan()
	str := strings.Split(input.Text(), " ")
	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Println(str[1])
}


// a + b
func fun2() {
	var a, b int
	for { // 循环获取输入
		if _, err := fmt.Scan(&a, &b); err != io.EOF {
			fmt.Println(a + b)
		} else {
			break
		}
	}
}

func fun3() {
	var t, a, b int
	fmt.Scan(&t)
	for t>0{
		fmt.Scan(&a, &b)
		fmt.Println(a+b)
		t--
	}
}
