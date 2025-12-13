package linkedlist

/*
160. 相交链表：https://leetcode.cn/problems/intersection-of-two-linked-lists/

给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
如果两个链表不存在相交节点，返回 null 。
*/

// 两个链表长度分别为 a+c, b+c
// 双指针同时移动 pa -> a+c+b, pb -> b+c+a, 正好到达相交节点
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
