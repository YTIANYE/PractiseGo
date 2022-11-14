/**

S 走1步
H 走2步
J 走3步
从第一个出发，问能否抵达最后一个位置
首尾位置均为-，表示可以落脚
* 表示障碍，不可以落脚

返回有所少种方案，结果太大应对1000000007 取模
*/
package main

import "fmt"

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param level string字符串
 * @return long长整型
 */

/**

"----"
4

//
"-**-"
1

//
"-*-****-*-"
0

*/

func Calc1(level string) int64 {
	// write code here
	//
	var k int64
	k = 1000000007
	n := len(level)
	if level[0] == '*' || level[n-1] == '*' {
		return 0
	}
	for i := 3; i <= n; i++ {
		if level[i-3:i] == "***" {
			return 0
		}
	}
	//
	dp := make([]int64, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		if level[i] == '*' {
			continue
		}
		if i-1 >= 0 && level[i-1] == '-' {
			dp[i] = (dp[i] + dp[i-1]) % k
		}
		if i-2 >= 0 && level[i-2] == '-' {
			dp[i] = (dp[i] + dp[i-2]) % k
		}
		if i-3 >= 0 && level[i-3] == '-' {
			dp[i] = (dp[i] + dp[i-3]) % k
		}
	}
	return dp[n-1]
}

func main() {



	strs := []string{	"----", 	"-**-", 	"-*-****-*-"}
	for _, s := range strs {
		fmt.Println(Calc1(s))
	}
}
