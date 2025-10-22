package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class SortListTest {

    @Test
    public void testSortListNormal() {
        // Test case: [4,2,1,3] -> [1,2,3,4]
        ListNode head = new ListNode(4);
        head.next = new ListNode(2);
        head.next.next = new ListNode(1);
        head.next.next.next = new ListNode(3);

        ListNode result = SortList.sortList(head);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertEquals(3, result.next.next.val);
        assertEquals(4, result.next.next.next.val);
        assertNull(result.next.next.next.next);
    }

    @Test
    public void testSortListAlreadySorted() {
        // Test case: [1,2,3,4] -> [1,2,3,4]
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);

        ListNode result = SortList.sortList(head);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertEquals(3, result.next.next.val);
        assertEquals(4, result.next.next.next.val);
        assertNull(result.next.next.next.next);
    }

    @Test
    public void testSortListReverseSorted() {
        // Test case: [5,4,3,2,1] -> [1,2,3,4,5]
        ListNode head = new ListNode(5);
        head.next = new ListNode(4);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(2);
        head.next.next.next.next = new ListNode(1);

        ListNode result = SortList.sortList(head);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertEquals(3, result.next.next.val);
        assertEquals(4, result.next.next.next.val);
        assertEquals(5, result.next.next.next.next.val);
        assertNull(result.next.next.next.next.next);
    }

    @Test
    public void testSortListEmpty() {
        // Test case: [] -> []
        ListNode result = SortList.sortList(null);
        assertNull(result, "Empty list should return null");
    }

    @Test
    public void testSortListSingleNode() {
        // Test case: [1] -> [1]
        ListNode head = new ListNode(1);

        ListNode result = SortList.sortList(head);

        assertEquals(1, result.val);
        assertNull(result.next);
    }

    @Test
    public void testSortListTwoNodes() {
        // Test case: [2,1] -> [1,2]
        ListNode head = new ListNode(2);
        head.next = new ListNode(1);

        ListNode result = SortList.sortList(head);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertNull(result.next.next);
    }

    @Test
    public void testSortListDuplicates() {
        // Test case: [3,1,2,1,3] -> [1,1,2,3,3]
        ListNode head = new ListNode(3);
        head.next = new ListNode(1);
        head.next.next = new ListNode(2);
        head.next.next.next = new ListNode(1);
        head.next.next.next.next = new ListNode(3);

        ListNode result = SortList.sortList(head);

        assertEquals(1, result.val);
        assertEquals(1, result.next.val);
        assertEquals(2, result.next.next.val);
        assertEquals(3, result.next.next.next.val);
        assertEquals(3, result.next.next.next.next.val);
        assertNull(result.next.next.next.next.next);
    }

    @Test
    public void testSortListNegativeNumbers() {
        // Test case: [-1,5,3,4,0] -> [-1,0,3,4,5]
        ListNode head = new ListNode(-1);
        head.next = new ListNode(5);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        head.next.next.next.next = new ListNode(0);

        ListNode result = SortList.sortList(head);

        assertEquals(-1, result.val);
        assertEquals(0, result.next.val);
        assertEquals(3, result.next.next.val);
        assertEquals(4, result.next.next.next.val);
        assertEquals(5, result.next.next.next.next.val);
        assertNull(result.next.next.next.next.next);
    }
}
