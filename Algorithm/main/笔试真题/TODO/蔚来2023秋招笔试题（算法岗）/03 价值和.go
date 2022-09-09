
/**
n个数，
K次特殊操作，每次可以选择一个数x 使得 x = x*x（每个数仅限一次）
m次选数，每次选P个数，求每次选择的数的最大和是多少



5 5 4
-3 -2 2 3 4
1
2
3
4
5



16
25
34
38
40
 */
package main

import (
	"fmt"
	"sort"
)

// 20%
func main() {
	var n,m,k int
	fmt.Scan(&n, &m, &k)
	vs := make([]int, n)
	for i:= range vs{
		fmt.Scan(&vs[i])
	}
	sort.Slice(vs, func(i, j int) bool {
		return abs3(vs[i]) > abs3(vs[j])
	})
	//
	for i:=0;i<m;i++ {
		var p int
		fmt.Scan(&p)

		if k != 0 {
			vs[i] = vs[i] * vs[i]
			k--
		}

		temp := make([]int, len(vs))
		copy(temp, vs)
		sort.Ints(temp)

		res := 0
		for j:=0;j<p;j++ {
			res += temp[len(temp)-1-j]
		}
		fmt.Println(res)
	}
	//

}


func abs3(a int) int {
	if a > 0 {
		return a
	}
	return -a
}