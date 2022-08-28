/**
一个字符串的k位置之前的字符大写，其余字符小写
 */

package main

import (
	"fmt"
	"strings"
)

// 100%
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var s string
	fmt.Scan(&s)

	res := strings.ToUpper(s[:k])
	res = res + strings.ToLower(s[k:])
	fmt.Println(res )
}
