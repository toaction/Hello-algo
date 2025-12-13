package linkedlist

/*
206. 反转链表：https://leetcode.cn/problems/reverse-linked-list/

给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：
	输入：head = [1,2,3,4,5]
	输出：[5,4,3,2,1]
*/

// 头插法反转链表
func ReverseList(head *ListNode) *ListNode {
	dump := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = dump.Next
		dump.Next = head
		head = next
	}
	return dump.Next
}
