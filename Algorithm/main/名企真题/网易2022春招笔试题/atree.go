/*

从1...n 的n个数，构建一棵完全二叉树，并层次遍历打印出这棵树
需要满足每个节点与其父节点的乘积为偶数
*/

package main

import "fmt"

/*
12345
			2
		1		3
	4
5
*/

// 我的题解：未完全通过样例

func main() {
	var n int
	fmt.Scan(&n)

	jique := []int{}
	ouque := []int{}

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			jique = append(jique, i)
		} else {
			ouque = append(ouque, i)
		}
	}

	fmt.Printf("%d", jique[0])
	jique = jique[1:]

	tag := false
	for i := 2; i <= n; i *= 2 {
		if tag {// jishu
			if i > len(jique) {
				printque(jique)
				printque(ouque)
				break
			} else {
				tempque := jique[:i]
				jique = jique[i:]
				printque(tempque)
			}
		} else {
			if i > len(ouque) {
				printque(ouque)
				printque(jique)
				break
			} else {
				tempque := ouque[:i]
				ouque = ouque[i:]
				printque(tempque)
			}
		}
		tag = !tag
	}

}

func printque(que []int) {
	for _, num := range que {
		fmt.Printf(" %d", num)
	}
}
