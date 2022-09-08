/**
给出一个字符串，存在相同字符
每次操作是删除2个相同字符增加一个字符，
问使结果中不存在相同字符，最少操作多少次
 */
package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 返回满足题意的最小操作数
 * @param str string字符串 给定字符串
 * @return int整型
 */
// 100%
func minOperations( str string ) int {
	// write code here

	hash := make(map[byte]int)
	for i := range str {
		hash[str[i]]++
	}
	count := 0
	for key, val := range hash {
		for val >= 2 {
			val -= 2
			count++
		}
		if val == 0 {
			delete(hash, key)
		}
	}
	//fmt.Println(len(hash))
	//fmt.Println(count)
	if 26 - len(hash) >= count {
		return count
	}else {
		return count + (count - (26 - len(hash)))
	}
}