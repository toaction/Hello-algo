package com.toaction.hot100.linkedList;

public class LinkedListCycleOne {

    /*
    * https://leetcode.cn/problems/linked-list-cycle/
    * 给你一个链表的头节点 head ，判断链表中是否有环。
    * */

    // 快慢指针判断是否存在环
    public static boolean hasCycle(ListNode head) {
        if (head == null) return false;
        ListNode slow = head;
        ListNode fast = head;
        while(fast.next!=null && fast.next.next!=null) {
            slow = slow.next;
            fast = fast.next.next;
            if (slow == fast) {
                return true;
            }
        }
        return false;
    }
}
