package linkedlist

/*
146. LRU 缓存：https://leetcode.cn/problems/lru-cache/

请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；
如果不存在，则向缓存中插入该组 key-value 。
如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
*/

type DListNode struct {
	key   int
	value int
	prev  *DListNode
	next  *DListNode
}

type LRUCache struct {
	capacity int
	size     int
	data     map[int]*DListNode
	head     *DListNode
	tail     *DListNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		size:     0,
		data:     make(map[int]*DListNode),
		head:     &DListNode{},
		tail:     &DListNode{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

// 若存在，将节点移动到 head
func (cache *LRUCache) Get(key int) int {
	DListNode := cache.data[key]
	if DListNode == nil {
		return -1
	}
	// 删除节点
	cache.remove(DListNode)
	// 插入到头部
	cache.addHead(DListNode)
	return DListNode.value
}

// 从链表中删除节点
func (cache *LRUCache) remove(DListNode *DListNode) {
	DListNode.prev.next, DListNode.next.prev = DListNode.next, DListNode.prev
	cache.size--
}

// 将节点添加到链表首部
func (cache *LRUCache) addHead(DListNode *DListNode) {
	DListNode.prev = cache.head
	DListNode.next = cache.head.next
	cache.head.next.prev = DListNode
	cache.head.next = DListNode
	cache.size++
}

func (cache *LRUCache) Put(key int, value int) {
	if v, ok := cache.data[key]; ok {
		v.value = value
		cache.Get(key)
		return
	}
	// 删除尾部节点
	if cache.size == cache.capacity {
		DListNode := cache.tail.prev
		delete(cache.data, DListNode.key)
		cache.remove(DListNode)
	}
	// 插入
	DListNode := &DListNode{key: key, value: value}
	cache.data[key] = DListNode
	cache.addHead(DListNode)
}
