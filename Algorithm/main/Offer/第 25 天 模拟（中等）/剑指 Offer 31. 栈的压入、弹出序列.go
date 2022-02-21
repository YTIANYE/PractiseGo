/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
假设压入栈的所有数字均不相等。
例如，序列 {1,2,3,4,5} 是某栈的压栈序列，
序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，
但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

示例 1：
输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：
输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

提示：
0 <= pushed.length == popped.length <= 1000
0 <= pushed[i], popped[i] < 1000
pushed是popped的排列。
*/

package main

// 我的题解：模拟
/*
执行用时：8 ms, 在所有 Go 提交中击败了22.90%的用户
内存消耗：3.6 MB, 在所有 Go 提交中击败了95.71%的用户
*/

func validateStackSequences1(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	if len(popped) == 0 && len(pushed) == 0 {
		return true
	}

	i := 0
	stack := []int{}

	for len(popped) != 0 {
		if in(stack, popped[0]) {
			if stack[len(stack)-1] != popped[0] {
				return false
			}
			// 相等
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		} else { // 栈中没有
			stack = append(stack, pushed[i])
			i++
		}
	}
	return true
}

func in(arr []int, x int) bool {
	for _, num := range arr {
		if num == x {
			return true
		}
	}
	return false
}

// 我实现的精选题解：另一种模拟方式
/*
执行用时：4 ms, 在所有 Go 提交中击败了96.06%的用户
内存消耗：3.6 MB, 在所有 Go 提交中击败了82.83%的用户
*/
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0

	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			i++
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
