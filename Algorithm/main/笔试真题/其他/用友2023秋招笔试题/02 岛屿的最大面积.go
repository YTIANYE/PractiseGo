/**
8 13
0,0,1,0,0,0,0,1,0,0,0,0,0
0,0,0,0,0,0,0,1,1,1,0,0,0
0,1,1,0,1,0,0,0,0,0,0,0,0
0,1,0,0,1,1,0,0,1,0,1,0,0
0,1,0,0,1,1,0,0,1,1,1,0,0
0,0,0,0,0,0,0,0,0,0,1,0,0
0,0,0,0,0,0,0,1,1,1,0,0,0
0,0,0,0,0,0,0,1,1,0,0,0,0
*/
package main

import "fmt"

func main() {

	//// 输入
	//tu := [][]int{}
	//var s string
	//fmt.Scan(&s)
	//arr := strings.Split(s, ",")
	//n := len(arr)
	////fmt.Println(m)
	//nums := []int{}
	//for i := range arr {
	//	num, _ := strconv.Atoi(arr[i])
	//	nums = append(nums, num)
	//}
	//tu = append(tu, nums)
	//
	//tnum := []int{}
	//for{
	//	var a int
	//	if _, ok := fmt.Scan(&a); ok != io.EOF {
	//		fmt.Scan(&a)
	//		tnum = append(tnum, a)
	//		if len(tnum) ==  n {
	//			temp := make([]int, len(tnum))
	//			copy(temp, tnum)
	//			tnum = []int{}
	//			tu = append(tu, temp)
	//		}
	//	}else {
	//		break
	//	}
	//}
	//m := len(tu)  // 行
	//for i:= range tu {
	//	fmt.Println(tu[i])
	//}

	var m, n int
	fmt.Scan(&m, &n)
	tu := [][]int{}
	for i := 0; i < m; i++ {
		nums := []int{}
		for j := 0; j < n; j++ {
			var num int
			fmt.Scan(&num)
			nums = append(nums, num)
		}
		tu = append(tu, nums)
	}

	//  是否访问标记
	tag := make([][]bool, m)
	for i := range tag {
		tag[i] = make([]bool, n)
	}
	// 探索
	var search func(x, y, num int) int
	search = func(x, y, num int) int {
		if x < 0 || x >= m || y < 0 || y >= n {
			return num
		}
		if tu[x][y] == 0 || tag[x][y] == true {
			return num
		}
		num++
		tag[x][y] = true
		num = search(x+1, y, num)
		num = search(x-1, y, num)
		num = search(x, y+1, num)
		num = search(x, y-1, num)
		return num
	}
	// 计算
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tag[i][j] == true || tu[i][j] == 0 { // 访问过 或者是 水
				continue
			}
			num := search(i, j, 0)
			res = max2(res, num)
		}
	}
	fmt.Println(res)
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
