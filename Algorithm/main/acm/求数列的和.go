/*
时间限制： 4000MS
内存限制： 557056KB

题目描述：
数列的定义如下： 数列的第一项为n，以后各项为前一项的平方根，求数列的前m项的和。

输入描述
输入数据有多组，每组占一行，由两个整数n（n<10000）和m(m<1000)组成，n和m的含义如前所述。

输出描述
对于每组输入数据，输出该数列的和，每个测试实例占一行，要求精度保留2位小数。

样例输入
81 4
2 2

样例输出
94.73
3.41
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	func3()

}

func func1() {
	var n, m int
	for { // 循环获取输入
		if _, err := fmt.Scan(&n, &m); err != io.EOF {

			temp := float64(n)
			res := temp
			for ; m != 1; m-- {
				temp = math.Pow(temp, 0.5)
				res += temp
			}
			fmt.Printf("%.2f\n", res)
		} else {
			break
		}
	}
}

func func2() {

	for {
		input := bufio.NewScanner(os.Stdin)
		if input.Scan() {
			strs := strings.Split(input.Text(), " ")
			n, _ := strconv.Atoi(strs[0])
			m, _ := strconv.Atoi(strs[1])

			temp := float64(n)
			res := temp
			for ; m != 1; m-- {
				temp = math.Pow(temp, 0.5)
				res += temp
			}
			fmt.Printf("%.2f\n", res)
		} else {
			break
		}

	}
}

func func3() {

	// input := bufio.NewScanner(os.Stdin)
	// input.Scan()
	// strs := strings.Split(input.Text(), " ")
	// fmt.Println(strs)

	var a, b int
	for {
		if _, err := fmt.Scan(&a, &b); err != io.EOF {
			fmt.Println(a + b)
		}else{
			break
		}
	}

}
