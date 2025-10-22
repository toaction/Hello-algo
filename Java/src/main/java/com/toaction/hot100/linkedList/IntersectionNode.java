package com.toaction.hot100.linkedList;

public class IntersectionNode {

    /* 
     * https://leetcode.cn/problems/intersection-of-two-linked-lists/
     * 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
     */


    /**
     * 先求出各自长度，长的先前进 k=m-n 步，然后共同前进，找到相同节点
     */
    public static ListNode getIntersectionNode1(ListNode headA, ListNode headB) {
        int lenA = 0;
        ListNode curA = headA;
        while(curA != null) {
            lenA++;
            curA = curA.next;
        }
        int lenB = 0;
        ListNode curB = headB;
        while(curB != null) {
            lenB++;
            curB = curB.next;
        }
        curA = headA;
        curB = headB;
        if (lenA > lenB) {
            int diff = lenA - lenB;
            while(diff > 0) {
                curA = curA.next;
                diff--;
            }
        }else {
            int diff = lenB - lenA;
            while(diff > 0) {
                curB = curB.next;
                diff--;
            }
        }
        while(curA != curB) {
            curA = curA.next;
            curB = curB.next;
        }
        return curA;
    }


    /**
     * 双指针
     * 指针A先遍历headA，遍历结束后遍历headB;
     * 指针B先遍历headB，遍历结束后遍历headA;
     * 两指针走过的长度为(a+b-c)和(b+a-c)，c为相交节点到尾部的长度
     */
    public static ListNode getIntersectionNode2(ListNode headA, ListNode headB) {
        if (headA == null || headB == null) {
            return null;
        }
        ListNode pA = headA, pB = headB;
        while (pA != pB) {
            pA = pA == null ? headB : pA.next;
            pB = pB == null ? headA : pB.next;
        }
        return pA;
    }

}
