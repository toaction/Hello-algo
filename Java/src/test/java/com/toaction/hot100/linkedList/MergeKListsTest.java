package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class MergeKListsTest {

    @Test
    public void testMergeKListsNormal() {
        // Create lists: [[1,4,5],[1,3,4],[2,6]]
        ListNode list1 = new ListNode(1);
        list1.next = new ListNode(4);
        list1.next.next = new ListNode(5);

        ListNode list2 = new ListNode(1);
        list2.next = new ListNode(3);
        list2.next.next = new ListNode(4);

        ListNode list3 = new ListNode(2);
        list3.next = new ListNode(6);

        ListNode[] lists = {list1, list2, list3};
        ListNode result = MergeKLists.mergeKLists(lists);

        // Expected result: [1,1,2,3,4,4,5,6]
        assertEquals(1, result.val);
        assertEquals(1, result.next.val);
        assertEquals(2, result.next.next.val);
        assertEquals(3, result.next.next.next.val);
        assertEquals(4, result.next.next.next.next.val);
        assertEquals(4, result.next.next.next.next.next.val);
        assertEquals(5, result.next.next.next.next.next.next.val);
        assertEquals(6, result.next.next.next.next.next.next.next.val);
        assertNull(result.next.next.next.next.next.next.next.next);
    }

    @Test
    public void testMergeKListsEmpty() {
        // Test with empty array
        ListNode[] lists = {};
        ListNode result = MergeKLists.mergeKLists(lists);
        assertNull(result, "Should return null for empty array");
    }

    @Test
    public void testMergeKListsSingleList() {
        // Test with single list: [[1,2,3]]
        ListNode list1 = new ListNode(1);
        list1.next = new ListNode(2);
        list1.next.next = new ListNode(3);

        ListNode[] lists = {list1};
        ListNode result = MergeKLists.mergeKLists(lists);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertEquals(3, result.next.next.val);
        assertNull(result.next.next.next);
    }

    @Test
    public void testMergeKListsWithNulls() {
        // Test with some null lists: [[1,2],[],null,[3,4]]
        ListNode list1 = new ListNode(1);
        list1.next = new ListNode(2);

        ListNode list3 = new ListNode(3);
        list3.next = new ListNode(4);

        ListNode[] lists = {list1, null, list3};
        ListNode result = MergeKLists.mergeKLists(lists);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertEquals(3, result.next.next.val);
        assertEquals(4, result.next.next.next.val);
        assertNull(result.next.next.next.next);
    }

    @Test
    public void testMergeKListsSameValues() {
        // Test with lists containing same values: [[1,1,1],[1,1,1]]
        ListNode list1 = new ListNode(1);
        list1.next = new ListNode(1);
        list1.next.next = new ListNode(1);

        ListNode list2 = new ListNode(1);
        list2.next = new ListNode(1);
        list2.next.next = new ListNode(1);

        ListNode[] lists = {list1, list2};
        ListNode result = MergeKLists.mergeKLists(lists);

        for (int i = 0; i < 6; i++) {
            assertEquals(1, result.val, "All values should be 1");
            if (i < 5) {
                result = result.next;
            }
        }
        assertNull(result.next);
    }

    @Test
    public void testMergeKListsNegativeNumbers() {
        // Test with negative numbers: [[-2,0,2],[-1,1]]
        ListNode list1 = new ListNode(-2);
        list1.next = new ListNode(0);
        list1.next.next = new ListNode(2);

        ListNode list2 = new ListNode(-1);
        list2.next = new ListNode(1);

        ListNode[] lists = {list1, list2};
        ListNode result = MergeKLists.mergeKLists(lists);

        assertEquals(-2, result.val);
        assertEquals(-1, result.next.val);
        assertEquals(0, result.next.next.val);
        assertEquals(1, result.next.next.next.val);
        assertEquals(2, result.next.next.next.next.val);
        assertNull(result.next.next.next.next.next);
    }
}
