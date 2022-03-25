/*

5个数，每次每4个数减1，结果不能小于0， 最多可以减多少次
*/

package main

import (
	"fmt"
)

func main(){
	var n int
	fmt.Scan(&n)

	for ;n >0;n--{
		var a, b , c, d, e int
		fmt.Scan(&a, &b, &c, &d, &e)
		nums := []int{a,b,c,d,e}
		// sort.Ints(nums)
		// fmt.Println(nums)





		nums[1] += nums[0]

		var res int
		if nums[1] < nums[2]{
			res = nums[1]
		}else{
			res = nums[2]
		}
		// nums = nums[1:]
		// sort.Ints(nums)
		// res = nums[0]
		fmt.Println(res )

	}
}
//
// func findMin(nums []int) (int , int ){
// 	index := 0
// 	minnum := math.MaxInt32
// 	for i, num := range nums{
// 		if num < minnum{
// 			index = i
// 			minnum = num
// 		}
// 	}
// 	return index, minnum
// }
//

