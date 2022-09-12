package main



/**
定位
时间限制： 3000MS
内存限制： 589824KB
题目描述：
小团在地图上放了三个定位装置，想依赖他们来进行定位！

小团的地图是一个n×n的一个棋盘，他在(x1,y1),(x2,y2),(x3,y3) xi,yi ∈ Z ∩ [1,n] 这三个位置分别放置了一个定位装置（两两不重叠）。然后小团在一个特定的位置(a,b)a,b ∈ Z ∩ [1,n]放置了一个信标。每个信标会告诉小团它自身到那个信标的曼哈顿距离，即对i=1,2,3 小团知道(|xi-a|+|yi-b|)，现在小团想让你帮他找出信标的位置！注意，题目保证最少有一个正确的信标位置。

因为小团不能定位装置确定出来的信标位置是否唯一，如果有多个，输出字典序最小的那个。(a,b)的字典序比(c,d)小，当且仅当 a<c或者a==c∧b<d



输入描述
第一行一个正整数n，表示棋盘大小。

第二行两个整数，分别表示x1与y1，即第一个定位器的位置。

第三行两个整数，分别表示x2与y2，即第二个定位器的位置。

第四行两个整数，分别表示x3与y3，即第三个定位器的位置。

第五行三个整数，分别表示第一、二、三个定位器到信标的曼哈顿距离。第i个定位器到信标的曼哈顿距离即(|xi-a|+|yi-b|)



数字间两两有空格隔开，对于所有数据， n≤50000,  1≤xi,yi≤n

输出描述
输出一行两个整数，表示字典序最小的可能的信标位置。


样例输入
3
2 1
2 2
2 3
2 1 2
样例输出
1 2

提示
样例解释：

与 (2, 1) 的哈曼顿距离为2的位置有三个，分别是 (1, 2), (2, 3), (3, 2)

与 (2, 2) 的哈曼顿距离为1的位置有四个，分别是 (1, 2), (2, 1), (2, 3), (3, 2)

与 (2, 3) 的哈曼顿距离为2的位置有三个，分别是 (1, 2), (2, 1), (3, 2)

所以只有 (1, 2), (3, 2) 这两个位置有可能是信标，而 (1, 2) 的字典序最小，所以输出 (1, 2)
 */
import (
	"fmt"
	"strconv"
	"strings"
)
// 通过样例 100%

func main() {
	var n int
	var x1, y1, x2, y2, x3, y3 int
	var l1, l2, l3 int
	fmt.Scan(&n)
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)
	fmt.Scan(&l1, &l2, &l3)

	//
	hash1 := manhadun(x1, y1, l1, n)
	hash2 := manhadun(x2, y2, l2, n)
	hash3 := manhadun(x3, y3, l3, n)
	fmt.Println(hash1)
	fmt.Println(hash2)
	fmt.Println(hash3)
	//
	x, y := n, n
	for str := range hash1 {
		if hash2[str] && hash3[str] {
			strs := strings.Split(str, ",")
			a, _ := strconv.Atoi(strs[0])
			b, _ := strconv.Atoi(strs[1])
			if a < x || (a == x && b < y) {
				x, y = a, b
			}
		}
	}
	fmt.Println(x, y)
}

func manhadun(x, y, l, n int) map[string]bool {

	isRight := func(x, y int) bool {
		if x >= 1 && x <= n && y >= 1 && y <= n {
			return true
		}
		return false
	}

	res := make(map[string]bool)
	for i := 0; i <= l; i++ {
		j := l - i
		if isRight(x+i, y+j) {
			res[strconv.Itoa(x+i)+","+strconv.Itoa(y+j)] = true
		}
		if isRight(x-i, y+j) {
			res[strconv.Itoa(x-i)+","+strconv.Itoa(y+j)] = true
		}
		if isRight(x+i, y-j) {
			res[strconv.Itoa(x+i)+","+strconv.Itoa(y-j)] = true
		}
		if isRight(x-i, y-j) {
			res[strconv.Itoa(x-i)+","+strconv.Itoa(y-j)] = true
		}
	}
	return res
}
