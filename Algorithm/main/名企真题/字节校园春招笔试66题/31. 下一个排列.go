/*
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。

 

示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]
 

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
123
132
213
231
312
321
*/

// 我实现的官方题解：双指针
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了79.70%的用户
*/

func nextPermutation(nums []int) []int {
	n := len(nums)

	p := n - 1
	pre := p - 1
	for pre >= 0 && nums[pre] >= nums[p] {
		pre--
		p--
	}

	if pre >= 0 {
		q := n - 1
		for q >= 0 && nums[q] <= nums[pre] {
			q--
		}
		nums[q], nums[pre] = nums[pre], nums[q]
	}
	sort.Ints(nums[pre+1:])

	fmt.Println("结果是：", nums)
	return nums
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		if input.Scan() { // 1 2 3 4
			strs := strings.Split(input.Text(), " ")
			nums := []int{}
			for i := range strs {
				num, _ := strconv.Atoi(strs[i])
				nums = append(nums, num)
			}
			for i := 0; i < 24; i++ {
				nums = nextPermutation(nums)
			}
		} else {
			break
		}

	}

}
