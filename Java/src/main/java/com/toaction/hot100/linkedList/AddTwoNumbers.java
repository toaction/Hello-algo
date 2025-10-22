package com.toaction.hot100.linkedList;

public class AddTwoNumbers {

    /*
    https://leetcode.cn/problems/add-two-numbers/
    给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
    请你将两个数相加，并以相同形式返回一个表示和的链表。
    你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
    * */

    public static ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dump = new ListNode();
        ListNode cur = dump;
        // 两数相加产生的进位
        int carry = 0;
        while(l1!=null && l2!=null){
            int sum = l1.val + l2.val + carry;
            cur.next = new ListNode(sum%10);
            carry = sum/10;
            cur = cur.next;
            l1 = l1.next;
            l2 = l2.next;
        }
        // 处理高位
        ListNode high = l1==null?l2:l1;
        while(high!=null){
            int sum = high.val + carry;
            cur.next = new ListNode(sum%10);
            carry = sum/10;
            cur = cur.next;
            high = high.next;
        }
        // 处理进位
        if(carry > 0){
            cur.next = new ListNode(carry);
        }
        return dump.next;
    }
}
