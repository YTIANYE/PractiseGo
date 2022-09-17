/*
两个妖怪生命值a b
两招x和y
x 是对一个妖怪造成x的伤害
y 是对两个妖怪都造成y的伤害
问：至少需要出多少招？
*/

package main

import "fmt"

// 我的题解：完全通过样例 100%
func main() {

	var a, b, x, y int
	fmt.Scan(&a, &b, &x, &y)
	res := 0

	if x > y*2 {
		// 留下a b 的最后一次
		if a%x == 0 {
			res += a/x - 1
		} else {
			res += a / x
		}
		a -= res * x

		var temp int
		if b%x == 0 {
			temp = b/x - 1
			res += temp
		} else {
			temp = b / x
			res += temp
		}
		b -= temp * x
		// 如果最后1个y可以解决，就不需要2个x了
		if a <= y && b <= y {
			res += 1
		} else {
			res += 2
		}

	} else {
		// 使 a 最小
		if a > b {
			a, b = b, a
		}

		if a%y == 0 {
			res += a / y
			b -= a
		} else {
			res += a/y + 1
			b -= res * y
		}

		// 单独处理b
		if b > 0 {
			if b%x == 0 {
				res += b / x
			} else {
				res += b/x + 1
			}
		}
	}
	fmt.Println(res)

}
