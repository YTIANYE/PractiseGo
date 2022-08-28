/**
题目：
一个字符串s 至少包含两个red即为漂亮串，
求长度为n的漂亮串有多少种？
结果取模 （1000000000 + 7）
 */

package main

/**
6
redred
1
//
7
?redred
red?red
redred?
3*26
//
red = @
?@@ + ! 有4种
3 * 4 * 26^2
//
？！@@ + #
3*4*5 * 26^3

*/
import "fmt"
//
//// 20%
//func main() {
//	modNum := 1000000000 + 7
//	var n int
//	fmt.Scan(&n)
//	//
//	//  k
//	// C
//	//  m
//	// 从m中选k 个
//	xuanC := func(k, m int) int {
//		fenzi := 1
//		fenmu := 1
//		for i := 1; i <= k; i++ {
//			fenzi *= m
//			m--
//			fenmu *= i
//		}
//		return fenzi / fenmu
//	}
//
//	// 几个red的情况
//	jigered := func(k int) int {
//		m := n - k*3 + k
//		rednum := xuanC(k, m)
//		res := rednum
//		for i := 0; i < n-k*3; i++ {
//			res = res * 26 % modNum
//		}
//		return res
//	}
//
//	//rednum := 0 // 两个red位置不同的种类数
//	//for i := 0; i < n-5; i++ {
//	//	rednum = (rednum + (n - (i + 3) - 2)) % modNum
//	//}
//	//
//	//res := rednum
//	//for i := 0; i < n-6; i++ {
//	//	res = (res * 26) % modNum
//	//}
//
//	//
//	res := 0
//	end := n / 3 // 最多这么多个red
//	for i := 2; i <= end; i++ {
//		res = (res + jigered(i)) % modNum
//	}
//	fmt.Println(res)
//}

func main() {
	modNum := 1000000000 + 7
	var n int
	fmt.Scan(&n)
	//
	//  k
	// C
	//  m
	// 从m中选k 个
	xuanC := func(k, m int) int {
		fenzi := 1
		fenmu := 1
		for i := 1; i <= k; i++ {
			fenzi *= m
			m--
			fenmu *= i
		}
		return fenzi / fenmu
	}



	rednum := xuanC(2, n-6+2)

	res := rednum
	for i := 0; i < n-6; i++ {
		res = (res * 26) % modNum
	}
	fmt.Println(res)
}
