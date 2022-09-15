/**
'b'可以分裂为‘aa'
'c'可以分裂为’bb‘
以此类推

问，一个长度为n的都是’a‘构成的字符串，最少是由多长的字符串分裂而成的，输出一种结果即可
*/
package main

import "fmt"

//100%
func main () {
	var n int
	fmt.Scan(&n)
	fmt.Println(fun1(n))

}


func fun1(n int) string {
	res := ""
	nums := []int{}


	for i:=1;i<=n;i*=2 {
		nums = append(nums, i)

	}
	//
	i:= len(nums)-1
	for  n != 0  {
		for nums[i] > n {
			i--
		}
		res += string('a' + byte(i))
		n -= nums[i]
	}
	return res
}