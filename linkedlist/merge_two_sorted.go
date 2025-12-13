package linkedlist

/*
21. 合并两个有序链表：https://leetcode.cn/problems/merge-two-sorted-lists/

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：
	输入：1->2->4, 1->3->4
	输出：1->1->2->3->4->4
*/

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dump := &ListNode{}
	cur := dump
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}
	return dump.Next
}
