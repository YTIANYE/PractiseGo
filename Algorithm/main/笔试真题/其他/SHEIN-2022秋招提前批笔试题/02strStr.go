/**
haystack字符串中是否存在needle字符串
 */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param haystack string字符串
 * @param needle string字符串
 * @return int整型
 */


/**

"shein","in"
3

"aaaaa", "bba"
-1

"",""
0
 */

func strStr( haystack string ,  needle string ) int {
	// write code here

	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	if n == 0 {
		return -1
	}

	// 比对

	for i:=0;i<n;i++{
		if haystack[i] == needle[0] {
			if fun(haystack[i:], needle) {
				return i
			}
		}
	}
	return -1
}

func fun(s1, s2 string) bool  {
	n, m := len(s1), len(s2)
	if m >n {
		return false
	}
	for i:=0;i<m;i++{
		if s2[i] != s1[i] {
			return false
		}
	}
	return true
}