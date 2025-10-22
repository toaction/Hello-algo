package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class ReverseNodesInKGroupTest {

    @Test
    public void testReverseKGroupNormal() {
        // Test case: [1,2,3,4,5], k=2 -> [2,1,4,3,5]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        head.next.next.next.next = new ListNode(5);

        ListNode result = ReverseNodesInKGroup.reverseKGroup(head, 2);

        assertEquals(2, result.val);
        assertEquals(1, result.next.val);
        assertEquals(4, result.next.next.val);
        assertEquals(3, result.next.next.next.val);
        assertEquals(5, result.next.next.next.next.val);
        assertNull(result.next.next.next.next.next);
    }

    @Test
    public void testReverseKGroupK3() {
        // Test case: [1,2,3,4,5], k=3 -> [3,2,1,4,5]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        head.next.next.next.next = new ListNode(5);

        ListNode result = ReverseNodesInKGroup.reverseKGroup(head, 3);

        assertEquals(3, result.val);
        assertEquals(2, result.next.val);
        assertEquals(1, result.next.next.val);
        assertEquals(4, result.next.next.next.val);
        assertEquals(5, result.next.next.next.next.val);
        assertNull(result.next.next.next.next.next);
    }

    @Test
    public void testReverseKGroupK1() {
        // Test case: [1,2,3], k=1 -> [1,2,3] (no change)
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);

        ListNode result = ReverseNodesInKGroup.reverseKGroup(head, 1);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertEquals(3, result.next.next.val);
        assertNull(result.next.next.next);
    }

    @Test
    public void testReverseKGroupExactMultiple() {
        // Test case: [1,2,3,4], k=2 -> [2,1,4,3]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);

        ListNode result = ReverseNodesInKGroup.reverseKGroup(head, 2);

        assertEquals(2, result.val);
        assertEquals(1, result.next.val);
        assertEquals(4, result.next.next.val);
        assertEquals(3, result.next.next.next.val);
        assertNull(result.next.next.next.next);
    }

    @Test
    public void testReverseKGroupSingleNode() {
        // Test case: [1], k=1 -> [1]
        ListNode head = new ListNode(1);

        ListNode result = ReverseNodesInKGroup.reverseKGroup(head, 1);

        assertEquals(1, result.val);
        assertNull(result.next);
    }

    @Test
    public void testReverseKGroupKEqualsLength() {
        // Test case: [1,2,3,4,5], k=5 -> [5,4,3,2,1]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        head.next.next.next.next = new ListNode(5);

        ListNode result = ReverseNodesInKGroup.reverseKGroup(head, 5);

        assertEquals(5, result.val);
        assertEquals(4, result.next.val);
        assertEquals(3, result.next.next.val);
        assertEquals(2, result.next.next.next.val);
        assertEquals(1, result.next.next.next.next.val);
        assertNull(result.next.next.next.next.next);
    }

    @Test
    public void testReverseKGroupKGreaterThanLength() {
        // Test case: [1,2,3], k=4 -> [1,2,3] (no change, k > length)
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);

        ListNode result = ReverseNodesInKGroup.reverseKGroup(head, 4);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertEquals(3, result.next.next.val);
        assertNull(result.next.next.next);
    }

    @Test
    public void testReverseKGroupTwoNodes() {
        // Test case: [1,2], k=2 -> [2,1]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);

        ListNode result = ReverseNodesInKGroup.reverseKGroup(head, 2);

        assertEquals(2, result.val);
        assertEquals(1, result.next.val);
        assertNull(result.next.next);
    }
}
