/**

输入商品的编号和价格
输出前N大的价格对应的商品
注意：相同价格可能不止一种商品
*/

package main

import (
	"fmt"
	"io"
	"sort"
)

// 66.67%
//func main() {
//	var n int
//	fmt.Scan(&n)
//	nums := [][]int{}
//	for {
//		var a, b int
//		if _, ok := fmt.Scan(&a, &b); ok != io.EOF{
//			nums = append(nums, []int{a, b})
//		}else {
//			break
//		}
//	}
//	//
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i][1] > nums[j][1]
//	})
//	//
//
//	i := 0
//	hash := make(map[int][]int)
//	for len(hash) != n+1 {
//		a, b := nums[i][0], nums[i][1]
//		i++
//		hash[b] = append(hash[b], a)
//	}
//	//
//	keys := []int{}
//	for key := range hash {
//		keys = append(keys, key)
//	}
//	sort.Ints(keys)
//	//
//	for j:=n;j>0;j-- {
//		fmt.Printf("%d",keys[j])
//		for _, num := range hash[keys[j]] {
//			fmt.Printf(" %d", num)
//		}
//		fmt.Println()
//	}
//}


// 100%
func main() {
	var n int
	fmt.Scan(&n)
	hash := make(map[int][]int)
	for {
		var a, b int
		if _, ok := fmt.Scan(&a, &b); ok != io.EOF{
			hash[b] = append(hash[b], a)
		}else {
			break
		}
	}
	keys := []int{}
	for key := range hash {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	//
	for j:=len(keys)-1;j> len(keys)-1-n;j-- {
		fmt.Printf("%d",keys[j])
		for _, num := range hash[keys[j]] {
			fmt.Printf(" %d", num)
		}
		fmt.Println()
	}
}