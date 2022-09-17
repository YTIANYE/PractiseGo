/**
实验数据
时间限制： 3000MS
内存限制： 589824KB
题目描述：
小昱做了很久的实验得到了一个用正整数表示的实验数据，并记录在了纸上。但是由于做完实验太过激动，他一不小心把墨水打翻溅在了纸上，导致数据中一些位置上的数字看不清楚。他仍记得这个数据有以下三个特征：

1. 这个数是正整数，且没有前导零（即数的最高位不是0）

2. 这个数任意两个相邻数位的数字不同

3. 这个数可以被3整除

他现在很关心在满足以上特征的条件下，这个数字最小为多少。



输入描述
输入一个由数字0-9和‘?’构成的字符串。若输入的第i个字符为问号，则表示数据从高位往低位数的第i位被墨水遮盖，不能确定是哪个数字；否则，表示这一位未被墨水遮盖，是一个确定的数。

输出描述
输出一个正整数，表示实验数据最小可能是多少。


样例输入
?12?0?9??
样例输出
212101902

提示
输入的字符串长度不超过100000，且至少为1。

所有数据均保证合法，保证存在合法解，且至少含有1个‘?’。

最高位为‘?’表示最高位被遮挡无法确定。因为第一条特征限制，最高位不能为0。因为第二位为1，根据第二条限制，最高位不能为1。所以最高位只能是2到9中的任意一个数字，当最高位为2时，实验数据会最小。第4、6、8、9位数字也被墨水遮挡，当第4、6位为数字1，第8位为数字为0，第9位为数字2时满足小昱记忆中数据的特征，且是可能出现的最小值。
 */
package main

import (
	"fmt"
	"strconv"
)

// 36%
func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	res := ""
	//
	var last int
	for i := n - 1; i >= 0; i++ {
		if s[i] == '?' {
			last = i
			break
		}
	}
	//
	for i := 0; i < n; i++ {
		if s[i] != '?' {
			res += string(s[i])
			continue
		}
		// == '?'
		for j := 0; j <= 9; j++ {
			if i == 0 && j == 0 {
				continue
			}
			num := byte(j) + '0'
			if (i-1 >= 0 && num == s[i-1]) || (i+1 < n && num == s[i+1]) {
				continue
			}
			res += string(num)
			break
		}
	}
	//

	sumNum := sum(res[:last])
	if last != n-1 {
		sumNum += sum(res[last+1:])
	}
	yushu := sumNum % 3
	temp := 0
	if yushu > 0 {
		temp = 3 -yushu
	}
	for j:= temp;j<=9;j+=3 {
		if last == 0 && j == 0 {
			continue
		}
		num := byte(j) + '0'
		if (last-1 >= 0 && num == s[last-1]) || (last+1 < n && num == s[last+1]) {
			continue
		}
		//
		if last == 0 {
			res = string(num) + res[1:]
		}else if last == n-1 {
			res = res[:n-1] + string(num)
		}else {
			res = res[:last] + string(num) + res[last+1:]
		}
		break
	}
	fmt.Println(res)
}

func sum(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	for i := range s {
		num, _ := strconv.Atoi(string(s[i]))
		res += num
	}
	return res
}
