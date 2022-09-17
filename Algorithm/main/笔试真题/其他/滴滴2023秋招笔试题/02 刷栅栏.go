/**
1~~1000000000个栅栏
1和2号两种油漆 分别需要刷p, q次
可以刷n次，
每次在l和r之间（应该是左闭右闭）刷t号油漆

输入 n,p,q
接下来的3*n矩阵，每一列表示第i次刷的 l , r, t


问：最后有多少个栅栏刷的次数是符合要求的？
*/
package main

import "fmt"

func main() {
	nums1 := [][]int{} // 第二维 ： l r count
	nums2 := [][]int{}

	var n, p, q int
	fmt.Scan(&n, &p, &q)
	ls, rs, ts := make([]int, n), make([]int, n), make([]int, n)
	for i := range ls {
		fmt.Scan(&ls[i])
	}
	for i := range rs {
		fmt.Scan(&rs[i])
	}
	for i := range ts {
		fmt.Scan(&ts[i])
	}
	//
	fun := func(nums [][]int, l, r int) {
		for i, num := range nums {
			if l >= num[0] && l <= num[1] {
				temp := [][]int{}
				if l >= i {
					temp = append(temp, []int{i, l, num[2]})
				}

				if r < num[1] {
					temp = append(temp, []int{l+1, r, num[2] + 1})
					temp = append(temp, []int{r+1, num[1], num[2]})
				} else if r > num[1] {
					temp = append(temp, []int{l+1, num[1], num[2] + 1})
					temp = append(temp, []int{num[1]+1, r, num[2]})
				}else {
					temp = append(temp, []int{l+1, r, num[2] + 1})
				}
				nums = append(nums[:i], append(temp, nums[i+1:]...)...)
				break
			}
		}
	}
	//
	for i := 0; i < n; i++ {
		l, r := ls[i], rs[i]
		if ts[i] == 1 {
			fun(nums1, l, r)
		} else {
			fun(nums2, l, r)
		}
	}
	//
	res := 0
	// 统计 nums1 和 nums2 的每个段符号要求的
	fmt.Println(res)
}
