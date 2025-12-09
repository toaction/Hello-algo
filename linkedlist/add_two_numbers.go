package linkedlist

/* 
2. 两数相加：https://leetcode.cn/problems/add-two-numbers/

给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
	输入：l1 = [2,4,3], l2 = [5,6,4]
	输出：[7,0,8]
	解释：342 + 465 = 807 
*/

// 模拟计算
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dump := &ListNode{}
	cur := dump
	carry := 0
	// 共有的低位
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		cur.Next = &ListNode{
			Val: sum % 10,
		}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	// 剩余的高位
	high := l1
	if l2 != nil {
		high = l2
	}
	for high != nil {
		sum := high.Val + carry
		carry = sum / 10
		cur.Next = &ListNode{
			Val: sum % 10,
		}
		cur = cur.Next
		high = high.Next
	}
	// 进位
	if carry != 0 {
		cur.Next = &ListNode{
			Val: carry,
		}
	}
	return dump.Next
}
