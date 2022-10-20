/**
长度为n的数组,操作k次
每次操作是将数组中的某个数改为 这个数的二进制数中 1 的个数
问：应该如何操作使得数组之和最小？输出最小的和

k <= 100
n很大
*/
package main

import (
	"fmt"
	"sort"
)

// 100%
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	nums := make([]int, n)
	arrs := make([]int, n)   // 转化成二进制的数
	chas := make([][]int, n) // 操作一次，整体减小的最小值
	sum := 0
	for i := range nums {
		fmt.Scan(&nums[i])
		arrs[i] = erjinzhi(nums[i])
		cha := nums[i] - arrs[i]
		chas[i] = []int{i, cha}
		sum += nums[i]
	}
	sort.Slice(chas, func(i, j int) bool {
		return chas[i][1] > chas[j][1]
	})
	// 每次操作后，需要将旧的数对应的cha删除，插入新的数对应的cha
	insert := func(i, newcha int) {
		cha := []int{i, newcha}
		if len(chas) == 1 { // 只有一个cha
			chas = [][]int{}
			chas = append(chas, cha)
			return
		}

		chas = chas[1:]
		if newcha >= chas[0][1] { // 队首
			chas = append([][]int{cha}, chas...)
		} else if newcha < chas[len(chas)-1][1] { // 队尾
			chas = append(chas, cha)
		} else { // 队中
			for j := range chas {
				if j > k { // 不加这句就会// 46.67% 超时
					return
				}
				if chas[j][1] < newcha {
					chas = append(chas[:i], append([][]int{cha}, chas[i:]...)...)
				}
			}
		}
	}
	// 每次操作
	for ; k != 0; k-- {
		i, cha := chas[0][0], chas[0][1]
		sum -= cha
		nums[i] = arrs[i]
		arrs[i] = erjinzhi(nums[i])
		newcha := nums[i] - arrs[i]
		insert(i, newcha)
	}
	fmt.Println(sum)
}

// 返回将一个数转化成二进制后，中1的个数
func erjinzhi(a int) int {
	res := 0
	for a != 0 {
		if a&1 == 1 { // 方便计数
			res++
		}
		a >>= 1
	}
	return res
}
