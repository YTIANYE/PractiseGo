/*
给定一个字符串a
全部由0和1构成，
找出两个不同的字串，使其中0的个数和1的个数相同
输出长度最大的任意两个字符串的下标
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Scan(&a)

	var s1, t1, s2, t2 int
	n := len(a)
	changdu := 1
	dp := make([][2]int, n+1) // index:起始点， value 0 1 的个数
	// 0 不用
	// maxChangDu := 0

	for ; changdu < n; changdu++ {
		temp := changdu - 1            // 不算开头，剩下的部分的长度
		for i := 1; i <= n-temp; i++ { // i dp 开头位置索引
			index := i - 1                               // 字符串的起始索引
			end := index + temp                          // 截至位置的索引
			num, _ := strconv.Atoi(string(a[end])) // 字符串上的1 或 0
			dp[i][num]++
		}

		// 当下长度中，是否有符合要求的
		// newDP := dp[:n-temp]
		// sort.Slice(newDP, func(i, j int) bool {
		// 	return newDP[i][0] < newDP[j][0]
		// })

		hashMap := make(map[int]int)  // key : 0的个数  value 索引
		for i:=1;i<= n-temp;i++{
			key := dp[i][0]
			count, ok := hashMap[key]
			if ok{
				j := i
				i := count
				s1, t1, s2, t2 = i, i+temp, j, j+temp
				break
			}else{
				hashMap[key] = i
			}
		}

		// find := false
		// for i := 1; i <= n-temp; i++ {
		// 	for j := i + 1; j <= n-temp; j++ {
		// 		if dp[i][0] == dp[j][0] && dp[i][1] == dp[j][1] {
		// 			s1, t1, s2, t2 = i, i+temp, j, j+temp
		// 			find = true
		// 			break
		// 		}
		// 	}
		// 	if find {
		// 		break
		// 	}
		// }

	}
	fmt.Println(s1, t1, s2, t2)
}


