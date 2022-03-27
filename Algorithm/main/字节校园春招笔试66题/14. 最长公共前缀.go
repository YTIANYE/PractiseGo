/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

 

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
 

提示：

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main


// 我的题解：暴力遍历
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1{
		return strs[0]
	}

	res := ""
	for i:=1;i<=len(strs[0]);i++{
		temp := strs[0][:i]
		for _, str := range strs{
			// 注意第一个字符串的长度可能比后面的字符串的长度大。
			if len(str) <len(temp) || temp != str[:i]{
				return res
			}
		}
		res = temp
	}
	return res

}