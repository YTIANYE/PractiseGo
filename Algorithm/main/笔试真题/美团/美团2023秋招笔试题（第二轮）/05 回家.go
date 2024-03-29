/**
回家
时间限制： 3000MS
内存限制： 589824KB
题目描述：
小美在回家路上碰见很多缠人的可爱猫猫！因为猫猫太可爱了以及小美十分有爱心，每遇到一只猫猫，小美忍不住停下来花费T的时间抚摸猫猫让猫猫不再缠着小美。而一路上小美能捡到很多亮闪闪的小玩具，这里我们给每个小玩具的种类都编了号，从1~k，一共k种小玩具，对于每个所属种类i的小玩具，小美可以选择将它送给遇到的一只猫猫玩，这样的话可以只花费ti的时间就可以让这只猫猫心满意足的离开。小美想知道，在她以最佳的对小玩具的用法下，她最少耗费多少时间在打发猫猫（即只考虑摸猫时间以及用小玩具打发猫的时间）。

注意，每个捡到的小玩具只能用一次。



输入描述
第一行三个正整数n、k和T，分别表示小美回家遇见的事件数、小玩具种类总数以及摸猫时间！

接下来一行k个数t1,t2, ...,tk, 含义如题所述，为每种小玩具打发猫猫所用时间。

接下来一行n个数e1,e2, ...,en ，表示n次事件，对第i次事件，如果ei=0，则表示遇到了一只猫猫，小美可以选择花费T的时间去抚摸，或者用一个小玩具送给猫猫来打发它(如果小美有的话)。如果ei>0，则表示小美在这里捡到了一个小玩具，种类为ei。初始时候小美身上没有任何小玩具，她可以携带任意多个小玩具。

对于所有数据，1≤n≤50000，1≤k≤50000，0≤ei≤k, 1≤T,ti≤104

输出描述
输出最少耗费多少时间在打发猫猫（即只考虑摸猫时间以及用小玩具打发猫的时间）。


样例输入
6 2 100
1 50
0 1 2 0 1 0
样例输出
102

提示
样例解释

一开始没有小玩具，遇到一只猫猫只能抚摸，花费了100的时间。

接下来获得了小玩具1和2，然后遇到一只猫猫，用了小玩具1，只花费了1的时间。

接下来又获得一个小玩具1之后又遇见一只猫猫，因为又有小玩具1了，所以还是只用花费1的时间。

总共用时102
 */


package main

import (
	"fmt"
	"sort"
)

// 没写完
func main() {
	var n, k , T int
	fmt.Scan(&n, &k, &T)
	kv := make([]int, k+1) // 耗时
	ks := make(map[int]int)// key 耗时 玩具数目
	for i:=1;i<=k;i++ {
		var a int
		fmt.Scan(&a)
		kv[i] = a
		ks[a] = 0
	}
	sort.Ints(kv)

	var res = 0
	for i:=0;i<n;i++ {
		var a int
		fmt.Scan(&a)
		if a == 0 {
			for i:=k;i>0;i-- {
				jianshi := kv[i]
				if ks[jianshi] >0 {
					ks[jianshi]--
					var haoshi int
					if jianshi > T {
						haoshi = 0
					}else {
						haoshi = T - jianshi
					}
					res += haoshi
				}
			}
		}else {
			ks[a]++
		}
	}
	fmt.Println(res )
}