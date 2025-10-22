package com.toaction.hot100.linkedList;

public class SortList {

    /*
    https://leetcode.cn/problems/sort-list/
    给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
    * */

    // 归并排序
    public static ListNode sortList(ListNode head) {
        if (head==null || head.next==null) {
            return head;
        }
        // 先找到中心节点
        ListNode mid = findMiddle(head);
        // 两部分分开
        ListNode head1 = head;
        ListNode head2 = mid.next;
        mid.next = null;
        // 排序
        head1 = sortList(head1);
        head2 = sortList(head2);
        // 合并
        return merge(head1, head2);
    }


    private static ListNode findMiddle(ListNode head) {
        ListNode slow = head;
        ListNode fast = head;
        while(fast.next!=null && fast.next.next!=null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        return slow;
    }

    private static ListNode merge(ListNode head1, ListNode head2) {
        ListNode dump = new ListNode();
        ListNode cur = dump;
        while(head1!=null && head2!=null) {
            if (head1.val < head2.val) {
                cur.next = head1;
                head1 = head1.next;
            }else {
                cur.next = head2;
                head2 = head2.next;
            }
            cur = cur.next;
        }
        cur.next = head1==null?head2:head1;
        return dump.next;
    }
}
