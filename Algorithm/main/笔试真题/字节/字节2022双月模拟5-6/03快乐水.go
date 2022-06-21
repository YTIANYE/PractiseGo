/**
题目：可以摆放一排可乐或者雪碧，每相邻的5个中有2瓶是可乐。n瓶快乐水，问有多少种情况？结果 mod 10e9+7

*/

package main

import "fmt"


// 答案不完全正确
func main() {

	hash5 := make(map[string]int)
	hash4 := make(map[string]int)

	var dfs func(string)
	dfs = func(str string) {
		if len(str) == 4 {
			hash4[str]++
			return
		}
		dfs(str + "0")
		dfs(str + "1")
	}
	dfs("") // 初始化hash4

	// 去除 Hash5 中不符合要求的结果
	var deleteHash5 func()
	deleteHash5 = func() {
		deleted := []string{"00000", "10000", "01000", "00100", "00010", "00001"}
		for _, key := range deleted {
			hash5[key] = 0
		}
	}

	var generateHash4 func()
	generateHash4 = func() {
		hash4 = make(map[string]int)
		for key, num := range hash5 {
			hash4[key[:4]] = (hash4[key[:4]] + num) % 1000000007
		}
	}

	var generateHash5 func()
	generateHash5 = func() {
		hash5 = make(map[string]int)
		for key, num := range hash4 {
			hash5[key+"0"] = (hash5[key+"0"] + num) % 1000000007
			hash5[key+"1"] = (hash5[key+"1"] + num) % 1000000007
		}
	}

	var n int
	fmt.Scan(&n)
	for i := 5; i <= n; i++ {
		generateHash5()
		deleteHash5()
		generateHash4()
	}

	// 计算最终结果
	res := 0
	for _, num := range hash5 {
		res = (res + num) % 1000000007
	}
	fmt.Println(res)

}

/**
1
5

26
*/
/**
1
6

48
*/

/**
2
666
6

897443582
48
*/
