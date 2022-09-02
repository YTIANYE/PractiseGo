/**
检测一个html中的所有标签<string></string> 是不是成对出现？
 */
/**
<te></te><st></st>
YES

test<tist>/<>
NO

<></>
NO
 */
package main

import "fmt"

// 100%
func main() {
	var s string
	fmt.Scan(&s)

	if isRight(s ) {
		fmt.Println("YES")
	}else {
		fmt.Println("NO")
	}
}

func isRight(s string) bool  {
	queue := []string{}
	n := len(s)
	//
	var findBiao func(start int)  (int, int)
	findBiao = func(start int)  (int, int) {
		left, right := -1, -1
		findRight := false
		for i:= start; i<n;i++ {
			if findRight == false && s[i] == '<' {
				left = i
				findRight = true
			}else if findRight == true && s[i] == '>'{
				right = i
				return left, right
			}
		}
		return left, right
	}
	//
	i:=0
	for i < n {

		left, right := findBiao(i)
		i = right+1
		//
		if left == -1 && right == -1 {
			return true
		}else if left != -1 && right == -1 {
			return false
		}
		//
		if s[left+1] == '/' {
			var s1, s2 string
			// 没有左标签
			if len(queue) == 0 {
				return false
			}
			s1 = queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if left+2 == right {
				s2 = ""
			}else {
				s2 = s[left+2:right]
			}
			// 标签不对应
			if !isSame(s1, s2){
				return false
			}
		}else if left + 1 == right {
			return false
		}else {
			var s1 string
			s1 = s[left+1:right]
			queue = append(queue, s1)
		}
	}
	// 是否有多余左标签
	if len(queue) != 0 {
		return false
	}
	return true
}



func isSame(a, b string) bool {
	if len(a) != len(b) || len(a) == 0 || len(b) == 0 {
		return false
	}
	for i:=0;i<len(a);i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}