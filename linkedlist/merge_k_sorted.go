package linkedlist

/*
23. 合并 K 个升序链表：https://leetcode.cn/problems/merge-k-sorted-lists/

给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/

// 分治
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	p, q := merge(lists, left, mid), merge(lists, mid+1, right)
	return MergeTwoLists(p, q)
}
