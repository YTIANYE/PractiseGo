/**
题目：
特征一共有9中，分别是 1 2 3 4 5 6 7 8 9

输入n，表示接下来的行数
第一行 输入要验证的特征
接下来的每一行中的第一个数是特征，接下来的数是它依赖的特征

判断第一行给定的特征的依赖关系中，是否存在循环依赖，存在输出0，不存在输出1
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
/**
3
1
1 2
2 3 1

0

//
3
1 2
1 2 3
2 4

1 1

//

3
1 2
1 2 3
2 4 1

0 0
//

3
1 4
1 2 3
2 4 1

0 1
 */

// 100%
func main() {
	var n int
	fmt.Scan(&n)
	n--
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := strings.Split(input.Text(), " ") // 特征
	hash := make(map[int][]int)
	tezheng := []int{}
	for _, str := range s {
		num, _ := strconv.Atoi(str)
		tezheng = append(tezheng, num)
	}

	for ; n != 0; n-- {
		input.Scan()
		strs := strings.Split(input.Text(), " ")
		nums := []int{}
		for _, str := range strs {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
		v := make([]int, len(nums)-1)
		copy(v, nums[1:])
		hash[nums[0]] = v
	}
	//
	var tag int
	var huan func(int) bool
	huan = func(num int) bool {
		if num == tag {
			return true
		}
		for _, start := range hash[num] {
			if huan(start) {
				return true
			}
		}
		return false
	}

	//
	res := []int{}

	for _, num := range tezheng {
		tag = num
		have := false
		for _, start := range hash[num] {
			if huan(start) {
				have = true
				break
			}
		}
		if have { // 有环
			res = append(res, 0)
		} else {
			res = append(res, 1)
		}
	}


	//
	fmt.Printf("%d", res[0])
	for i:=1;i<len(res);i++ {
		fmt.Printf(" %d", res[i])
	}
}
