package main

import "fmt"

// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// 分别完成在队列尾部插入整数和在队列头部删除整数的功能。
// (若队列中没有元素，deleteHead操作返回 -1 )
//
// 示例 1：
// 输入：
// ["CQueue","appendTail","deleteHead","deleteHead"]
// [[],[3],[],[]]
// 输出：[null,null,3,-1]
//
// 示例 2：
// 输入：
// ["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
// [[],[],[5],[2],[],[]]
// 输出：[null,-1,null,null,5,2]
//
// 提示：
// 1 <= values <= 10000
// 最多会对appendTail、deleteHead 进行10000次调用

// 栈相关操作
type Stack struct {
	nums []int
}

func pop(stack *Stack) (num int) {
	if len(stack.nums) != 0 {
		num = stack.nums[len(stack.nums)-1]
		stack.nums = stack.nums[:len(stack.nums)-1]
		return num
	}
	return
}

// 执行用时：300 ms, 在所有 Go 提交中击败了5.16%的用户
// 内存消耗：8.4 MB, 在所有 Go 提交中击败了71.89%的用户

type CQueue struct {
	left  Stack
	right Stack
}

func Constructor() CQueue {
	var cqueue CQueue
	return cqueue

}

func (this *CQueue) AppendTail(value int) {
	this.right.nums = append(this.right.nums, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.left.nums) != 0 {
		return pop(&this.left)
	}

	for len(this.right.nums) != 0 {
		this.left.nums = append(this.left.nums, pop(&this.right))
	}
	if len(this.left.nums) != 0 {
		return pop(&this.left)
	} else {
		return -1
	}
}

func main() {
	// in := []string{"CQueue", "appendTail", "deleteHead", "deleteHead"}
	in := []string{"CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"}
	var cqueue CQueue
	for _,str := range in {
		if str == "CQueue" {
			cqueue = Constructor()
		} else if str == "appendTail" {
			var value int
			fmt.Scan(&value)
			cqueue.AppendTail(value)
		} else if str == "deleteHead" {
			fmt.Println(cqueue.DeleteHead())
		}

	}
}
