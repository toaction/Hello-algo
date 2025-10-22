package com.toaction.hot100.linkedList;

public class ReverseList {

    /*
     https://leetcode.cn/problems/reverse-linked-list
     给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
     */


    /**
     头插法反转链表
     */
    public static ListNode reverseList(ListNode head) {
        ListNode dump = new ListNode();
        while(head!=null){
            ListNode node = dump.next;
            dump.next = head;
            head = head.next;
            dump.next.next = node;
        }
        return dump.next;
    }
}
