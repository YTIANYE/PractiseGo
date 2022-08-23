/**
数组中第k个最大元素

在未排序的数组中找到第k个最大的元素，（注意，不是第k个不同的元素），即数组中存在相同元素

3,2,3,1,2,4,5,5,6  k= 4
4

3,2,1,5,6,4		k=2
5

 */

package main
import fmt "fmt"
import "math"

func main() {
	// fmt.Printf("Hello, World!");
	//  nums := []int{3,2,1,5,6,4}
	//  k := 2
	nums := []int{3,2,3,1,2,4,5,5,6}
	k := 4
	fmt.Println(kth(nums, k))
}

func kth(nums []int, k int ) int {
	temps := make([]int, k)
	n := len(nums)
	for i:=0;i<k;i++ {
		temps[i] = nums[i]
	}
	index, minNum := findMin(temps)
	//
	for i:=k;i<n;i++ {
		if nums[i] > minNum {
			temps[index] = nums[i]
			index, minNum = findMin(temps)
		}
	}

	_, res := findMin(temps)
	return res
}

func findMin(nums []int) (int, int ) { // 返回索引和value
	// res := 1000
	res := math.MaxInt64
	index := -1
	for i, num := range nums {
		if num < res {
			res = num
			index = i
		}
	}
	return index , res
}