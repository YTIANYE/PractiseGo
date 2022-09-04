/**
乒乓球比赛
给出两人分数，判断至少多少次之后可以使a获胜
 */
package main

import "fmt"
// 100%
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	count := 0
	for !(a >= 11 && a - b >= 2) {
		count++
		a++
	}
	fmt.Println(count)
}