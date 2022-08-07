/**
涂色
时间限制： 3000MS
内存限制： 589824KB
题目描述：
小A的门前有n个排成一排的栅栏，编号分别为1,2，...，n。每个栅栏都是红色或者蓝色的。但小A觉得目前的上色方案看起来有些杂乱，便想要重新对栅栏进行涂色。具体地，小A认为，如果栅栏的颜色交替次数多于1次，那么就是杂乱的，否则就是整齐的。换言之，如果栅栏是全红/全蓝/前一段红后一段蓝/前一段蓝后一段红，那么都能符合小A的要求。请问小A至少需要对几个栅栏进行重新涂色，才能满足他的要求呢？



输入描述
第一行是一个整数n，表示有n个栅栏，1<=n<=100000。

第二行是一个字符串s，字符串只包含’r’和’b’，对于第i个字符，若为’r’表示第i个栅栏为红色，若为’b’则表示第i个栅栏为蓝色。

输出描述
一行一个整数，表示小A需要进行重新涂色的最少栅栏数。


样例输入
5
rbrbb
样例输出
1

提示
输入样例2：

6

rbrbbrr

输出样例2：

2

输入样例3

4

rrbr

输出样例3

1

样例解释：将第二个栅栏从’b’涂为’r’，则重新涂色后的栅栏为rrrbb，符合小A的要求，可以证明1是需要重新涂色的最少栅栏数。
 */

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var str string
	fmt.Scan(&str)
	//
	var rleft, rright, bleft, bright int
	for i := range str {
		if str[i] == 'r' {
			rright++
		}else {
			bright++
		}
	}

	//
	res := min(rright, bright)
	for i := 0; i<n;i++ {
		if str[i] == 'r' {
			rleft++
			rright--
		}else {
			bleft++
			bright--
		}

		temp := min(rleft, bleft) + min(rright, bright)
		res = min(res, temp)
	}
	fmt.Println(res )
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
