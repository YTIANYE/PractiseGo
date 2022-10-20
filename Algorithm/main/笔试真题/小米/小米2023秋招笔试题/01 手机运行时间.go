/**
手机运行时长
时间限制： 3000MS
内存限制： 589824KB
题目描述：
米小兔新买了一台小米mix fold 2，配备8英寸大屏，骁龙8+处理器和12G内存，拥有强劲性能，后台可以轻松在同一时间运行多款App。假设当后台没有App在运行时，手机进入休眠；否则手机处于运行状态。

第i个App的开始运行时间starti，运行结束时间endi，组成第i个App的运行区间intervals[i] = [starti, endi]。

数组intervals表示今天米小兔的mix fold 2的各个App的运行区间的集合，求今天米小兔的手机运行了多长时间



输入描述
输入所有App的运行时间区间，每行都是一个App运行区间starti和endi

输出描述
运行时间


样例输入
1 3
2 6
8 10
15 18
样例输出
10

 */
package main

import (
	"fmt"
	"io"
	"sort"
)

// 100%
func main() {
	nums := [][]int{}
	for {
		var a, b int
		if _, ok := fmt.Scan(&a, &b); ok != io.EOF {
			num := []int{a, b}
			nums = append(nums, num)
		} else {
			break
		}
	}
	//
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] <= nums[j][0]
	})
	//
	res  := 0
	last := -1
	for i := range nums {
		a, b := nums[i][0], nums[i][1]
		if last <= a {
			res += b-a
			last = b
		}else if last > a && last <=b {
			res += b-last
			last = b
		}
	}
	//
	fmt.Println(res )
}
