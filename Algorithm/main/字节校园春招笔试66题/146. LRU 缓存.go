/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

 

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

提示：

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put

*/

package main

import "fmt"

// 方法：哈希+双向链表

type Node struct {
	val   int
	left  *Node
	right *Node
}

type LRUCache struct {
	nums     map[int]*Node
	head     *Node // 头插入新结点
	tail     *Node
	count    int
	capacity int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.nums = make(map[int]*Node)

	lru.head = &Node{-1, nil, nil}
	lru.tail = &Node{-1, nil, nil}
	lru.head.right = lru.tail
	lru.tail.left = lru.head

	lru.count = 0
	lru.capacity = capacity
	return lru
}

func (this *LRUCache) Get(key int) int {
	node := this.nums[key]
	if node == nil || node.val == -1 {
		return -1
	}
	node.left.right = node.right
	node.right.left = node.left
	node.right = this.head.right
	this.head.right.left = node // 注意this.head.right
	node.left = this.head
	this.head.right = node
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.nums[key]
	if ok && node.val != -1 {
		node.val = value
		node.left.right = node.right
		node.right.left = node.left
	} else {
		node = &Node{value, nil, nil}
		this.nums[key] = node
		this.count++
	}

	node.right = this.head.right
	node.right.left = node // 注意this.head.right的左指针
	node.left = this.head
	this.head.right = node

	if this.count > this.capacity {
		remove := this.tail.left
		this.tail.left = remove.left
		remove.left.right = this.tail
		// 删除remove
		// map 中的删除可以用delete(this.nums, key)来实现
		remove.val = -1
		remove.left = nil
		remove.right = nil
		this.count--
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	// ["LRUCache", "put", "put", "get", "put","get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]

	// ["LRUCache", "get"]
	// [[1],[0]]

	// ["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
	// [[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]

	var lru LRUCache
	var LRU string
	fmt.Scan(&LRU)
	if LRU == "LRUCache" {
		var capacity int
		fmt.Scan(&capacity)
		lru = Constructor(capacity)
	}
	for true {
		var order string
		fmt.Scan(&order)
		if order == "put" {
			var value, key int
			fmt.Scan(&key, &value)
			lru.Put(key, value)
		} else if order == "get" {
			var key int
			fmt.Scan(&key)
			fmt.Println(lru.Get(key))
		}
	}
}
