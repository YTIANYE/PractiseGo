/**
每次参加游戏的时间长度为m
一共n组游戏，每个小游戏的时间为[left, right)
问完成所有游戏，需要进行多少次？
 */
// 可以引⼊的库和版本相关请参考 “环境说明”
// 必须定义⼀个 包名为 `main` 的包

package main

import "sort"
// import "fmt"

// 本题面试官已设置测试用例
func calcGameCount(m int, n [][]int) int {
	// 在这⾥写代码
	res := []int{}
	l := len(n)
	// 排序
	// for i:=0;i<l;i++ {
	//     for j:=i+1;j<l;j++ {
	//         if !bijiao(n[j-1], n[j]) {
	//             n[j-1], n[j] = n[j], n[j-1]
	//         }
	//     }
	// }

	sort.Slice(n, func(i, j int) bool {
		return n[i][0] < n[j][0] || (n[i][0] == n[j][0] && n[i][1] < n[j][1])
	})
	// fmt.Println(n)


	for i:=0;i<l;i++ {
		j := 0
		left := n[i][0]
		right := n[i][1]
		tag := false
		for j< len(res) {
			if left >= res[j] {
				res[j] = right
				tag = true  // 接上了
				break
			}
			j++
		}
		if !tag{
			res = append(res, right)
		}
	}

	return len(res)
}

// func bijiao(a []int, b []int) bool {  // <= false , > true 
//     if a[0] < b [0] {
//         return false 
//     }else if a[0] > b[0] {
//         return true 
//     }else {
//         if a[1] <= b[1] {
//             return false
//         }else {
//             return true 
//         }
//     }

// }