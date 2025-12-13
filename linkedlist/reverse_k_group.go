package linkedlist

/*
25. K 个一组翻转链表：https://leetcode.cn/problems/reverse-nodes-in-k-group/

给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

示例 1：
	输入：head = [1,2,3,4,5], k = 2
	输出：[2,1,4,3,5]
*/

func ReverseKGroup(head *ListNode, k int) *ListNode {
	dump := &ListNode{Next: head}
	cur := dump
	for cur.Next != nil {
		start, end := cur.Next, cur
		// 找到该组尾节点
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		next := end.Next
		// 截断
		end.Next = nil
		cur.Next = Reverse(start)
		start.Next = next
		cur = start
	}
	return dump.Next
}
