/*
输入n个地点m个边
输入两行数，每一行m个，代表地点之间存在道路

输入q代表询问次数，接下来q行
每行输入两个地点，判断两个地点之间是否有直接的道路（经过其他地点到达的不算）
*/

package main

import (
	"fmt"
	"strconv"
)

// 我的题解：通过样例 100%
func main(){
	var n , m int
	fmt.Scan(&n, &m)
	us := make([]int, m)
	vs := make([]int, m)
	for i:=0;i<m;i++{
		var num int
		fmt.Scan(&num)
		us[i] = num
	}
	for i:=0;i<m;i++{
		var num int
		fmt.Scan(&num)
		vs[i] = num
	}

	hashmap := make(map[string]bool)
	for i:=0;i<m;i++{
		a, b := min(us[i], vs[i])
		u, v:= strconv.Itoa(a), strconv.Itoa(b)
		key := u + "-" + v
		hashmap[key] = true
	}
	// fmt.Println(hashmap)

	// 询问
	var q int
	fmt.Scan(&q)
	for i:=0;i<q;i++{
		var u, v int
		fmt.Scan(&u, &v)
		a, b := min(u, v)
		key := strconv.Itoa(a) + "-" + strconv.Itoa(b)
		// fmt.Println(key + "#")
		if hashmap[key]{
			fmt.Println("Yes")
		}else{
			fmt.Println("No")
		}
	}
}

func min(a, b int) (int, int){
	if a < b {
		return a, b
	}
	return b, a
}
