/**
打印一个‘U’
长宽为4n

n = 1
*..*
*..*
*..*
.**.

n = 2
**....**
**....**
**....**
**....**
**....**
**....**
.**..**.
..****..


 */

package main

import "fmt"

// 通过样例 100%
func main(){
	var n int
	fmt.Scan(&n)
	s1 := ""
	for i:=0;i<n;i++{
		s1 += "*"
	}
	for i:=0;i<n *2;i++{
		s1 += "."
	}
	for i:=0;i<n;i++{
		s1 += "*"
	}
	//
	for i:=0;i<n*4-n;i++{
		fmt.Println(s1)
	}

	//
	for i:=0;i<n;i++{
		s1 = "." + s1[:2*n-1] + s1[2*n+1:] + "."
		fmt.Println(s1)
	}



}
