package linkedlist

/*
24. 两两交换链表中的节点：https://leetcode.cn/problems/swap-nodes-in-pairs/

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

示例 1：
	输入：head = [1,2,3,4]
	输出：[2,1,4,3]
*/

func SwapPairs(head *ListNode) *ListNode {
	dump := &ListNode{Next: head}
	cur := dump
	for head != nil && head.Next != nil {
		next := head.Next.Next
		cur.Next = head.Next
		cur.Next.Next = head
		head.Next = next
		cur = head
		head = next
	}
	return dump.Next
}
