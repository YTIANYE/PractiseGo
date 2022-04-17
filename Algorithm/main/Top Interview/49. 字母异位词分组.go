/**
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

 

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]
 

提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母

 */

package main

import (
	"sort"
	"strconv"
)
// 我的题解：哈希+排序
/**
执行用时：28 ms, 在所有 Go 提交中击败了23.07的用户
内存消耗：8.6 MB, 在所有 Go 提交中击败了40.07%的用户
 */

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	hash := make(map[string][]string)
	for _, str := range strs{
		key := getKey(str)
		hash[key] = append(hash[key], str)
	}
	for _, val := range hash{
		res = append(res, val)
	}
	return res
}

func getKey(str string) string {
	s := []uint8{}
	for i := range str{
		s = append(s, str[i])
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	key := ""
	for i:= range s{
		k := strconv.Itoa(int(s[i]))
		key += k
	}
	return key
}