/*
问题描述：
只包含'A' 'B' 的字符串
经过最少多少步可以转换为互斥字符串（即字符串中每个字符与其相邻字符均不相同）
*/

package 虾皮2022春招笔试题

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 *
 * @param s string字符串
 * @return int整型
 */

// 我的题解：通过全部测试样例
func MinOperations(s string) int {
	// write code here
	n := len(s)
	res1, res2 := 0, 0

	for i := 0; i < n; i++ {
		// ABABABA
		if i%2 == 0 && s[i] != 'A' {
			res1++
		} else if i%2 == 1 && s[i] != 'B' {
			res1++
		}
		// BABABAB
		if i%2 == 0 && s[i] != 'B' {
			res2++
		} else if i%2 == 1 && s[i] != 'A' {
			res2++
		}
	}

	if res1 < res2 {
		return res1
	}
	return res2

}
