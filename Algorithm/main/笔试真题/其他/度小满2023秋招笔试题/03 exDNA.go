/**
exDNA
时间限制： 3000MS
内存限制： 589824KB
题目描述：
       科学家在外星球发现了一种全新的生物，该种生物的遗传物质与DNA类似，被称为exDNA。

       与DNA类似，exDNA由9种碱基对依次连接而成，将这9种碱基对分别编号为1至9，则一段exDNA链可以写成一个字符串，例如：

                                 1376329476294518......

        进一步研究发现，如果exDNA中存在两个相邻的1号碱基对，则该段exDNA的表征将出现问题。因此科学家将exDNA分为两类：正常exDNA和错误exDNA。正常exDNA中不含相邻的1号碱基对，错误exDNA包含至少一处相邻的1号碱基对。形式化地，设exDNA长度为n，第i个碱基对编号为ai，则当存在一个i使ai=ai+1=1(1≤i≤n-1)时，该exDNA为错误exDNA，反之为正常exDNA。

       例如12445，2414183，1都是正常exDNA；11244，52112，67113都是错误exDNA。



       更进一步的研究发现，exDNA在复制时可能产生复制错误，表现为相邻的两个碱基对交换位置，例如12345复制后可能得到13245。由于发生错误的可能性很低，可以认为一段exDNA在复制时至多只会发生一次复制错误。在至多产生一次复制错误的条件下，如果一段exDNA复制后得到的一定是正常exDNA，则称其为安全exDNA；如果复制后可能得到错误exDNA，则称其为不安全exDNA。

       例如12445，1，123123都是安全exDNA，2414183，1214，311都是不安全exDNA。



       请编写一个程序，计算有多少种不同的长度为n的安全exDNA链，结果对1000000007取模。



输入描述
输入一行，包含一个整数n，其中1≤n≤100000，表示exDNA链的长度。

输出描述
输出一行，包含一个整数，表示长度为n的安全exDNA链的种类数。

答案对1000000007取模。
 */


package main

import (
	"fmt"
	"math"
)

/**

11

111 1x1 x11 11x

1111 x111 1x11 11x1 111x x1x1 1x1x xx11 x11x 11xx
*/

// 27%

func main() {
	var n int
	fmt.Scan(&n)
	hash := make(map[string]bool)
	hash["11"] = true
	for i := 3; i <= n; i++ {
		newHash := make(map[string]bool)
		str := ""
		for j:=0;j<i;j++ {
			str += "1"
		}
		newHash[str] = true
		for s := range hash {
			add(newHash, s)
		}
		hash = newHash
	}
	//
	res := math.Pow(9, float64(n))
	nums := make([]int, n)

	for s := range hash {
		nums[cishu(s)]+=1
		// res -= cishu(s)
	}

	for i, num := range nums {
		if num == 0 {
			continue
		}
		res -= float64(num) * math.Pow(8, float64(i))
	}
	//
	ans := int(res - (float64(int(res/1000000007)) * 1000000007.0))
	fmt.Println(ans)
}

func cishu(s string) int {
	var count int
	for i := range s {
		if s[i] == 'x' {
			count+=1.0
		}
	}
	return count
	//if count == 0 {
	//	return 1.0
	//}
	//return math.Pow(8, count)
}

func add(hash map[string]bool, s string) {
	for i := 0; i <= len(s); i++ {
		var str string
		if i == 0 {
			str = "x" + s
		} else if i == len(s) {
			str = s + "x"
		} else {
			str = s[:i] + "x" + s[i:]
		}
		hash[str] = true
	}
}
