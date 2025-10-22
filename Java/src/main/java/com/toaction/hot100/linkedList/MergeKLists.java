package com.toaction.hot100.linkedList;

public class MergeKLists {

    /*
    https://leetcode.cn/problems/merge-k-sorted-lists/
    给你一个链表数组，每个链表都已经按升序排列。
    请你将所有链表合并到一个升序链表中，返回合并后的链表。
    * */

    public static ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0) return null;
        return mergeLists(lists, 0, lists.length-1);
    }

    private static ListNode mergeLists(ListNode[] lists, int left, int right) {
        if (left == right) {
            return lists[left];
        }
        int mid = left + (right-left)/2;
        ListNode head1 = mergeLists(lists, left, mid);
        ListNode head2 = mergeLists(lists, mid+1, right);
        return merge(head1, head2);
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
