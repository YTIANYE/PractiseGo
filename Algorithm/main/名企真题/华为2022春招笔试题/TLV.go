/**
题目：
接收一个一段数据
0F04ABABABAB...
一个信息段分为3个部分
第一个字节是目标号  0F=15，第二个字节表示长度 04 = 4个字节，第三段是前面描述的长度的信息，
在之后是新的信息段
如果最后一个信息段残缺不全，抛弃该信息段

给出全部的目标号，输出每个目标号对应的信息的长度和偏移量（描述的信息开始位置，如AB在第2个字节上）
如：15 4 2

输入：
信息段
几个目标号
目标号

例子1：
输入：
0F04ABABABAB
1
15
输出：
4 2

例子2：
输入：
0F04ABABABAB0102AB
2
15
1
输出：
4 2
2 8
*/

package main

import "fmt"

// 我的题解：通过率 100%

func main() {
	var str string
	var n int
	hashTag := make(map[int][]int) // key：tag             value：length offset
	tags := []int{}                // 目标号【为了实现hashTag 的有序输出】

	fmt.Scan(&str)
	fmt.Scan(&n)
	for ; n != 0; n-- {
		var tag int
		fmt.Scan(&tag)
		hashTag[tag] = []int{0, 0}
		tags = append(tags, tag)
	}

	// 计算
	hashMap := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
	}
	i := 0
	for i < len(str) {
		if i+4 >= len(str) { // 连目标号和长度都未能满足的残缺字段
			break
		}

		tag := toInt(str[i:i+2], hashMap)
		length := toInt(str[i+2:i+4], hashMap) // value 中的字节数
		offset := (i + 4) / 2                  // str中的两个字符 视为 offset的一个字节
		i += 4 + length*2                      // 下一字段的开始位置
		if i > len(str) {                      // 当前字段不完整
			break
		}
		if _, ok := hashTag[tag]; ok {
			hashTag[tag][0] += length
			hashTag[tag][1] = offset
		}
	}

	// 输出
	for _, tag := range tags {
		fmt.Println(hashTag[tag][0], hashTag[tag][1])
	}

}

// 一个字节转换成 int
func toInt(str string, hashMap map[string]int) int {
	a := hashMap[string(str[0])]
	b := hashMap[string(str[1])]
	res := b + a*16
	return res
}
