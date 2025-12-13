package linkedlist

/*
234. 回文链表：https://leetcode.cn/problems/palindrome-linked-list/

给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
*/

// 先使用快慢指针找到链表的中间节点
// 将后半部分链表反转
// 双指针同时移动判断 val 是否相同
func IsPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	mid := findMiddle(head)
	pa, pb := head, Reverse(mid)
	for pa != nil && pb != nil {
		if pa.Val != pb.Val {
			return false
		}
		pa = pa.Next
		pb = pb.Next
	}
	return true
}

// 返回的是中间节点（奇数）或后半部分头节点（偶数）
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func Reverse(head *ListNode) *ListNode {
	dump := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = dump.Next
		dump.Next = head
		head = next
	}
	return dump.Next
}
