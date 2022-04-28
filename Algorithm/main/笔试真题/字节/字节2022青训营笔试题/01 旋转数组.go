/*
题目：一个有序数组，在某一位置切成两端，前后顺序颠倒后拼接
在新数组中查找一个目标值k是否存在，
数组中存在重复元素
空间复杂度O(1)
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	strs := strings.Split(input.Text(), "]")
	nums := strings.Split(strs[0][1:], ",")
	k, _ := strconv.Atoi(strs[1][1:])

	// fmt.Println(nums)
	// fmt.Println(k)

	res := false
	var biaFind func(left, right int)
	biaFind = func(left, right int) {
		if left > right {
			return
		}
		mid := (left + right) / 2
		l, _ := strconv.Atoi(nums[left])
		r, _ := strconv.Atoi(nums[right])
		m, _ := strconv.Atoi(nums[mid])
		// [2,5,6,0,0,1,2] 1
		// 这种旋转数组一般考虑 > < =
		if k > m {
			if l < r {
				biaFind(mid+1, right)
			} else if l > r {
				biaFind(left, mid-1)
			} else { // l == r
				biaFind(mid+1, right)
				biaFind(left, mid-1)
			}
		} else if k < m {
			if l < r {
				biaFind(left, mid-1)
			} else if l > r {
				biaFind(mid+1, right)
			} else { // l == r
				biaFind(mid+1, right)
				biaFind(left, mid-1)
			}
		} else { // m == k
			res = true
			return
		}
	}

	biaFind(0, len(nums)-1)
	fmt.Println(res)
}

/**
[2,5,6,0,0,1,2] 1
[2,5,6,0,0,1,2] 3

*/
