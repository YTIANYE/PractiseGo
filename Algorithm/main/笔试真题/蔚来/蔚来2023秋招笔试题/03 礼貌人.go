/**

题目：n个人序号为1...n，排成一排过q扇门，
门有A 和 B两种，
过A门，队首的人去最后
过B门，整个队伍向后转，形成新的队伍
输出最后的队伍顺序

*/


/**
5
4
A
A
B
A

1
5
4
3
2
 */
package main

import "fmt"

// 100%
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	nums := make([]int, n)
	for i:=0;i<n;i++ {
		nums[i] = i+1
	}

	tag := true // true 从头开始  false 从尾开始

	fun1 := func() {
		if tag {
			nums = append(nums[1:], nums[0])
		}else {
			nums = append([]int{nums[n-1]}, nums[:n-1]...)
		}
	}

	fun2 := func() {
		tag = !tag
	}


	for i:=0;i<m;i++ {
		var a string
		fmt.Scan(&a)
		if a == "A" {
			fun1()
		}else if a == "B" {
			fun2()
		}
	}
	// 输出
	if tag {
		for i:=0;i<n;i++ {
			fmt.Println(nums[i])
		}
	}else {
		for i:= n-1;i>=0;i-- {
			fmt.Println(nums[i])
		}
	}

}