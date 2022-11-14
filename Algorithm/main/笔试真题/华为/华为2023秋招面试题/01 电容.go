/**
电容对于终端来说是原材料当中的易耗品，现把仓库中的电容分成 N 份，第 i 份有 parts[i] 个电容。假如：

1.   生产线要消耗完一份电容，才能开始消耗下一份；

2.   生产线每分钟可消耗 K  个；如果某一份的电容数量小于 K 个，则生产线在这一分钟内消耗完这份电容后，剩余时间就会空闲，直到下一分钟开始消耗下一份。

距离下批电容上线前有 lengthTime 分钟，请计算在下批电容上线前，消耗完当前所有电容的最小速度 K（个/分钟，K 为整数）。

示例 ：

输入: parts = [3,6,7,11], lengthTime = 8

输出: 4

解释: 假如速度是 4 个/分钟的话：
第 1 份 parts[0] 消耗 1 分钟，
第 2 份 parts[1] 消耗 2 分钟，
第 3 份 parts[2] 消耗 2 分钟，
第 4 份 parts[3] 消耗 3 分钟，
合计 8 分钟，符合要求。
如果速度是 3 个/分钟的话，则需要 10 分钟，超过时间（8）要求；
如果是 5 个/分钟的话，也是在 8 分钟内消耗完、满足时间要求，但是比 4 个/分钟速度快，不是最小速度。



限制：

1 <= parts.length <= 10^4

parts.length <= lengthTime <= 10^9

1 <= parts[i] <= 10^9
*/
package main

import (
	"fmt"
	"math"
)

func xiaohao(parts []int, T int) int {
	res := 0
	c := math.MaxInt
	var fun func(i int) int
	fun = func(i int) int {
		count := 0
		for _, part := range parts {
			count += part / i
			if part%i != 0 {
				count++
			}
		}
		return count
	}

	for c > T {
		res++
		c = fun(res)
	}

	return res
}

func main() {
	//parts := []int{3,6,7,11}
	//T := 8
	//res := xiaohao(parts, T) // 4

	parts := []int{3,6,7,11,3}
	T := 9
	res := xiaohao(parts, T) // 4

	fmt.Println(res)
}
