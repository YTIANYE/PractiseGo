/**
组队吃鸡
时间限制： 3000MS
内存限制： 589824KB
题目描述：
最近一款吃鸡类型的游戏火爆全球。在组队模式下，你可以邀请自己的好友组建自己的小队，并选择是否填充（是否同意和非好友游玩），然后加入游戏。现在有A个单人队伍，B个双人队伍，C个三人队伍，D个四人队伍，并且全都同意填充，但已有的多人队伍的队员不能被拆开填充到其他队伍，请问最多能组成多少个四人队伍。



输入描述
第一行一个正整数T，表示数据组数。（1≤T≤100）

接下来T行，每行四个非负整数，A，B，C，D。（0≤A, B, C, D≤150）

输出描述
共T行，每行输出一个队伍数。


样例输入
4
1 2 3 4
4 3 2 1
2 2 2 1
0 2 0 1
样例输出
6
5
4
2
 */

package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for ;T!=0;T-- {
		var a, b, c,d int
		fmt.Scan(&a, &b,&c,&d)
		res := 0
		res += d + b/2 + min(a, c)
		numb := b % 2
		if min(a, c) == c {  // a有剩余
			numa := a - c
			if numb == 1 && numa >= 2 {
				res += 1 + (numa-2) / 4
			}else {
				res += numa / 4
			}
		}
		fmt.Println(res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}