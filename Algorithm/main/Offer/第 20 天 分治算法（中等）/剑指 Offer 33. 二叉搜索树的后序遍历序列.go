/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true

*/

package main

// 我的题解：分治
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了98.36%的用户
*/
func verifyPostorder1(postorder []int) bool {
	l := len(postorder)
	if l == 1 || l == 0 {
		return true
	}
	num := postorder[l-1]
	i := l - 2
	for ; i >= 0; i-- {
		if postorder[i] < num {
			break
		}
	}

	right := postorder[i+1 : l-1]
	left := postorder[:i+1]

	for _, temp := range left {
		if temp > num {
			return false
		}
	}
	return verifyPostorder(left) && verifyPostorder(right)

}


// 我实现的精选题解：分治
func verifyPostorder2(postorder []int) bool{
	return fun(postorder, 0, len(postorder)-1)
}

func fun(postorder []int, i, j int) bool{
	if i>=j{
		return true
	}

	p := i
	for postorder[p] < postorder[j]{
		p++
	}
	m := p
	for postorder[p] > postorder[j] {
		p++
	}
	return p==j && fun(postorder, i, m-1) && fun(postorder, m, j-1)
}

// 我实现的精选题解:辅助单调栈
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了95.08%的用户
*/

func verifyPostorder(postorder []int) bool{
	stack := []int{}
	root := int(^uint(0) >> 1)
	for i:= len(postorder)-1;i>-1;i--{
		if postorder[i]>root{
			return false
		}
		for len(stack)!= 0 && postorder[i] < stack[len(stack)-1]{
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
