/*
像素扩大

01
10

0011
0011
1100
1100

*/

package main

import "fmt"

func main(){
	var n , k int
	fmt.Scan(&n, &k)
	// nums := [][]int{}

	for i:=0;i<n;i++{
		temp := []int{}
		for j:=0;j<n;j++{
			var num int
			fmt.Scan(&num)
			for t:=0;t<k;t++{
				temp = append(temp, num)
			}
		}

		for m:=0;m<k;m++{
			// nums = append(nums, temp)
			printTemp(temp)
		}
	}

}

func printTemp(temp []int){
	for i:=0;i<len(temp);i++{
		if i == 0{
			fmt.Print(temp[i])
		}else{
			fmt.Print(" ", temp[i])
		}

	}
	fmt.Println()
}
