package main

import (
	"fmt"
	"strconv"
)

func he(num1, num2 string) string {
	jinwei := 0
	n, m := len(num1), len(num2)
	res := ""
	i := n - 1
	j := m - 1
	for i>=0 && j>=0 {
		a,_ := strconv.Atoi(string(num1[i]))
		b, _ := strconv.Atoi(string(num2[j]))
		sum := jinwei + a + b
		c := sum % 10
		res = strconv.Itoa(c) + res
		jinwei = sum /10
		i--
		j--
	}
	for i>=0 {
		a,_ := strconv.Atoi(string(num1[i]))
		sum := jinwei + a
		c := sum % 10
		res = strconv.Itoa(c) + res
		jinwei = sum /10
		i--
	}
	for j>=0 {
		b, _ := strconv.Atoi(string(num2[j]))
		sum := jinwei + b
		c := sum % 10
		res = strconv.Itoa(c) + res
		jinwei = sum /10
		j--
	}
	if jinwei != 0 {
		res = strconv.Itoa(jinwei) + res
	}
	return res
}

func main() {
	fmt.Println(he("111", "999"))
}
