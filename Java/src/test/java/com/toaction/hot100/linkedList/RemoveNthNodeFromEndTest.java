package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class RemoveNthNodeFromEndTest {

    @Test
    public void testRemoveNthFromEndNormal() {
        // Test case: [1,2,3,4,5], n=2 -> [1,2,3,5]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        head.next.next.next.next = new ListNode(5);

        ListNode result = RemoveNthNodeFromEnd.removeNthFromEnd(head, 2);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertEquals(3, result.next.next.val);
        assertEquals(5, result.next.next.next.val);
        assertNull(result.next.next.next.next);
    }

    @Test
    public void testRemoveNthFromEndSingleNode() {
        // Test case: [1], n=1 -> []
        ListNode head = new ListNode(1);

        ListNode result = RemoveNthNodeFromEnd.removeNthFromEnd(head, 1);

        assertNull(result, "Removing only node should return null");
    }

    @Test
    public void testRemoveNthFromEndFirstNode() {
        // Test case: [1,2], n=2 -> [2]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);

        ListNode result = RemoveNthNodeFromEnd.removeNthFromEnd(head, 2);

        assertEquals(2, result.val);
        assertNull(result.next);
    }

    @Test
    public void testRemoveNthFromEndLastNode() {
        // Test case: [1,2,3], n=1 -> [1,2]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);

        ListNode result = RemoveNthNodeFromEnd.removeNthFromEnd(head, 1);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertNull(result.next.next);
    }

    @Test
    public void testRemoveNthFromEndTwoNodes() {
        // Test case: [1,2], n=1 -> [1]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);

        ListNode result = RemoveNthNodeFromEnd.removeNthFromEnd(head, 1);

        assertEquals(1, result.val);
        assertNull(result.next);
    }

    @Test
    public void testRemoveNthFromEndMiddleNode() {
        // Test case: [1,2,3,4,5], n=3 -> [1,2,4,5]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        head.next.next.next.next = new ListNode(5);

        ListNode result = RemoveNthNodeFromEnd.removeNthFromEnd(head, 3);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertEquals(4, result.next.next.val);
        assertEquals(5, result.next.next.next.val);
        assertNull(result.next.next.next.next);
    }
}
