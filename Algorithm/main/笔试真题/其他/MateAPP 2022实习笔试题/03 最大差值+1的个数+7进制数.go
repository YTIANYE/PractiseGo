/**
题目：
第一步：
数组x
如果长度<2，返回0
对数组进行排序
如果数组长度>8, 取数组的前8位，
返回相邻数的差的最大的那一个

第二步：
如果n<0, n = 0 ,如果n > 1000, n = 1000 。
返回1...n*3范围内所有数中的每一位中,1出现的次数

第三步：
如果n > 100000, n = 100000
取n的绝对值
返回n的7进制数
*/

package main

// import "sort"
// import "math"
// import "strconv"
// 10,57,12,34,81  [2,4]
// 20 行
/**
 * 请勿修改返回值类型
 */
func answer(x []int) []int {

	// 处理结果
	ans1 := step1(x)
	ans2 := step2(ans1)
	ans3 := step3(ans2)
	// fmt.Println("ans1:", ans1)
	// fmt.Println("ans2:", ans2)
	// fmt.Println("ans3:", ans3)

	return ans3
}

//
// 第一步
func step1(x []int) int {
	if len(x) < 2 {
		return 0
	}

	// sort.Ints(x)
	var quicksort func(left, right int)
	quicksort = func(left, right int) {
		if left >= right {
			return
		}

		t := x[left]
		i, j := left, right
		for i < j {
			for i < j && x[j] >= t { //
				j--
			}
			for i < j && x[i] < t {
				i++
			}
			x[i], x[j] = x[j], x[i]
		}
		x[left], x[i] = x[i], x[left]
		quicksort(left, i-1)
		quicksort(i+1, right)
	}
	quicksort(0, len(x)-1)

	if len(x) > 8 {
		x = x[:8]
	}

	res := cha(x[1], x[0])
	for i := 1; i < len(x); i++ {
		res = max(res, cha(x[i], x[i-1]))
	}
	return res
}

func cha(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 第二步
func step2(n int) int {
	if n < 0 {
		n = 0
	}
	if n > 1000 {
		n = 1000
	}

	res := 0
	for i := 1; i <= n*3; i++ {
		res += cishu(i)
	}
	return res
}

func cishu(n int) int {

	res := 0
	for n != 0 {
		t := n % 10
		n /= 10
		if t == 1 {
			res++
		}
	}
	return res
}

// 第三步
func step3(n int) []int {
	if n > 100000 {
		n = 100000
	}
	if n < 0 {
		n = -1 * n
	}

	res := []int{}
	for n != 0 {
		num := n % 7
		res = append([]int{num}, res...)
		n = n / 7
	}
	return res
}

/*
13
6   1
3   0
1   1
0   1
*/
