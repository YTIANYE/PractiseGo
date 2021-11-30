// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
// F(0) = 0, (1)= 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
// 斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
// 示例 1：
//
// 输入：n = 2
// 输出：1
// 示例 2：
//
// 输入：n = 5
// 输出：5
//
//
// 提示：
//
// 0 <= n <= 100
package main

import "fmt"


// 我的题解：迭代
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了70.86%的用户
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	num1 := 0
	num2 := 1
	for i := 2; i <= n; i++ {
		// num1 += num2
		// num1, num2 = num2, num1		// 像python一样的方式
		// num2 %= 1000000007
		num1, num2 = num2, num1 + num2
		num2 %= 1000000007
	}
	return num2
}


// 我的题解：递归 超时
// func fib(n int) int {
// 	if n == 0  || n == 1{
// 		return n
// 	}else {
// 		return (fib(n -1) + fib(n - 2)) % 1000000007
// 	}
// }

func main() {
	for n:= 0; n < 100; n ++{
		result := fib(n)
		fmt.Println(n, result)
	}
}