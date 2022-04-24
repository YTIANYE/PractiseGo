/*
时间限制： 4000MS
内存限制： 557056KB

题目描述：
春天是鲜花的季节，水仙花就是其中最迷人的代表，数学上有个水仙花数，
他是这样定义的： “水仙花数”是指一个三位数，它的各位数字的立方和等于其本身，比如：153=1^3+5^3+3^3。
现在要求输出所有在m和n范围内的水仙花数。

输入描述
输入数据有多组，每组占一行，包括两个整数m和n（100<=m<=n<=999）。

输出描述
对于每个测试实例，要求输出所有在给定范围内的水仙花数，就是说，
输出的水仙花数必须大于等于m,并且小于等于n，如果有多个，则要求从小到大排列在一行内输出，之间用一个空格隔开;
如果给定的范围内不存在水仙花数，则输出no; 每个测试实例的输出占一行。

样例输入
100 120
300 380

样例输出
no
370 371

*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main(){
	// for {
	// 	var m, n int
	// 	if _, err := fmt.Scan(&m, &n); err != io.EOF{
	// 		shuixianhua(m, n)
	// 	}else{
	// 		break // 不要忘记 break
	// 	}
	// }

	input := bufio.NewScanner(os.Stdin)
	for {
		if input.Scan(){
			strs := strings.Split(input.Text(), " ")
			m, _:= strconv.Atoi(strs[0])
			n, _ := strconv.Atoi(strs[1])
			shuixianhua(m, n )
		}else{
			break
		}

	}
}

func shuixianhua(m, n int ){
	hua := false
	for i:=m;i<=n;i++{
		if isHua(i){
			fmt.Print(i, " ")
			hua = true
		}
	}
	if hua{
		fmt.Printf("\n")
	}else{
		fmt.Println("no")
	}

}

func isHua(num int) bool{
	temp := num
	c := float64(temp % 10)
	temp /= 10
	b := float64(temp % 10)
	temp /= 10
	a := float64(temp)

	return float64(num) == math.Pow(a, 3) + math.Pow(b, 3) + math.Pow(c, 3)
}