package linkedlist

/*
19. 删除链表的倒数第 N 个结点：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例：
	输入：head = [1,2,3,4,5], n = 2
	输出：[1,2,3,5]
	解释：删除了倒数第 2 个节点 4
*/

// 双指针
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dump := &ListNode{Next: head}
	// 要删除节点的前一个节点、链表的尾节点
	pre, end := dump, head
	i := 0
	for i = 0; i < n && end != nil; i++ {
		end = end.Next
	}
	if i != n {
		return nil
	}
	// 双指针同时移动
	for end != nil {
		pre = pre.Next
		end = end.Next
	}
	// 删除节点
	pre.Next = pre.Next.Next
	return dump.Next
}
