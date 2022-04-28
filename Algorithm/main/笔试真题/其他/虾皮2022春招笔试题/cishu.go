/**
题目：
将 sourceX, sourceY 转换成 targetX, targetY 最少需要多少次计算
如：
1,2,7,14
4次
1-2-3-6-7
*/

package 虾皮2022春招笔试题

import (
	"math"
)

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 * 将 sourceX, sourceY 转换成 targetX, targetY 最少需要多少次计算
 * @param sourceX long长整型  x初始值
 * @param sourceY long长整型  y初始值
 * @param targetX long长整型  x目标值
 * @param targetY long长整型  y目标值
 * @return long长整型
 */

// 我的题解：通过全部测试样例
func GetMinCalculateCount(sourceX int64, sourceY int64, targetX int64, targetY int64) int64 {
	// write code here

	if sourceX > targetX || sourceY > targetY {
		return -1
	}

	lenx, leny := targetX+1, targetY+1
	dpx := make([]int64, lenx)
	dpy := make([]int64, leny)
	var i int64
	for i = 0; i < lenx; i++ {
		dpx[i] = math.MaxInt64
	}
	for i = 0; i < leny; i++ {
		dpy[i] = math.MaxInt64
	}
	x, y := sourceX, sourceY
	dpx[x], dpy[y] = 0, 0

	// 计算：动态规划
	for x < lenx || y < leny {
		if x+1 < lenx {
			dpx[x+1] = min(dpx[x+1], dpx[x]+1)
		}
		if y+1 < leny {
			dpy[y+1] = min(dpy[y+1], dpy[y]+1)
		}
		if x*2 < lenx {
			dpx[x*2] = min(dpx[x*2], dpx[x]+1)
		}
		if y*2 < leny {
			dpy[y*2] = min(dpy[y*2], dpy[y]+1)
		}
		x++
		y++
	}

	// 结果
	if dpx[targetX] == dpy[targetY] {
		return dpx[targetX]
	} else {
		return -1
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

//
// func main(){
// 	fmt.Println(GetMinCalculateCount(1,2,7,14))
// }
