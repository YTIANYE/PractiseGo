/**
有n个数，求出X和Y，使得每个数乘以X或Y的结果之和为0
并给出每个数乘以的是X还是Y
 */
package main

import "fmt"
/**
例子1 ：
2
1 1

输出
1 -1
XY

例子2：
5
248 537 207 611 711

1148 -1226

*/

// 没做完
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	sum := 0
	for i:= range nums {
		fmt.Scan(&nums[i])
		sum += nums[i]
	}

	//
	x := sum / 2
	y := -(sum - x)

	//

	// 递归所有XY的情况
	sX := ""
	sY := ""
	for i:=0;i<n;i++ {
		sX += "X"
		sY += "Y"
	}
	strs := []string{}
	s := ""

	var allXY func()
	allXY = func()  {
		if len(s) == n {
			if s == sX || s == sY {
				return
			}
			temp := ""
			for i:= range s {
				temp += string(s[i])
			}
			strs = append(strs, temp)
			return
		}
		s += "X"
		allXY( )
		s = s[:len(s)-1]
		s += "Y"
		allXY()
		s = s[:len(s)-1]
	}

	// 计算结果是不是 0
	jisuan := func() bool {
		var res int64
		for _, str := range strs {
			for i := range str {
				if str[i] == 'X' {
					res += int64(x * nums[i])
				}else {
					res += int64(y * nums[i])
				}
			}
			//
			if res == 0 {
				s = str
				return true
			}
		}
		return false
	}

	// 判断x,y   -x,-y是不是符合
	isRight := func() bool {
		if jisuan() {
			return true
		}
		x,y = -x, -y
		if jisuan() {
			return true
		}
		x,y = -x, -y
		return false
	}

	//
	allXY()
	for !isRight() && x <= sum{
		x++
		y++
	}
	////
	//if x > sum {
	//	fmt.Println(-1)
	//}else {
	//	fmt.Println(x, y)
	//	fmt.Println(s)
	//}
	fmt.Println(x, y)
	fmt.Println(s)

}

