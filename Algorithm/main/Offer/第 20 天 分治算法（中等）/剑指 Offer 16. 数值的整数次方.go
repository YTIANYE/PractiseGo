/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。


示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25

*/

package main

import "fmt"

/*// 1 我的题解：暴力解法：超时
func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	res := x
	var count int
	if n < 0 {
		count = n * (-1)
	} else {
		count = n
	}
	for i := 1; i < count; i++ {
		res *= x
	}
	if n < 0 {
		res = 1.0 / res
	}
	return res
}

// 2 我的题解：改进暴力：超时
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	flag := true
	if n < 0 {
		flag = false
		n *= -1
	}

	result := fun(x, n)

	if flag == false {
		result = 1.0 / result
	}
	return result
}

// 返回将计算结果

func fun(x float64, n int) float64 {

	cifangs := cf(n)

	res := 1.0

	num := 1.0
	j := 0
	for i := len(cifangs) - 1; i >= 0; i-- {

		temp := 1
		for k := 0; k < cifangs[i]; k++ {
			temp *= 2
		}

		for j < temp {
			num = num * x
			j++
		}
		res *= num

	}
	return res
}

// 返回n的二进制数中为1的位数
func cf(n int) []int {

	cifangs := []int{}
	num := n
	for num > 0 {
		cifang, temp := compute(num)
		num -= temp
		cifangs = append(cifangs, cifang)
	}
	return cifangs
}

// 返回 n 中最大的2的幂 的幂数 转化为二进制数的位数
func compute(n int) (int, int) {
	count := 0
	num := 1
	for num <= n {
		count++
		num *= 2
	}
	return count - 1, num / 2
}*/

// 3 我的题解：分治
func myPow3(x float64, n int) float64 {
	if n < 0 {
		return 1 / comp(x, -n)
	}
	return comp(x, n)
}

func comp(x float64, n int) float64 {

	if n == 0 {
		return 1.0
	} else if n == 1 {
		return x
	}
	return comp(x, n/2) * comp(x, n-n/2)
}

// 4 我的题解：分治——改进版
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func myPow4(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	} else if n < 0 {
		return 1 / myPow(x, -n)
	}

	temp := myPow(x, n/2)
	if n%2 == 1 {
		return temp * temp * x
	} else {
		return temp * temp
	}
}

// 5 我实现的精选题解：快速幂
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了97.79%的用户
*/
func myPow(x float64, n int) float64{
	if x == 0{
		return 0
	}

	res := 1.0

	if n<0{
		x = 1.0/x
		n = -n
	}

	for n != 0{
		if n & 1 == 1{
			res *= x
		}
		x*= x
		n >>= 1
	}
	return res
}

func main() {
	x := 2.0
	n := 10
	// n := 21// 10101
	fmt.Println(myPow(x, n))
}
