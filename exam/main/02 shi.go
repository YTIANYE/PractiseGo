/*
n个士兵排成一列，身高小的排前面，身高一样遵照名字的字典序列排序，小的在前
输入：
n
n个士兵的身高
n个士兵的名字
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 我的题解：通过样例 100%
func main() {
	var n int
	fmt.Scan(&n)
	shibing := [][]string{}

	// nums := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		shengao := strconv.Itoa(num)
		shibing = append(shibing, []string{shengao})
	}
	// names := make(map[int]string)// key 身高 value 名字
	for i := 0; i < n; i++ {
		var name string
		fmt.Scan(&name)
		shibing[i] = append(shibing[i], name)
	}

	// 自定义排序切片方式
	sort.Slice(shibing, func(i, j int) bool {
		// return shibing[i][0] < shibing[j][0] && shibing[i][1] < shibing[j][1]
		return comp(shibing[i][0] , shibing[j][0] , shibing[i][1] ,shibing[j][1])
	})

	// 打印
	for i, val := range shibing {
		if i == 0 {
			fmt.Print(val[1])
		} else {
			fmt.Print(" " + val[1])
		}
	}

}

func comp(num1, num2 , s1, s2 string) bool{
	a, _ := strconv.Atoi(num1)
	b, _ := strconv.Atoi(num2)
	if a < b {
		return true
	}else if a == b {
		return s1 < s2
	}else{
		return false
	}
}

//
// func printstring(s []string) {
// 	for i := range s{
// 		if i == 0{
// 			fmt.Println(s[i])
// 		}else{
// 			fmt.Println(" "+ s[i])
// 		}
// 	}
// 	fmt.Println('\n')
// }
