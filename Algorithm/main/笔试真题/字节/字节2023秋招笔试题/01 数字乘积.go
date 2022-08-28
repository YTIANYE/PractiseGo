/**
题目：
一个数组，数组中的元素只包括[0,1,2,4,8,16,32,64,128,256,512,1024]
找出乘积最大的子数组的左右坐标，如果存在乘积相同的，取做坐标小的，如果左坐标相同，取右坐标小的
 */


package main

import "fmt"
/**
2
5
1 2 4 0 8
7
1 2 4 8 0 256 0

1 3
6 6
*/

// 0%
func main() {
	var t int
	fmt.Scan(&t)
	hash := make(map[int]int)
	temp := 1
	for i := 0; i <= 10; i++ {
		hash[temp] = i
		temp *= 2
	}
	//
	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		nums := make([]int, n)
		index := []int{}
		index = append(index, -1)
		for i := 0; i < n; i++ {
			var a int
			fmt.Scan(&a)
			if a == 0 {
				index = append(index, i)
				nums[i] = -1
			}else {
				nums[i] = hash[a]
			}
		}
		index = append(index, n)

		//
		m := nums[0]
		left, right := 0, 0
		for i := 1; i < len(index); i++ {
			l := index[i-1] + 1
			r := index[i] - 1
			sum := 0
			for j := l; j <= r; j++ {
				sum += nums[j]
			}
			//
			for r > l && nums[r] == 0 {
				r--
			}
			//
			if sum > m {
				left = l
				right = r
				m = sum
			} else if sum == m {
				if l < left {
					left = l
					right = r
					m = sum
				} else if l == left {
					for r > l && nums[r] == 0 {
						r--
					}
					tright := right
					for right > left && nums[right] == 0 {
						right--
					}
					if r < right {
						left = l
						right = r
						m = sum
					} else {
						right = tright
					}
				}
			}
		}
		//
		fmt.Println(left+1, right+1)
	}
}
