package com.toaction.hot100.linkedList;

public class SwapNodesInPairs {

    /*
    https://leetcode.cn/problems/swap-nodes-in-pairs/
    给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
    * */

    public static ListNode swapPairs(ListNode head) {
        ListNode dump = new ListNode(0, head);
        ListNode cur = dump;
        // cur为要交换的节点的前一个节点
        while(cur.next!=null && cur.next.next!=null) {
            ListNode first = cur.next;
            ListNode second = cur.next.next;
            first.next = second.next;
            second.next = first;
            cur.next = second;
            cur = first;
        }
        return dump.next;
    }

}
