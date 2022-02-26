package main

import "math"

func printNumbers(n int) []int {
	res := []int{}
	end := int(math.Pow(10, float64(n)))
	for i:=1;i<end;i++{
		res = append(res, i)
	}
	return res
}

