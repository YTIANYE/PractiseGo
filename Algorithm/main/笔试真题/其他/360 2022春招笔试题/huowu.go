/*
所有的箱子堆在一起，长宽高的个数分别为l,r,h,
偷走后变成(l-2, r-2, h-1)
被偷走后的箱子共n个，问最大损失是多少?

输入 n = 4
输出 ： 41
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			jk := n / i
			for j := 1; j < jk; j++ {
				if jk%j == 0 {
					k := jk/j
					temp := maxijk(i,j,k)
					res = max(res, temp-n)
				}
			}
		}
	}

	fmt.Println(res)

}

// 返回最大乘积
func maxijk(i, j, k int) int{
	res := 0
	res = max(res, (i+1)*(j+2)*(k+2))
	res = max(res, (i+2)*(j+1)*(k+2))
	res = max(res, (i+2)*(j+2)*(k+1))
	return res
}

/*func max(a, b int) int{
	if a>b{
		return a
	}
	return b
}*/
