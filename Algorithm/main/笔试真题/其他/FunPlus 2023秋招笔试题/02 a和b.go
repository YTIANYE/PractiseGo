/**
	有a b 两个字母组成一个长度为m的字符串
	a可以连续，b不可以连续
	问有多少种组合方式？
	最多可以连续n次【注意这个条件】
 */

// 可以引⼊的库和版本相关请参考 “环境说明”
// 必须定义⼀个 包名为 `main` 的包

package main

// 本题面试官已设置测试用例
func calcCount(m int, n int) int {
	// 在这⾥写代码
	dpa := make([]int, m+1)
	dpb := make([]int, m+1)
	dpa[1] = 1
	dpb[1] = 1
	for i:=2;i<=m;i++ {
		dpa[i] = dpa[i-1] + dpb[i-1]
		dpb[i] = dpa[i-1]
	}
	return dpa[m] + dpb[m]
}