/**
合并两个数组，并由小到大排序
 */

package main

import (
	"fmt"
	"sort"
)
func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	fmt.Printf("Hello World!\n");
}

func hebing(nums1, nums2 []int) []int{
	sort.Ints(nums1)
	sort.Ints(nums2)

	i:=0
	j:=0
	res := []int{}
	for i<len(nums1) && j < len(nums2){
		if nums1[i] < nums2[j]{
			res = append(res, nums1[i])
			i++
		}else{
			res = append(res, nums2[j])
			j++
		}
	}

	for i < len(nums1){
		res = append(res, nums1[i])
		i++
	}
	for j < len(nums2){
		res = append(res, nums2[j])
		j++
	}

	return res
}