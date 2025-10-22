package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.junit.jupiter.api.Assertions.assertFalse;

public class LinkedListCycleOneTest {

    @Test
    public void testHasCycleWithCycle() {
        // Create a linked list with cycle: 3 -> 2 -> 0 -> -4 -> (back to 2)
        ListNode head = new ListNode(3);
        ListNode node2 = new ListNode(2);
        ListNode node0 = new ListNode(0);
        ListNode node4 = new ListNode(-4);

        head.next = node2;
        node2.next = node0;
        node0.next = node4;
        node4.next = node2; // Create cycle

        boolean result = LinkedListCycleOne.hasCycle(head);
        assertTrue(result, "Should return true when the linked list has a cycle");
    }

    @Test
    public void testHasCycleNoCycle() {
        // Create a linked list without cycle: 1 -> 2 -> 3 -> 4
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);

        boolean result = LinkedListCycleOne.hasCycle(head);
        assertFalse(result, "Should return false when the linked list has no cycle");
    }
}
