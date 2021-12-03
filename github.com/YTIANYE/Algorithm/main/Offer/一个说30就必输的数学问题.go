package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 游戏规则： 从1开始到30结束，两个人a和b轮流说一个数或者两个数，谁先开始都可以，最后一个说出30的人就输了。

// 分析：
// 【1】
// 如果某一次a说的是26或者25 26时，
// 那么b：27， a：28 29；或者b:27 28， a：29；最终结果都是b输，
// 则a一定说出26就可以赢了
//
// 【2】
// 为了保证a说25 26 或者 26，应该使b说23 24， 24， 24 25， 25，
// 如果a说了23， b只能说24， 24 25， 就可以保证a说25 26 或者 26了
// 则a一定说出23就可以赢了
//
// 【3】
// 以此类推，a只要说出以下的数，并以该数做自己说的数的结尾，就可以赢了
// [29, 26, 23, 20, 17, 14, 11, 8, 5, 2] 称之为 必赢数组 [2^n-1]
// 说了上面的数之后，接下来需要做的就是对方说一个数，我说两个数，对方说两个数，我说一个数，
// 就可以以3为一轮次持续踩中上面的数，并且一直赢到最后。
//
// 如 a:14
// 若 b:15, a:16 17
// 若 b:15 16, a:17
//
// 【4】
// 如果a先说，就说1 2, a就可以赢了
// 如果b先说，
// 若b: 1,  a: 2， a就可以赢了
// 若b:1 2,
// 此时a说数字，之后如果b不遵守与a说的数字的数目不同的原则，a就可以踩中必赢数组中的数字并取得胜利。
// 换句话说，之后a说了数字之后，b若与a说的数字个数相同，b就输了。
//
// 如：b:1 2
// 若a:3,b:4,a:5, a就可以赢了。
// 若a:3 4,b:5 six,a:7 8, a就可以赢了。

// 另一个秘诀：如果b说了3的倍数的数，b就输了。因为a只需要说两个数，就可以踩中必赢数组[2^n-1]
// 在a知道规则，b不知道规则时，b赢的情况只有1种，即：b先说12，并持续踩中 必赢数组[2^n-1]

// 类似的游戏
// 游戏规则：从1开始到100结束，两个人a和b轮流说1到7个数，谁先开始都可以，最后一个说出100的人就输了。
// 分析： 谁最后说99谁就赢了，类似得出必赢数组[2+8^n]，踩中必赢数组后，保持对方说x个数，我说8-x个数就可以赢了。

// 通用规律：
// 游戏规则：从1开始到num结束，两个人a和b轮流说1...x个数，谁先开始都可以，最后一个说出num的人就输了。
// 分析：必赢数组的第一个数：first = (num-1) % (x+1), 必赢数组[first + (1+x)^n],之后对方说 p,我说 1+x-p,就可以赢了。

// 假设a 知道策略
func policy_say_a(b int) (result int) {
	winArray := [...]int{29, 26, 23, 20, 17, 14, 11, 8, 5, 2}
	for _, num := range winArray{
		if b + 1 == num || b + 2 == num{
			return num
		}
	}
	return say_a(b)

}

// 随机方法
func say_a(b int) int {
	a := b + 1 + rand.Intn(2) // 1, 2
	// fmt.Println("a:", a)
	return a
}

func say_b(a int) int {
	b := a + 1 + rand.Intn(2) // 1, 2
	// fmt.Println("b:", b)
	return b
}

func game(say string) int {
	// 将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())

	a := 0
	b := 0
	if say == "a" {		// a 先说
		for a < 30 && b < 30 {
			// a 也不知道策略
			// a = say_a(b)
			// a 知道策略
			a = policy_say_a(b)
			if a >= 30 {
				// fmt.Println("b赢了这个游戏")
				return 0
			}
			b = say_b(a)
			if b >= 30 {
				// fmt.Println("a赢了这个游戏")
				return 1
			}
		}
	} else if say == "b" {		// b 先说
		var arr_a []int
		var arr_b []int
		for a < 30 && b < 30 {
			b = say_b(a)
			arr_b = append(arr_b, b)
			if b >= 30 {
				// fmt.Println("a赢了这个游戏")
				return 1
			}
			// a 不知道策略
			// a = say_a(b)
			// a 知道策略
			a = policy_say_a(b)
			arr_a = append(arr_a, a)
			if a >= 30 {
				fmt.Println("b赢了这个游戏")
				fmt.Println("b:", arr_b)
				fmt.Println("a:", arr_a)
				return 0
			}
		}
	}

	// fmt.Printf("出问题了，a：%v，b:%v\n", a, b)
	return -1
}

func start(say string) {
	a := 0
	b := 0
	c := 0
	//  重复进行游戏 10000 次，记录输赢结果
	for i := 0; i < 10000; i++ {
		g := game(say)
		if g == 1 {
			a += 1
		} else if g == 0 {
			b += 1
		} else {
			c += 1
		}
	}
	fmt.Printf("%s先说，进行游戏%v次：\n",say, 10000)
	fmt.Println("a赢了：", a, "次，b赢了", b, "次, 出现问题", c, "次。")
}

func main() {
	// “a” a 先说  ； "b" b 先说
	start("b")
	start("a")
	// game("a")
	// game("b")
}
