package linkedlist

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
//
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 示例：
//
//	输入：head = [1,2,3,4,5], n = 2
//	输出：[1,2,3,5]
//	解释：删除了倒数第 2 个节点 4
//
// 思路：快慢双指针（一次遍历）
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	k := n
	dump := &ListNode{Val: 0, Next: head}
	cur := head
	for cur != nil && k > 0 {
		cur = cur.Next
		k--
	}
	if k > 0 {
		return nil
	}
	// 要删除节点的前一个结点
	ahead := dump
	for cur != nil {
		cur = cur.Next
		ahead = ahead.Next
	}
	ahead.Next = ahead.Next.Next
	return dump.Next
}
