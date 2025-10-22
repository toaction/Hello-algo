package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class SwapNodesInPairsTest {

    @Test
    public void testSwapPairsNormal() {
        // Test case: [1,2,3,4] -> [2,1,4,3]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);

        ListNode result = SwapNodesInPairs.swapPairs(head);

        assertEquals(2, result.val);
        assertEquals(1, result.next.val);
        assertEquals(4, result.next.next.val);
        assertEquals(3, result.next.next.next.val);
        assertNull(result.next.next.next.next);
    }

    @Test
    public void testSwapPairsEmpty() {
        // Test case: [] -> []
        ListNode result = SwapNodesInPairs.swapPairs(null);
        assertNull(result, "Empty list should return null");
    }

    @Test
    public void testSwapPairsSingleNode() {
        // Test case: [1] -> [1]
        ListNode head = new ListNode(1);

        ListNode result = SwapNodesInPairs.swapPairs(head);

        assertEquals(1, result.val);
        assertNull(result.next);
    }

    @Test
    public void testSwapPairsOddLength() {
        // Test case: [1,2,3,4,5] -> [2,1,4,3,5]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        head.next.next.next.next = new ListNode(5);

        ListNode result = SwapNodesInPairs.swapPairs(head);

        assertEquals(2, result.val);
        assertEquals(1, result.next.val);
        assertEquals(4, result.next.next.val);
        assertEquals(3, result.next.next.next.val);
        assertEquals(5, result.next.next.next.next.val);
        assertNull(result.next.next.next.next.next);
    }

    @Test
    public void testSwapPairsTwoNodes() {
        // Test case: [1,2] -> [2,1]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);

        ListNode result = SwapNodesInPairs.swapPairs(head);

        assertEquals(2, result.val);
        assertEquals(1, result.next.val);
        assertNull(result.next.next);
    }

    @Test
    public void testSwapPairsThreeNodes() {
        // Test case: [1,2,3] -> [2,1,3]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);

        ListNode result = SwapNodesInPairs.swapPairs(head);

        assertEquals(2, result.val);
        assertEquals(1, result.next.val);
        assertEquals(3, result.next.next.val);
        assertNull(result.next.next.next);
    }

    @Test
    public void testSwapPairsLongerList() {
        // Test case: [1,2,3,4,5,6] -> [2,1,4,3,6,5]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        head.next.next.next.next = new ListNode(5);
        head.next.next.next.next.next = new ListNode(6);

        ListNode result = SwapNodesInPairs.swapPairs(head);

        assertEquals(2, result.val);
        assertEquals(1, result.next.val);
        assertEquals(4, result.next.next.val);
        assertEquals(3, result.next.next.next.val);
        assertEquals(6, result.next.next.next.next.val);
        assertEquals(5, result.next.next.next.next.next.val);
        assertNull(result.next.next.next.next.next.next);
    }
}
