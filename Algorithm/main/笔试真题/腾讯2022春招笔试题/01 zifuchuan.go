/**
n个长度相等的字符串，每个字符串中只包含数字
n <= 9
遍历每一列的数字组成一个整数（去掉前导0）
将所有遍历到的数字由小到大输出
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 通过样例 100%
func main() {
	var n int
	fmt.Scan(&n)
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&strs[i])
	}

	// 遍历整数
	l := len(strs[0])
	nums := make([]int, l)
	for i := 0; i < l; i++ {
		num := ""
		for j := 0; j < n; j++ {
			num = num + string(strs[j][i])
		}
		nums[i], _ = strconv.Atoi(num)
	}

	sort.Ints(nums)

	// 打印结果
	for i := range nums {
		if i == 0 {
			fmt.Print(nums[i])
		} else {
			fmt.Print(" ")
			fmt.Print(nums[i])
		}
	}

}

/**
4
0000
0101
1011
0111

10 11 101 111
// 0 101 111 1011

*/
