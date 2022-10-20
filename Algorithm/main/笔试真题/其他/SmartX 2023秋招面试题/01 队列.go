/**
一个二维数组，数组中每个元素包含两个数，分别是“一段”的左右下标，（闭区间），
二维数组中的每个元素不重叠，且已排序
插入一个新的元素，如果元素发生重叠，请合并重叠的元素
 */
package main

/**
[1,3], [4,6], [9,12], [14,15]
[7,8]
[7,13]
*/

// 思路：对插入的元素的 2个数进行 二分查找
func insert(arrs [][]int, num []int) [][]int {
	n := len(arrs)

	// 二分查找到 小于num的所有的数中最大的那个的下标
	var find func(l, r int, t int, num int) int // t: 0 / 1
	find = func(l, r int, t int, num int) int {
		if l > r {
			return r
		}
		mid := (l + r) / 2
		if (num > arrs[mid][t] && (mid+1 < n && num < arrs[mid+1][t])) || num > arrs[mid][t] && mid == n-1 {
			return mid
		} else if num > arrs[mid+1][t] {
			return find(mid+1, r, t, num)
		} else { // <=
			return find(l, mid-1, t, num)
		}
	}
	//	删除掉重叠的元素
	var del func(int, int)
	del = func(l, r int) {
		if r < n {
			arrs = append(arrs[:l], arrs[r+1:]...)
		} else {
			arrs = arrs[:l]
		}
	}
	//
	/**
	  [1,3], [4,6], [10,12], [13,15]
	  [7,8]
	  [7,14]
	*/
	left := find(0, n-1, 0, num[0])
	right := find(0, n-1, 1, num[1])
	//
	var index int
	if arrs[left][1] >= num[0] {
		arrs[left][1] = num[1]
		index = left
	} else {
		arrs = append(arrs[:left+1], append(num, arrs[left+1]...))
		index = left + 1
	}
	// right: 9
	if right+1 < n && arrs[index][1] >= arrs[right+1][0] { //
		arrs[index][1] = arrs[right+1][1]
		del(index+1, right+1)
	} else {
		arrs[index][1] = arrs[right][0]
		del(index+1, right)
	}
	//
	return arrs

}
