/**
票数最多的获胜，票数相同，名字字典序小的获胜
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strs := strings.Split(input.Text(), ",")
	// fmt.Println(strs)
	//
	hash := make(map[string]int)
	for _, str := range strs {
		if isRight(str) {
			hash[str]++
		} else {
			fmt.Println("error.0001")
			return
		}
	}

	var name string
	num := 0
	for key, value := range hash {
		if value > num {
			name = key
			num = value
		} else if value == num {
			name = comp(name, key)
		}
	}
	fmt.Println(name)
}

func isRight(s string) bool {

	for i := range s {
		if i == 0 && (s[i] < 'A' || s[i] > 'Z') {
			return false
		} else if i != 0 && (s[i] < 'a' || s[i] > 'z') {
			return false
		}
	}
	return true
}

func comp(a, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	// a短
	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return a
		} else if a[i] > b[i] {
			return b
		}
	}
	return a
}
