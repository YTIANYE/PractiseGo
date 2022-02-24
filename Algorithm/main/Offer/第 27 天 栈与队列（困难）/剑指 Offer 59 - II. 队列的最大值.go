/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，
要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value需要返回 -1

示例 1：
输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出:[null,null,null,2,1,2]
示例 2：
输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出:[null,-1,-1]

限制：
1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5
*/
package main

// 我的题解：[单调双端队列]维护一个maxs的递减序列
/*
执行用时：72 ms, 在所有 Go 提交中击败了92.17%的用户
内存消耗：8.2 MB, 在所有 Go 提交中击败了71.59%的用户
*/

type MaxQueue struct {
	maxs   []int
	values []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxs) == 0 {
		return -1
	}
	return this.maxs[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.maxs) != 0 && value > this.maxs[len(this.maxs)-1] {
		this.maxs = this.maxs[:len(this.maxs)-1]
	}
	this.maxs = append(this.maxs, value)
	this.values = append(this.values, value)
}

func (this *MaxQueue) Pop_front() int {
	res := -1
	if len(this.values) == 0 {
		return res
	}

	if len(this.values) == 1 {
		res = this.values[0]
		this.maxs = []int{}
		this.values = []int{}
		return res
	}

	if this.values[0] == this.maxs[0] {
		this.maxs = this.maxs[1:]
	}
	res = this.values[0]
	this.values = this.values[1:]
	return res
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
