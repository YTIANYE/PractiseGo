/**
将一个数分成若干个数相加，使得这些数的乘积最大，输出乘积结果的分子和分母
例子
7
可以分成三个 7/3
4
可以分成两个2
 */
package main

import (
	"fmt"
	"math"
)
// 85%
func main() {
	var n int
	fmt.Scan(&n)
	fenzi, fenmu := 1,1
	for i:=2;i<=n;i++ {
		a :=  result(fenzi, fenmu)
		b := result(n, i)
		if b > a {
			fenzi = n
			fenmu = i
		}
	}
	//
	m := fenmu
	fenzi, fenmu = yue(fenzi, fenmu)
	fenzi = int(math.Pow(float64(fenzi), float64(m)))
	fenmu = int(math.Pow(float64(fenmu), float64(m)))
	fmt.Println(fenzi, fenmu)
}

func result(fenzi, fenmu int) float64 {
	return math.Pow(float64(fenzi), float64(fenmu)) / math.Pow(float64(fenmu), float64(fenmu))
}


func yue(fenzi int, fenmu int) (int, int) {
	for i:= 2;i<=fenmu;i++ {
		if fenzi % i == 0 && fenmu % i == 0 {
			fenzi /= i
			fenmu /= i
			i = 2
		}
	}
	return fenzi, fenmu
}