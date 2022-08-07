/**
回合制游戏，v，k是两个英雄，速度快的先攻击，速冻相同k先。
每三回合发动一次主动技能，不能发动主动技能的时候使用普通攻击。

v的主动技能：
普通攻击一次，使k陷入混乱状态（即，k的下次普通攻击，攻击自己）

v的被动技能：
如果v的生命不足31，v,k生命各+20, v的攻击永久+15,被动技能每场对局只能发动一次

k的主动技能：
k生命-10，对v造成45伤害，并造成20元素伤害（元素伤害无视防御值），之后k需要休息一回合，但是技能cd正常减少。 // 不清楚休息的时候被动能不能触发，不过好像也不影响
k生命小于11，只能使用普通攻击，不能使用主动技能

k的被动技能：
k每次损失5生命值，获得攻击力+1【（初始生命-当前生命）/2 后向下取整】  // 不理解这里为什么是题目给出的 /2 ，不是 /5

*/

/**
3
100 100 100 100
100 1 1 1
100 100 1 100
100 100 1 1
100 20 12 25
100 23 9 26


I love V2V forever!
Kalpas yame te!
I love V2V forever!
*/
package main

import "fmt"

// 我的题解：通过样例 20%+

func main() {
	var T int
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var vsheng, vgong, vfang, vsu int
		var ksheng, kgong, kfang, ksu int
		fmt.Scan(&vsheng, &vgong, &vfang, &vsu)
		fmt.Scan(&ksheng, &kgong, &kfang, &ksu)
		begin(vsheng, vgong, vfang, vsu, ksheng, kgong, kfang, ksu)
	}
}

func begin(vsheng, vgong, vfang, vsu, ksheng, kgong, kfang, ksu int) {

	// v 主动
	khunluan := false // k 是否混乱状态
	vzhutag := 3
	vZhuDong := func() {
		vzhutag--
		if vzhutag == 0 {
			vzhutag = 3
			if vgong-kfang > 0 {
				ksheng -= vgong - kfang
			}
			khunluan = true
		} else {
			if vgong-kfang > 0 {
				ksheng -= vgong - kfang
			}
		}
	}

	// v 被动
	vbeitag := true
	vBeiDong := func() {
		if vbeitag && vsheng < 31 {
			vsheng += 20
			ksheng += 20
			vgong += 15
			vbeitag = false
		}
	}

	// k 普通攻击
	kPu := func() {
		khunluan = false // 修改后的位置
		if kgong-vfang > 0 {
			if khunluan {
				// khunluan = false	// 原答案的位置
				ksheng -= kgong - kfang
			} else {
				vsheng -= kgong - vfang
			}
		}
	}

	// k 主动
	kxiuxi := false // k 是否休息
	kzhutag := 3
	kZhuDong := func() {
		kzhutag--
		if kxiuxi {
			kxiuxi = false
			return
		}
		if kzhutag == 0 {
			kzhutag = 3
			if ksheng < 11 {
				kPu()
			} else {
				ksheng -= 10
				if 45-vfang > 0 {
					vsheng -= 45 - vfang
				}
				vsheng -= 20
				kxiuxi = true
			}
		} else {
			kPu()
		}
	}

	// k 被动
	kchushisheng := ksheng
	kchushigong := kgong
	kBeiDong := func() {
		if kxiuxi == false {
			kgong = kchushigong + (kchushisheng-ksheng)/5
		}
	}

	// 开始
	for vsheng > 0 && ksheng > 0 {
		if vsu > ksu {
			vBeiDong()
			vZhuDong()
			kBeiDong()
			kZhuDong()
		} else {
			kBeiDong()
			kZhuDong()
			vBeiDong()
			vZhuDong()
		}
	}
	// 结果
	if vsheng > 0 { // V胜利
		fmt.Println("I love V2V forever!")
	} else { // v失败
		fmt.Println("Kalpas yame te!")
	}
}
