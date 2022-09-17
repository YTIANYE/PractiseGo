/**
w个格子，n中颜色，每个颜色i必须连着涂ai个，i必须在i-1的颜色的后面，且每种颜色之间至少一个空格不涂颜色
问，在所有涂色结果中，保持一直涂上了颜色且颜色不改变的格子有几个，在什么位置（1...n）
 */
package main

import "fmt"
// 100%
func main() {
	var w, n int
	fmt.Scan(&w, &n)
	a := make([]int, n)
	for i:=0;i<n;i++ {
		fmt.Scan(&a[i])
	}
	//
	color := []int{}//每种颜色之间间隔1个空格时的涂色方法
	local := []int{}
	for i:=0;i<n;i++ {
		if i != 0 {
			color = append(color, -1)// -1 white
			local = append(local, len(color)-1)// 记录white的位置
		}
		for j:=0;j<a[i];j++ {
			color = append(color, i)
		}

	}
	//
	res := []int{}
	num := w - len(color)// 剩余num个白色格子可以插空
	if num == 0 {
		for i:= range color {
			if color[i] == -1 {
				continue
			}
			res = append(res, i)
		}
	}else {
		index := 0

		for i:=0;i<n;i++ {
			if a[i] <= num {
				index += a[i] + 1
			} else {
				index += num
				for index < len(color) && color[index] != -1 {
					temp := index
					res = append(res, temp)
					index++
				}
				index += 1
			}
		}
	}

	//
	fmt.Println(len(res))
	for i:= range res {
		if i == 0 {
			fmt.Printf("%d", res[0]+1)
		}else {
			fmt.Printf(" %d", res[i]+1)
		}
	}
}

/**
i=2, a[i]=3 num = 2
因为a[i]>num，所以无论2 2 2前后怎么移动（或白色格子怎么插空）得到的的涂色方法中，都能保证最后一个会涂上颜色，且颜色不变
1 0 2 2 2 0 0
1 0 0 2 2 2 0
1 0 0 0 2 2 2
0 0 1 0 2 2 2

1 0 2 2 0 0
0 0 1 0 2 2

1 0 2 2 2 0
 */
