/*
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，
其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积,
即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

示例:
输入: [1,2,3,4,5]
输出: [120,60,40,30,24]

提示：
所有元素乘积之和不会溢出 32 位整数
a.length <= 100000
*/

package main

import "fmt"

/*
1	2	3	4

sum 2
2	1

sum 2 	sum 12
2	1	4	3

sum 24
24 	12	 8	 6


*/

// 我的题解：分治
/*
执行用时：20 ms, 在所有 Go 提交中击败了92.55%的用户
内存消耗：9.6 MB, 在所有 Go 提交中击败了11.97%的用户
*/
func constructArr1(a []int) []int {
	if len(a) == 0 {
		return nil
	}

	b := []int{}
	for i := 0; i < len(a); i++ {
		b = append(b, 1)
	}

	_, b = fun(a, b, 0, len(a)-1)
	return b
}

func fun(a []int, b []int, left, right int) (int, []int) { // suml 左边的数的乘积  sumr 右边的数的乘积
	if left == right {
		return a[left], b
	}

	mid := (left + right) / 2
	suml, b := fun(a, b, left, mid)
	sumr, b := fun(a, b, mid+1, right)

	for i := left; i <= mid; i++ {
		b[i] *= sumr
	}
	for i := mid + 1; i <= right; i++ {
		b[i] *= suml
	}
	return suml * sumr, b
}

// 我实现的精选题解：表格分区【上三角和下三角】
/*
执行用时：20 ms, 在所有 Go 提交中击败了92.59%的用户
内存消耗：9.5 MB, 在所有 Go 提交中击败了12.43%的用户
*/
func constructArr(a []int) []int {
	b := []int{}
	for i := 0; i < len(a); i++ {
		b = append(b, 1)
	}
	temp := 1
	// 下三角
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	// 上三角
	for i := len(a) - 2; i > -1; i-- {
		temp *= a[i+1]
		b[i] *= temp
	}

	return b
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	// nums := []int{}
	fmt.Println(constructArr(nums))
}
