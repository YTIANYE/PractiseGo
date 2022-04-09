/*
问：一个整数n中存在的圆圈的个数？
0， 6， 9 有一个
8 有两个
其他没有
*/

package main

import "fmt"

// 我的题解：通过样例 100%
func main(){
	var n int
	fmt.Scan(&n)

	hashmap := map[int]int{
		0:1,
		1:0,
		2:0,
		3:0,
		4:0,
		5:0,
		6:1,
		7:0,
		8:2,
		9:1,
	}

	res := 0
	for n > 0{
		num := n % 10
		n /= 10
		res += hashmap[num]
	}
	fmt.Println(res )
}
