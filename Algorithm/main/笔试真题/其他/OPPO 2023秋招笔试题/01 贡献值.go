/**

若干个应用，每个引用有工作时间段，1-2， 40-80 ...
某一时间段2-4，应用1在工作，则贡献2
若某一时间段 a 2-4  b 1-8
则a贡献 2/2, b贡献 1+2/2+4
重叠部分需要平分给各应用

求每个应用的贡献值
 */

package main

import (
	"math"
	"sort"
)

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 输入参数是每个应用持有WakeLock的起始时间及结束时间，
函数返回值是每个应用对总WakeLock时长的贡献值，
返回值vector里面的每个元素的数值累加即为手机总的持有WakeLock的时长。
 * @param wakeLock int整型二维数组 vector<vector<int>>的每个元素代表某个应用的持有WakeLock的信息，
内部的vector<int>，元素按照奇偶排列，偶数元素代表拿WakeLock的起始时间点，奇数代表释放WakeLock的时间点
 * @return int整型一维数组
*/
// 80%
func getWakeLockContrib(wakeLock [][]int) []int {
	// write code here
	end := 0
	begin := math.MaxInt32
	for i := range wakeLock {
		sort.Ints(wakeLock[i])
		begin = min(begin, wakeLock[i][0])
		end = max(end, wakeLock[i][len(wakeLock[i])-1])
	}
	nums := make([]float64, len(wakeLock))
	//
	for index := begin; index <= end; index++ {
		sum := 0
		temp := []int{}

		for i := range wakeLock {
			if len(wakeLock[i]) == 0 {
				continue
			}
			if index >= wakeLock[i][0] && index < wakeLock[i][1] {
				sum++
				temp = append(temp, i)
			}
			if index == wakeLock[i][1] {
				if len(wakeLock[i]) == 2 {
					wakeLock[i] = []int{}
				} else {
					wakeLock[i] = wakeLock[i][2:]
				}
			}
		}
		//
		if sum == 0 {
			continue
		}
		for _, j := range temp {
			nums[j] += 1.0 / float64(sum)
		}
	}

	//
	res := []int{}
	for i := range nums {
		var num int
		if nums[i] - float64(int(nums[i])) > 0.99 {
			num = int(nums[i]) + 1
		}else {
			num = int(nums[i])
		}
		res = append(res, num)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
