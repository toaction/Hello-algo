package com.toaction.hot100.linkedList;

public class RemoveNthNodeFromEnd {

    /*
    https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
    给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
    * */

    // 双指针
    public static ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode dump = new ListNode(0, head);
        // 指向要删除节点的前一个元素
        ListNode cur = dump;
        // 先前进 n 步
        ListNode fast = dump;
        for(int i=0; i<n; i++) {
            fast = fast.next;
        }
        while(fast.next != null) {
            cur = cur.next;
            fast = fast.next;
        }
        // 删除节点
        ListNode nxt = cur.next.next;
        cur.next = nxt;
        return dump.next;
    }
}
