package com.toaction.hot100.linkedList;

public class ReverseNodesInKGroup {

    /*
    https://leetcode.cn/problems/reverse-nodes-in-k-group/
    给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
    * */

    public static ListNode reverseKGroup(ListNode head, int k) {
        ListNode dump = new ListNode(0, head);
        ListNode ahead = dump;
        while(ahead != null) {
            ahead = reverse(ahead, k);
        }
        return dump.next;
    }


    private static ListNode reverse(ListNode ahead, int k) {
        if (ahead.next == null) {
            return null;
        }
        ListNode dump = null;
        ListNode cur = ahead.next;
        // 要返回的节点，作为下一次遍历的头节点
        ListNode tail = cur;
        while(cur!=null && k>0) {
            ListNode nxt = cur.next;
            cur.next = dump;
            dump = cur;
            cur = nxt;
            k--;
        }
        if (k == 0) {
            tail.next = cur;
            ahead.next = dump;
            return tail;
        }
        // 不足 k 个，反转回去
        ahead.next = null;
        while(dump != null) {
            ListNode nxt = dump.next;
            dump.next = ahead.next;
            ahead.next = dump;
            dump = nxt;
        }
        return null;
    }

}
