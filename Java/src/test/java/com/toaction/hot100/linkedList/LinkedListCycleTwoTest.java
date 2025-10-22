package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class LinkedListCycleTwoTest {

    @Test
    public void testDetectCycleWithCycle() {
        // Create a linked list with cycle: 3 -> 2 -> 0 -> -4 -> (back to 2)
        ListNode head = new ListNode(3);
        ListNode node2 = new ListNode(2);
        ListNode node0 = new ListNode(0);
        ListNode node4 = new ListNode(-4);

        head.next = node2;
        node2.next = node0;
        node0.next = node4;
        node4.next = node2; // Create cycle, entry point is node2

        ListNode result = LinkedListCycleTwo.detectCycle(head);
        assertEquals(node2, result, "Should return the node where the cycle begins");
    }

    @Test
    public void testDetectCycleNoCycle() {
        // Create a linked list without cycle: 1 -> 2 -> 3 -> 4 -> null
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);

        ListNode result = LinkedListCycleTwo.detectCycle(head);
        assertNull(result, "Should return null when there is no cycle");
    }
}
