package linkedlist

/*
142. 环形链表 II：https://leetcode.cn/problems/linked-list-cycle-ii/

给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
*/

// 快慢指针找到相遇节点
// 相遇时，慢指针走过 a+b, 快指针走过 a+k(b+c)+b,
// 此时  2(a+b)=a+b+k(b+c) -> a=(k-1)(b+c)+c
// 相遇后，一个指针指向head，另一个指向相遇节点，同时移动a距离再次相遇，恰好到达入环节点
func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	// 找到相遇节点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
