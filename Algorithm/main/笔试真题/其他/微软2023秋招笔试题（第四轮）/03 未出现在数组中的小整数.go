/**

A := []int{1,2,4,3}
B := []int{1,3,2,3}
输出
C ：= []int{1,3,4,3}
2

3,2,1,6,5
4,2,1,3,3
输出
4,2,1,6,5
3

1,2
1,2
输出
3
*/

package main

func Solution(A []int, B []int) int {
	// write your code in Go (Go 1.4)

	hash := make(map[int]bool)
	res := 1
	n := len(A)
	for i := 0; i < n; i++ {
		num := max(A[i], B[i])
		if num == res {
			res++
			for hash[res] {
				delete(hash, res)
				res++
			}
		} else if num > res {
			hash[num] = true
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


/**


func Solution(A []int, B []int) int {
	// write your code in Go (Go 1.4)

	hash := make(map[int]bool)
	n := len(A)
	for i:=0;i<n;i++ {
		num := max(A[i], B[i])
		hash[num] = true
	}

	i := 1
	for {
		if hash[i] == false {
			return i
		}
		i++
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

 */