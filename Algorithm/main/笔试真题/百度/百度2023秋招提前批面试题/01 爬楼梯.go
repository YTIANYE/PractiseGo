/**

上楼梯：每次可以 1 2 3个台阶
*/

package main

import "fmt"

// 方法三：递归+记录
func main() {


	hash := make(map[int]int)
	hash[0] = 0
	hash[1] = 1
	hash[2] = 2
	hash[3] = 4

	var fun func(int) int
	fun = func(n int) int {
		if hash[n] != 0 {
			return hash[n]
		}

		//res := 0
		//if n -1 > 0 {
		//	res += fun(n-1)
		//}
		//if n - 2 > 0 {
		//	res += fun(n-2)
		//}
		//if n - 3 > 0 {
		//	res += fun(n-3)
		//}
		//return res
		hash[n] = fun(n-3) + fun(n-2) + fun(n-1)
		return hash[n]
	}

	for i := 1; i < 10; i++ {
		fmt.Println(fun(i))
	}

}

//方法一：动态规划
//func fun(n int){
//
//	dp := make([]int, n+1)
//	dp[0] = 0
//	for i := 1; i <= n; i++ {
//		var num1, num2, num3 int
//		if i-1 >= 0 {
//			num1 = dp[i-1]
//		}
//		if i-2 >= 0 {
//			num2 = dp[i-2]
//		}
//		if i-3 >= 0 {
//			num3 = dp[i-3]
//		}
//		dp[i] += num1 + num2 + num3
//	}
//	fmt.Println(dp[n])
//}
//

// 方法二：递归
//func fun(n int) int {
//	if n == 0 {
//		return 0
//	} else if n == 1 {
//		return 1
//	} else if n == 2 {
//		return 2
//	} else if n == 3 {
//		return 4 // 111 12 21 3
//	}
//
//	//res := 0
//	//if n -1 > 0 {
//	//	res += fun(n-1)
//	//}
//	//if n - 2 > 0 {
//	//	res += fun(n-2)
//	//}
//	//if n - 3 > 0 {
//	//	res += fun(n-3)
//	//}
//	//return res
//	return fun(n-3) + fun(n-2) + fun(n-1)
//}
