package main

import "fmt"

func main() {
	a := []int{1, 5, 9, 10, 11, 12, 19, 20, 22, 24, 28}
	// 找到有序数组中，>=t的第一个数
	var biaFind func(t int, l, r int) int
	biaFind = func(t, l, r int) int {
		if l >= r {
			return a[l]
		}
		mid := l + (r-l)/2
		if a[mid] < t { // 小于的全删去
			return biaFind(t, mid+1, r)
		} else {
			return biaFind(t, l, mid) // 查找第一个大于等于的方法
		}
	}

	for i:=0;i<28;i++ {
		res := biaFind(i, 0, len(a)-1)
		fmt.Println(res)
	}
}
