package main

import "fmt"

/*测试map默认value*/
func testmap() {
	set := make(map[string]int)
	s := "abcd"
	for _, str := range s {
		set[string(str)] += 1
	}
	fmt.Println(set["e"]) // map中找不到时默认返回0
}
