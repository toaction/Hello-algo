package linkedlist

/*
141. 环形链表: https://leetcode.cn/problems/linked-list-cycle/

给你一个链表的头节点 head ，判断链表中是否有环。

示例
	输入：head = [3,2,0,-4], pos = 1
	输出：true
	解释：链表中有一个环，其尾部连接到第二个节点。
*/

// 思路：快慢指针判断是否有环
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
