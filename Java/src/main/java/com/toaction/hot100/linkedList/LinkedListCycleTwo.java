package com.toaction.hot100.linkedList;

public class LinkedListCycleTwo {

    /*
    * https://leetcode.cn/problems/linked-list-cycle-ii/
    * 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
    * */


    /**
     * 使用快慢指针，快指针相对于慢指针速度快1个单位，
     * 当慢指针进入环后，两者距离最大为n-1，因此快指针一定会在慢指针走完第一圈之前追上
     * 当两者相遇时，慢指针走过的距离：a+b，快指针：a+b+k(b+c), 此时有2(a+b) = a+b+k(b+c)
     * 整理得 a = (k-1)(b+c)+c，此时再让慢指针走a步，相对环入口位置为b+(k-1)(b+c)+c = 0
     */
    public static ListNode detectCycle(ListNode head) {
        ListNode meet = meetCycle(head);
        if (meet == null) return null;
        while(head != meet) {
            head = head.next;
            meet = meet.next;
        }
        return head;
    }


    private static ListNode meetCycle(ListNode head) {
        if (head == null) return null;
        ListNode slow = head;
        ListNode fast = head;
        while(fast.next!=null && fast.next.next!=null) {
            slow = slow.next;
            fast = fast.next.next;
            if (slow == fast) {
                return slow;
            }
        }
        return null;
    }
}
