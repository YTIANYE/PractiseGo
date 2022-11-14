/**
题目：
m个面试官n个面试者，每个面试官最多面试x次，每个面试者需要面试两次
依次输入面试官会的语言，以及面试者面试的语言
问：能否安排面试，能的话输出面试match数组，不能输出false


测试样例
4 6 4
python c++ java
python java
go java
c++ go
java
python
go
c++
java
python

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var m, n, x int
	fmt.Scan(&m, &n, &x)
	guan := make(map[int][]string, m)
	guanx := make(map[int]int) // 剩余x
	guanl := [][]int{}         // 面试官会的语言种类数 和 面试官序号
	zhe := make(map[int]string, n)
	zhex := make(map[int]int) // 剩余次数
	match := [][]int{}

	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < m; i++ {
		input.Scan()
		strs := strings.Split(input.Text(), " ")
		for j := range strs {
			// fmt.Println("语言：", strs[j])
			guan[i] = append(guan[i], strs[j])
		}
		guanx[i] = x
		guanl = append(guanl, []int{len(strs), i})

		temp := make([]int, n)
		match = append(match, temp)
	}
	for i := 0; i < n; i++ {
		input.Scan()
		strs := strings.Split(input.Text(), " ")
		zhe[i] = strs[0]
		zhex[i] = 2
	}
	// 语言种类会的少的排在前面
	sort.Slice(guanl, func(i, j int) bool {
		return guanl[i][0] < guanl[j][0]
	})

	// 打印
	// fmt.Println(zhe)
	// fmt.Println(zhex)
	// fmt.Println(guan)
	// fmt.Println(guanx)
	// fmt.Println(guanl)

	// 计算
	end := false                   // 不用结束
	for zheid, lang := range zhe { // 对于每个面试者
		for zhex[zheid] != 0 { // 每个面试者两次面试
			find := false
			for _, g := range guanl { // 对于每个面试官
				guanid := g[1]
				if inLang(lang, guan[guanid]) && guanx[guanid] != 0 && match[guanid][zheid] != 1 {
					match[guanid][zheid] = 1
					zhex[zheid]--
					guanx[guanid]--
					find = true
					break
				}
			}
			// 每个面试官都不能满足了
			if !find {
				end = true
			}
		}

		// 若当前面试者不能满足，结束面试
		if end {
			break
		}
	}
	// 输出结果
	if end { // 不能匹配而结束
		fmt.Println("false")
		// printmatch(match)
	} else {
		printmatch(match)
	}
}

// 面试官是否可以面该语言
func inLang(lang string, langs []string) bool {
	for _, val := range langs {
		if val == lang {
			return true
		}
	}
	return false
}

func printmatch(match [][]int) {
	for i := 0; i < len(match); i++ {
		for j := 0; j < len(match[i]); j++ {
			num := match[i][j]
			if j == 0 {
				fmt.Printf("%d", num)
			} else {
				fmt.Printf(" %d", num)
			}
		}
		fmt.Printf("\n")
	}
}
