/*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。

 

示例 1：

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
示例 2：

输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。
 

提示：

n == ratings.length
1 <= n <= 2 * 104
0 <= ratings[i] <= 2 * 104

*/

package main

import "fmt"

func candy(ratings []int) int {

	n := len(ratings)
	if n == 1 {
		return 1
	}

	candies := make([]int, n)
	candies[0] = 1
	i := 1
	for i < n {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
			i++
		} else if ratings[i] < ratings[i-1] {
			j := i
			for j < n && ratings[j] < ratings[j-1] {
				j++
			}
			nexti := j
			j--
			k := 1
			for ; j != i-1; j-- {
				candies[j] = k
				k++
			}

			candies[i-1] = max(candies[i-1], k)
			i = nexti
		} else {
			candies[i] = 1
			i++
		}
	}

	// fmt.Println(candies)
	return sum(candies)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}

func main(){
	nums := []int{1,0,2}
	fmt.Println(candy(nums))
}
