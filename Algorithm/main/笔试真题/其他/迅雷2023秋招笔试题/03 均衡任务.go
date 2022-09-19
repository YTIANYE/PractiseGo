/**
一个数组若干个数，
把每个数调整成一样的，最少操作多少次？

每次操作是可以把nums[i] 中的数 -temp,其相邻数中的一个 +temp
 */

/**
9 8 17  6
3
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
// 100%
func main() {
	imput := bufio.NewScanner(os.Stdin)
	imput.Scan()
	strs := strings.Split(imput.Text(), " ")

	nums :=[]int{}
	sum := 0
	for i := range strs {
		if strs[i] == "" {
			continue
		}
		num, _ := strconv.Atoi(strs[i])
		nums = append(nums, num)
		sum += num
	}
	n := len(nums)
	//
	ave := sum / n
	res := 0
	for i:=0;i<n-1;i++ {
		if nums[i] > ave {
			res++
			nums[i+1] += nums[i] - ave
		} else if nums[i] < ave {
			res++
			nums[i+1] -= ave - nums[i]
		}

	}
	fmt.Println(res)
}
