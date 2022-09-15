/**
题目：
输入n表示n个树节点
每个树节点编号为1...n
接下来的n - 1 行为边上的两个结点

一次操作：该节点及其子树的值+1
初始时每个结点的值都是 1

问：最少多少次操作，可以使每个结点的值等于结点编号
*/
package main

import "fmt"



// 90%
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n+1) // 注意长度
	for i := range nums {
		nums[i] = 1
	}
	//
	hash := make(map[int]int, n)

	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		if b < a {
			a, b = b, a
		}
		hash[b] = a // a < b
	}

	//
	var res = 0
	for i := 2; i <= n; i++ {

		temp := i - nums[hash[i]]
		nums[i] = i
		res += temp

	}
	//
	fmt.Println(res)
}


//
//// 5%
//func main() {
//	var n int
//	fmt.Scan(&n)
//	nums := make([]int, n+1)  // 注意长度
//	for i:= range nums {
//		nums[i] = 1
//	}
//	//
//	hash := make(map[int][]int, n)
//	for i:=1;i<=n;i++ {
//		hash[i] = []int{}
//	}
//	for i:=0;i<n-1;i++ {
//		var a, b int
//		fmt.Scan(&a, &b)
//		if b < a {
//			a, b = b, a
//		}
//		hash[a]  = append(hash[a], b)// a < b
//	}
//	for i := 1;i<=n;i++ {
//		queue := make([]int, len(hash[i]))
//		copy(queue , hash[i])
//		for len(queue) != 0 {
//			tque := make([]int, len(queue))
//			copy(tque, queue)
//			queue = []int{}
//			for len(tque) != 0 {
//				node := tque[0]
//				hash[i] = append(hash[i], node)
//				t := make([]int, len(hash[node]))
//				queue = append(queue, t...)
//				if len(tque) == 1 {
//					tque = []int{}
//				}else {
//					tque = tque[1:]
//				}
//			}
//		}
//	}
//
//
//	//
//	var res = 0
//	for i:=1;i<=n;i++ {
//		if nums[i] >= i {
//			continue
//		}
//		temp := i - nums[i]
//		nums[i] += temp
//		res += temp
//		//
//		for _, num := range hash[i] {
//			nums[num] += temp
//		}
//
//	}
//	//
//	fmt.Println(res)
//}