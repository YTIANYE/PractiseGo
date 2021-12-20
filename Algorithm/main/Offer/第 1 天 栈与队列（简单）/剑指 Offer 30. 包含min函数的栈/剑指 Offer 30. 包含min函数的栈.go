/*
定义栈的数据结构，
请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

提示：
各函数的调用总次数不超过 20000 次

*/

package main

import (
	"fmt"
)

/*我实现的官方题解*/
type MinStack struct {
	nums []int
	mins []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var minstack MinStack
	return minstack
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, x)
	} else if this.mins[len(this.mins)-1] >= x { // 注意相等的时候也添加，方便pop
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.nums) == 0{
		return
	}
	if this.nums[len(this.nums)-1] == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) Min() int {
	return this.mins[len(this.mins)-1]
}

/*我的题解*/
// 执行用时：16 ms, 在所有 Go 提交中击败了79.72%的用户
// 内存消耗：8.1 MB, 在所有 Go 提交中击败了39.41%的用户

// const INT_MAX = int(^uint(0) >> 1) // 正无穷
//
// type MinStack struct {
// 	min  int
// 	nums []int
// }
//
// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	} else {
// 		return y
// 	}
// }
//
// func Constructor() MinStack {
// 	var minstack MinStack
// 	minstack.min = INT_MAX
// 	minstack.nums = []int{}
// 	return minstack
// }
//
// func (this *MinStack) Push(x int) {
// 	this.nums = append(this.nums, x)
// 	this.min = min(x, this.min)
// }
//
// func (this *MinStack) Pop() {
// 	if len(this.nums) == 1 {
// 		this.min = INT_MAX
// 		this.nums = nil
// 		return
// 	}
//
// 	if this.min == this.nums[len(this.nums)-1] {
// 		this.min = INT_MAX
// 		for _, val := range this.nums[:len(this.nums)-1] {
// 			this.min = min(this.min, val)
// 		}
// 	}
// 	this.nums = this.nums[:len(this.nums)-1]
//
// }
//
// func (this *MinStack) Top() int {
// 	return this.nums[len(this.nums)-1]
// }
//
// func (this *MinStack) Min() int {
// 	return this.min
// }

func main() {
	var minstack MinStack
	// ["MinStack","push","push","push","top","pop","min","pop","min","pop","push","top","min","push","top","min","pop","min"]
	// 	[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
	for {
		var order string
		fmt.Scan(&order)
		switch order {
		case "init":
			minstack = Constructor()
		case "push":
			var num int
			fmt.Scan(&num)
			minstack.Push(num)
		case "pop":
			minstack.Pop()
		case "min":
			fmt.Println(minstack.Min())
		case "top":
			fmt.Println(minstack.Top())
		}

	}

}
