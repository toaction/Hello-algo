package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class ReverseListTest {

    @Test
    public void testReverseListNormal() {
        // Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        head.next.next.next.next = new ListNode(5);

        ListNode result = ReverseList.reverseList(head);

        // Check reversed list: 5 -> 4 -> 3 -> 2 -> 1
        assertEquals(5, result.val, "First node should be 5");
        assertEquals(4, result.next.val, "Second node should be 4");
        assertEquals(3, result.next.next.val, "Third node should be 3");
        assertEquals(2, result.next.next.next.val, "Fourth node should be 2");
        assertEquals(1, result.next.next.next.next.val, "Fifth node should be 1");
        assertNull(result.next.next.next.next.next, "Last node should point to null");
    }

    @Test
    public void testReverseListSingleNode() {
        // Create a linked list with single node: 1
        ListNode head = new ListNode(1);

        ListNode result = ReverseList.reverseList(head);

        // Check result
        assertEquals(1, result.val, "Single node value should remain 1");
        assertNull(result.next, "Single node should point to null");
    }
}
