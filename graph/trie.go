package graph

/*
208. 实现 Trie (前缀树)：https://leetcode.cn/problems/implement-trie-prefix-tree/


Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。
这一数据结构有相当多的应用情景，例如自动补全和拼写检查。
*/

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

// 构造函数
func Constructor() Trie {
	return Trie{}
}

// 插入
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
	node.isEnd = true // 标记完整单词结束
}

// 搜索完整单词
func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd // 必须 isEnd 为 true
}

// 搜索前缀
func (t *Trie) StartsWith(prefix string) bool {
	node := t.searchPrefix(prefix)
	return node != nil // 不需要 isEnd
}

// 辅助函数：找到 prefix 对应的末尾节点
func (t *Trie) searchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}
