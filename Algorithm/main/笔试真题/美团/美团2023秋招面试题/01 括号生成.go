/**
括号生成
给出n对括号，请编写一个函数来生成所有的由n对括号组成的合法组合。


	例如，给出n=3，解集为：

	"((()))", "(()())", "(())()", "()()()", "()(())"

	数据范围：0 \le n \le 100≤n≤10

	要求：空间复杂度 O(n)O(n)，时间复杂度 O(2^n)O(2^n)

示例1
输入： 1
输出： ["()"]
示例2
输入： 2
输出： ["(())","()()"]
 */
package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param n int整型
 * @return string字符串一维数组
 */
func generateParenthesis( n int ) []string {
	// write code here
	hash := make(map[string]bool)
	hash["()"] = true
	for i:=1;i<n;i++ {
		thash :=  make(map[string]bool)
		for key := range hash{
			// 前后加括号
			thash["()" + key] = true
			thash[key + "()"] = true
			// 每个'('后面加“()”
			for j := range key {
				if key[j] == '(' {
					kuohao := key[:j+1] + "()" + key[j+1:]
					thash[kuohao] = true
				}
			}
		}
		hash = make(map[string]bool)
		for key := range thash{
			hash[key] = true
		}
	}
	//
	res := []string{}
	for key := range hash{
		res = append(res, key)
	}
	return res
}

/**
()

()() (())

()()() (())() ()(()) (()()) ((()))

()()()() (())()() ()(())() ()()(()); (()())() ((()))() (())(()) ; ()(()()) ()((())); (()()()) ((())()) (()(())); (()(())) ((()())) (((())))
["()()(())","((())())","(((())))","()()()()","(())()()","(())(())","()(())()","()((()))","((()))()","((()()))","()(()())","(()())()","(()(()))","(()()())"]
*/