/**
数组A中存储N个点的坐标，坐标可能相同
从中尽可能取出最多的点，
这些点之间的距离都可以是M的倍数



-3, -2, 1, 0, 8, 7, 1
3
输出 4

7,1,11,8,4,10
8
输出 1

*/

package main

func Solution2(A []int, M int) int {
	// write your code in Go (Go 1.4)

	minNum := A[0]
	for _, num := range A {
		minNum = min(minNum, num)
	}

	// 处理负数的方式，是使所有的数加上最小数的绝对值，多有的数都变成整数
	// (-2%3) // -2
	if minNum < 0 {
		for i := range A {
			A[i] += -minNum
		}
	}

	hash := make(map[int]int) // 长度是 M，key是每个数%M后得到的余数
	for _, num := range A {
		hash[num%M]++
	}

	res := 1 //只有一个点组成，距离是0，也算
	for _, value := range hash {
		res = max2(res, value)
	}
	return res
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//
//func main() {
//	A := []int{-3, -2, 1, 0, 8, 7, 1}
//	// 0 1 4 3 11 10 4
//	M := 3
//
//	//A := []int{7,1,11,8,4,10}
//	//M := 8
//
//	fmt.Println(Solution(A, M))
//}
