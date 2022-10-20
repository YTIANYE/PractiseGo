/**
一个字符串
开始是
10
然后对之前的1-2按位取反续接在后面变成：10 01
再然后对1-4位按位取反续接在后面变成：1001 0110
再然后对1-8位按位取反后续接在后面变成：10010110 01101001
...
这个字符串无限长

为给定范围[l,r]，左闭右闭，求这个范围上 1 的个数是多少？

 */
package main

import "fmt"

// 100%
/**
由于所有的数都是由最开始的 10 反转得来，所以两个为一组，其中必有一个是1，只需要单独判断首尾的情况
 */
func main () {
	var l, r int64
	fmt.Scan(&l, &r)
	var res int64
	if l % 2 == 1 {  // l 奇数
		if r % 2 == 0 {
			res = (r + 1 -l) / 2
		}else {
			res = (r -l) / 2
			if isOne(r) {
				res++
			}
		}
	}else {  // l 偶数
		if isOne(l) {
			res++
		}
		if r % 2 == 0 {
			res += (r - l) / 2
		}else {
			res += (r-l) / 2
			if isOne(r) {
				res++
			}
		}
	}
	fmt.Println(res)

}

// 判断一个数是不是 1 ,如果不是，就探索它是由谁取反得来的，递归向前探索
func isOne(i int64) bool {
	if i == 1 {
		return true
	}
	if i == 2 {
		return false
	}
	var t int64
	t = 2
	for t < i {
		t<<=1
	}
	// t >= i
	j := t>>1 - (t-i)  // t一直是2的倍数
	return !isOne(j)
}
/**
10
10 01
1001 0110
1001 0110 0110 1001

1 1
0 2

1
0


 */
