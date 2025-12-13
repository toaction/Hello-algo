package linkedlist

/*
148. 排序链表：https://leetcode.cn/problems/sort-list/

给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

示例：
	输入：head = [4,2,1,3]
	输出：[1,2,3,4]
*/

// 归并排序
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMiddleNode(head)
	next := mid.Next
	mid.Next = nil
	first, second := SortList(head), SortList(next)
	return MergeTwoLists(first, second)
}

// 返回链表中间节点（奇数）或前半部分的尾节点（偶数）
func findMiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
