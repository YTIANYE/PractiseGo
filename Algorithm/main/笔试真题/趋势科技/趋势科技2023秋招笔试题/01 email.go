/**
题目：
一个模板，在两个 % 之间存在的key替换成value

示例：
第一行，模板内容
第二行，key
第三行，value

"%FILENAME%VIRUSNAME%balabalabala%VIRUSNAME%FILENAME%",
["SENDER","RCPTS","SUBJECT","FILENAME","VIRUSNAME","ACTION"],
["abc@test.com","test@company.com","Test","Myfile","ICAN","Delete"]

"MyfileVIRUSNAME%balabalabalaICANFILENAME%"
*/

package main

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param my_template string字符串
 * @param keys string字符串一维数组
 * @param values string字符串一维数组
 * @return string字符串
 */

// 我的题解： 100%
func token_replace(my_template string, keys []string, values []string) string {
	// write code here
	hashmap := make(map[string]string, len(keys))
	for i := 0; i < len(keys); i++ {
		hashmap[keys[i]] = values[i]
	}
	//
	res := ""
	var count, left, right int
	var find func(string) bool
	find = func(s string) bool {
		count = 0
		left, right = 0, 0

		for i := range s {
			if s[i] == '%' {
				count++

				if count == 1 {
					left = i
				} else if count == 2 {
					right = i
					key := s[left+1 : right]
					if value, ok := hashmap[key]; ok { // 是key
						res = res + s[:left] + value
						count = 0
						my_template = my_template[right+1:]
						return true
					} else { // 不是key
						count = 1
						left = right
						res = res + s[:left]
						my_template = my_template[left:]
						return true
					}
				}
			}
		}
		return false
	}

	//
	for find(my_template) {
	}
	res = res + my_template
	return res
}
