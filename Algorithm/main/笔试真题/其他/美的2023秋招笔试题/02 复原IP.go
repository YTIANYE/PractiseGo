/**

一个字符串仅有数字构成，求这个字符串拆分成四段之后能否组成一个IP地址
输出所有可以组成的IP地址

（地址中不应包含前导0 ）
 */
package main

import (
	"fmt"
	"strconv"
)

// 97%
func main() {
	var str string
	fmt.Scan(&str)
	res := []string{}

	var dfs func(s string, nums []int)
	dfs = func(s string, nums []int) {
		if len(nums) == 4 {
			if s == "" {
				res = append(res, zujian(nums))
			}
			return
		}
		if s[0] == '0' {
			return
		}
		//
		if len(nums) == 3 {
			num, _ := strconv.Atoi(s)
			if num > 255 {
				return
			}
			nums = append(nums, num)
			dfs("", nums)
			nums = nums[:len(nums)-1]
		} else {
			//
			for i := 1; i <= 3 || i < len(s); i++ {
				num, _ := strconv.Atoi(s[:i])
				if num > 255 {
					continue
				}
				nums = append(nums, num)
				dfs(s[i:], nums)
				nums = nums[:len(nums)-1]
			}
		}

	}
	//
	dfs(str, []int{})
	for i := range res {
		fmt.Println(res[i])
	}
}

func zujian(nums []int) string {
	res := ""
	for i, num := range nums {
		t := strconv.Itoa(num)
		if i == 0 {
			res = t
		} else {
			res = res + "." + t
		}

	}
	return res
}
