/**
不一样的小目标

题目描述：
小A和小B都是搬砖人，因为各自的身体素质不同，所以他们为自己制定了不同的小目标。小A可以一次搬a1块砖，但是每搬一次就要休息b1分钟，他的目标是一天搬c1块砖;同样的，小B可以一次搬a2块砖，但是每搬一次就要休息b2分钟，他的目标是一天搬c2块砖。

每个人情况不同，只要完成自己的目标就好，已知某天小A和小B同时开始搬砖，请问他们谁先完成自己的目标？如果是小A则输出‘A’，如果是小B则输出‘B’，如果两个人同时完成，则输出‘A&B’。


输入描述
第一行是一个整数T，表示数据组数（1≤T≤100）

接下来T行，每行有六个正整数a1,b1,c1,a2,b2,c2,含义如题所示。(a1,b1,c1,a2,b2,c2<=10^9)

输出描述
输出包含T行，每行一个字符串如题所示。

样例输入
2
7 9 19 9 1 89
7 9 95 4 5 9
样例输出
B
B
 */

package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for ; N > 0; N-- {
		var a1, b1, c1, a2, b2, c2 int
		fmt.Scan(&a1, &b1, &c1, &a2, &b2, &c2)
		t1, t2 := comp(a1, b1, c1), comp(a2, b2, c2)
		if t1 < t2 {
			fmt.Println("A")
		} else if t1 > t2 {
			fmt.Println("B")
		} else {
			fmt.Println("A&B")
		}
	}
}

func comp(a, b, c int) int {
	var t int
	if c%a == 0 {
		t = c / a * b
	} else {
		t = (c/a + 1) * b
	}
	// t = c / a * b
	return t
}
