//
//func main() {
//	var n, m int
//	fmt.Scan(&n,&m)
//	nums := make([]int, n+1)
//	for i:=0;i<=n;i++ {
//		nums[i] = i
//	}
//
//	for ;m != 0 ;m-- {
//		var a int
//		fmt.Scan(&a)
//
//		if a == 1 {
//			continue
//		}else if a == n {
//			s1 := make([]int, 1)
//			copy(s1, nums[:1])
//			s2 := make([]int, 1)
//			copy(s2, nums[n:])
//			s3 := make([]int, len(nums[1:n]))
//			copy(s3, nums[1:n])
//			t1 := append(s1, s2...)
//			t2 := s3
//			nums = append(t1, t2...)
//		}else {
//			b := a
//
//			s1 := make([]int, len(nums[:1]))
//			copy(s1, nums[:1])
//			s2 := make([]int, len(nums[b:b+1]))
//			copy(s2, nums[b:b+1])
//			s3 := make([]int, len(nums[1:b]))
//			copy(s3, nums[1:b])
//			s4 := make([]int, len(nums[b+1:]))
//			copy(s4, nums[b+1:])
//			t1 := append(s1, s2...)
//			t2 := append(s3, s4...)
//			nums = append(t1, t2...)
//		}
//		// fmt.Println(nums)
//	}
//	//
//	fmt.Printf("%d", nums[1])
//	for i:=2;i<=n;i++ {
//		fmt.Printf(" %d", nums[i])
//	}
//}

/**
挑剔
时间限制： 3000MS
内存限制： 589824KB
题目描述：
小美有一个精致的珠宝链子。初始这个链子上有n个宝石，从左到右分别编号为1~n （宝石上的编号不会因为交换位置而改变编号）。接着，小美为了美观对这个项链进行微调，有m次操作，每次选择一个编号 x ,将编号 x 的宝石放到最左边（不改变其他宝石的相对位置）。小美想知道，所有操作完成后最终链子从左到右宝石编号是多少。



输入描述
第一行两个正整数n和m，分别表示链子上的宝石数和操作次数。

接下来一行m个数 x1,x2,...,xm ，依次表示每次操作选择的编号x值。

数字间两两有空格隔开

对于所有数据，1≤m,n≤50000, 1≤xi≤n

输出描述
输出一行 n 个整数，表示答案。


样例输入
5 3
2 3 4
样例输出
4 3 2 1 5

提示
第一次微调完，链子为 2 1 3 4 5

第二次微调完，链子为 3 2 1 4 5

第三次微调完，链子为 4 3 2 1 5
 */
package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Pre *Node
}

// 64%
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	nodes := make(map[int]*Node)
	head := &Node{0, nil, nil}
	p := head
	nodes[0] = head
	for i := 1; i <= n; i++ {
		node := &Node{i, nil, p}
		p.Next = node
		p = p.Next
		nodes[i] = node
	}
	//
	var move func(pre, t *Node)
	move = func(pre, t *Node) {
		pre.Next = t.Next
		t.Next.Pre = pre
		t.Next = head.Next
		head.Next.Pre = t
		head.Next = t
		t.Pre = head
	}
	//
	nums := make([]int, m)
	for i:=0;i<m;i++ {
		var a int
		fmt.Scan(&a)
		nums[i] = a
	}
	for i:= range nums {
		a := nums[i]
		t := nodes[a]
		pre := t.Pre
		move(pre, t)
	}

	//
	q := head.Next
	fmt.Printf("%d", q.Val)
	q = q.Next
	for q != nil {
		fmt.Printf(" %d", q.Val)
		q = q.Next
	}
}
