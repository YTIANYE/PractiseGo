/**
同花顺游戏
N 副牌， N<8，求这副牌组成最大的牌是多少？
 */
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 翻牌
 * @param inHand string字符串 一组以单空格间隔的手牌(例如：SA HK H9 D8 C5 S5 H4)
 * @return string字符串
 */

// 30分的66.7%
func showDown(inHand string) string {
	// write code here
	strs := strings.Split(inHand, " ")

	huas := []int{}
	dians := []int{}
	for i := range strs {

		var j int
		if strs[i][0] == 'S' {
			j = 0
		}else if strs[i][0] == 'H' {
			j = 1
		}else if strs[i][0] == 'C' {
			j = 2
		}else {
			j = 3
		}
		huas = append(huas, j)
		num := strs[i][1:]
		dianshu := 0
		if num == "A" {
			dianshu = 1
		} else if num == "K" {
			dianshu = 13
		} else if num == "Q" {
			dianshu = 12
		} else if num == "J" {
			dianshu = 11
		} else {
			dianshu, _ = strconv.Atoi(num)
		}
		dians = append(dians, dianshu)
	}
	//
	res := judge(huas, dians)
	if res == 1 {
		return "HuangJiaTongHuaShun"
	} else if res == 2 {
		return "TongHuaShun"
	} else if res == 3 {
		return "SiTiao"
	} else if res == 4 {
		return "HuLu"
	} else if res == 5 {
		return "TongHua"
	} else if res == 6 {
		return "ShunZi"
	} else if res == 7 {
		return "SanTiao"
	} else if res == 8 {
		return "LiangDui"
	} else if res == 9{
		return "YiDui"
	}else {
		return "GaoPai"
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func judge(huas []int, dians []int) int {

	res := 10
	n := len(dians)
	hashdian := make(map[int]int) // key：dian val:count
	for _, dian := range dians {
		hashdian[dian]++
	}
	var count2, count3, count4 int
	for _, val := range hashdian {
		if val == 2 {
			count2++
		} else if val == 3 {
			count3++
		} else if val == 4 {
			count4++
		}
	}

	// 3 4 7 8 9
	if count4 != 0 { // 3 四条
		res = min(res, 3)
	}
	if count3 != 0 {
		if count2 != 0 { // 4 葫芦
			res = min(res, 4)
		} else { // 7 三条
			res = min(res, 7)
		}
	}
	if count2 == 1 { // 9 一对
		res = min(res, 9)
	} else if count2 == 2 { // 8 两对
		res = min(res, 8)
	}
	//
	nums := []int{}  // 存放每一种数
	for key := range hashdian {
		nums = append(nums, key)
	}
	sort.Ints(nums)
	l := len(nums)
	shunzis := [][]int{}  // 每一行由小到大
	if l >= 5 {
		if nums[0] == 1 && nums[l-1] == 13 && nums[l-2] == 12 && nums[l-3] == 11 && nums[l-4] == 10{
			res = min(res, 6)  // 6 顺子
			shunzi := []int{}
			shunzi = append(shunzi, nums[l-4])
			shunzi = append(shunzi, nums[l-3])
			shunzi = append(shunzi, nums[l-2])
			shunzi = append(shunzi, nums[l-1])
			shunzi = append(shunzi, nums[0])
			shunzis = append(shunzis, shunzi)
		}else {
			for i:=0;i<=l-5;i++ {
				tag := true // 是顺子

				for j:=i;j<j+4;j++ {
					if nums[j+1] - nums[j] != 1 {
						tag = false
						break
					}
				}
				if tag {
					res = min(res, 6)  // 6 顺子
					shunzi := []int{}
					for j:=i;j<=j+4;j++ {
						shunzi = append(shunzi, nums[j])
					}
				}
			}
		}
	}
	// 判断花色
	hashhua := make(map[int]int) // key:huase val :count
	for _, hua := range huas {
		hashhua[hua]++
	}
	for _, val := range hashhua {
		if val == 5 {// 5 同花
			res = min(res, 5)
		}
	}
	// 判断同花顺、皇家同花顺

	yanse := make([][]bool, 14) // key:点  key：花色 val :是否存在
	for i:= range yanse {
		yanse[i] = make([]bool, 4) // SHCD
	}

	for i:=0;i<n;i++ {
		yanse[dians[i]][huas[i]] = true
	}
	for _, shunzi := range shunzis {
		hashcolor := make(map[int]int)
		for _, num:= range shunzi {
			hashtemp := yanse[num]
			for key := range hashtemp {
				if hashtemp[key] {
					hashcolor[key]++
				}
			}
		}
		//
		for _, val := range hashcolor {
			if val == 5 {
				res = min(res, 2) // 2 同花顺
				if shunzi[0] == 10 {
					res = min(res, 1)  // 1 皇家同花顺
					return 1
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(showDown("SA SK SQ SJ S10 H10 C9"))
}
