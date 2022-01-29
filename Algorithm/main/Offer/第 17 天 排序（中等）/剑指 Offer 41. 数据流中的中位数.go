/*
如何得到一个数据流中的中位数？
如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，

[2,3,4]的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例 1：

输入：
["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]
输出：[null,null,null,1.50000,null,2.00000]
示例 2：

输入：
["MedianFinder","addNum","findMedian","addNum","findMedian"]
[[],[2],[],[3],[]]
输出：[null,null,2.00000,null,2.50000]

限制：

最多会对 addNum、findMedian 进行 50000 次调用。

*/
package main

import "fmt"

type MedianFinder struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{}
	return mf
}

func (this *MedianFinder) AddNum1(num int) {
	if len(this.nums) == 0 {
		this.nums = append(this.nums, num)
	} else {
		for i, n := range this.nums {
			if n > num {
				this.nums = append(this.nums[:i], append([]int{num}, this.nums[i:]...)...)
				break
			} else if i == len(this.nums)-1 {
				this.nums = append(this.nums, num)
			}
		}
	}
}

/*
执行用时：360 ms, 在所有 Go 提交中击败了14.87%的用户
内存消耗：23 MB, 在所有 Go 提交中击败了12.84%的用户
*/

func (this *MedianFinder) AddNum(num int) {
	if len(this.nums) == 0 {
		this.nums = append(this.nums, num)
		return
	}

	index := this.fast(0, len(this.nums)-1, num)
	this.nums = append(this.nums[:index], append([]int{num}, this.nums[index:]...)...)
}

// 二分查找
func (this *MedianFinder) fast(left, right, num int) int {

	mid := (left + right) / 2
	if left >= right {
		if this.nums[mid] < num {
			return mid + 1
		} else {
			return mid
		}
	}

	if this.nums[mid] > num {
		return this.fast(left, mid-1, num)
	} else if this.nums[mid] < num {
		return this.fast(mid+1, right, num)
	} else {
		return mid
	}
}

func (this *MedianFinder) FindMedian() float64 {
	l := len(this.nums)

	if l%2 == 0 {
		return float64(this.nums[l/2]+this.nums[l/2-1]) / 2
	} else {
		return float64(this.nums[l/2])
	}
}

func main() {
	var mf MedianFinder
	//orders := []string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"}
	//nums := []int{0, -1, 0, -2, 0, -3, 0, -4, 0, -5, 0}
	orders := []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"}
	nums := []int{0, 1, 2, 0, 3, 0}

	for i, order := range orders {
		if order == "MedianFinder" {
			mf = MedianFinder{}
			fmt.Println("nil")
		} else if order == "addNum" {
			mf.AddNum(nums[i])
			fmt.Println("nil")
		} else if order == "findMedian" {
			fmt.Println(mf.FindMedian())
		}
	}
}
