/*
题目：名字在6-12个字符之间，只含有大小写字母，
输入 n 表示n 行数据
接下来每一行一个名字
判断每个名字的情况

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	names := make(map[string]int)

	for ; n > 0; n-- {
		// 注意读取用空格
		// var name string
		// fmt.Scan(&name)
		// 应该读取一整行
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		name := input.Text()
		if len(name) > 12 || len(name) < 6 {
			fmt.Println("illegal length")
		} else if !isright(name) {
			fmt.Println("illegal charactor")
		} else {
			_, ok := names[name]
			if ok {
				fmt.Println("acount existed")
			} else {
				names[name]++
				fmt.Println("registration complete")
			}
		}
	}
	fmt.Println(names)
}

func isright(name string) bool {
	for i := range name {
		// char := string(name[i])
		// if !(char >= "A" && char<= "Z") && !(char >= "a" && char<= "z"){
		// 	return false
		// }
		if !(name[i] >= 'A' && name[i] <= 'Z') && !(name[i] >= 'a' && name[i] <= 'z') {
			return false
		}
	}
	return true
}
