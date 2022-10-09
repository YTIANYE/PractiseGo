/**
黑白翻转
时间限制： 3000MS
内存限制： 589824KB
题目描述：
有n个黑白棋子，它们的一面是黑色，一面是白色。它们被排成一行，位置可以用标号1~n来表示。一开始，所有的棋子都是黑色向上，有q次操作，每次操作将位置标号在区间[L, R]内的所有棋子翻转（原来黑色变白色，原来白色变黑色）。请在每次操作后，求这n个棋子中，黑色向上的棋子个数。



输入描述
第一行两个整数 n, q，1 <= n <= 1018, q <= 300;

后面q行，每行两个整数 L,R，1 <= L <=R <= n。

输出描述
q行，每行一个整数，表示每次操作后黑色向上的棋子个数。


样例输入
100 2
1 30
21 40
样例输出
70
70
*/

package main

import (
	"fmt"
)

func main() {
	var n int64
	var q int
	fmt.Scan(&n, &q)
	//
	bai := make(map[int64]int64)
	res := n
	var fun func(l, r int64)
	fun = func(l, r int64) {
		if len(bai) == 0 {
			bai[l] = r
			res -= r - l + 1
			fmt.Println(res)
			return
		}
		for left, right := range bai {
			if l > right || r < left {
				continue
			}

			delete(bai, left)
			res += (right - left + 1)
			if l > left {
				bai[left] = l - 1
				res -= l - 1 - left + 1
			} else if l < left {
				bai[l] = left - 1
				res -= left - 1 - l + 1
			}

			if r < right {
				bai[r+1] = right
				res -= right - (r + 1) + 1
			} else if r > right {
				bai[right+1] = r
				res -= r - (right + 1) + 1
			}
			//
			fmt.Println(res)
			return
		}
		//

	}
	//
	for q != 0 {
		q--
		var l, r int64
		fmt.Scan(&l, &r)
		fun(l, r)
	}

}
