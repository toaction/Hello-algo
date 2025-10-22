package com.toaction.hot100.linkedList;

public class PalindromeList {

    /*
    https://leetcode.cn/problems/palindrome-linked-list/
    给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
    * */

    // 先找到中间节点，然后将 [middle, end] 反转，遍历
    public static boolean isPalindrome(ListNode head) {
        ListNode mid = getMiddle(head);
        // [middle, ]和(middle, ]效果等同
        ListNode cur = reverse(mid.next);
        while(cur!=null && head!=null) {
            if (cur.val != head.val) {
                return false;
            }
            cur = cur.next;
            head = head.next;
        }
        return true;
    }

    // 奇数返回中间，偶数返回中间前一个
    private static ListNode getMiddle(ListNode head) {
        ListNode slow = head;
        ListNode fast = head;
        while(fast.next!=null && fast.next.next!=null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        return slow;
    }

    private static ListNode reverse(ListNode head) {
        ListNode dump = null;
        while(head != null) {
            ListNode nxt = head.next;
            head.next = dump;
            dump = head;
            head = nxt;
        }
        return dump;
    }
}
