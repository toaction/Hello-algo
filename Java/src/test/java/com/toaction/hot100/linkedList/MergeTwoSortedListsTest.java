package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class MergeTwoSortedListsTest {

    @Test
    public void testMergeTwoSortedListsNormal() {
        // Test case: [1,2,4] and [1,3,4] -> [1,1,2,3,4,4]
        ListNode list1 = new ListNode(1);
        list1.next = new ListNode(2);
        list1.next.next = new ListNode(4);

        ListNode list2 = new ListNode(1);
        list2.next = new ListNode(3);
        list2.next.next = new ListNode(4);

        ListNode result = MergeTwoSortedLists.mergeTwoLists(list1, list2);

        assertEquals(1, result.val);
        assertEquals(1, result.next.val);
        assertEquals(2, result.next.next.val);
        assertEquals(3, result.next.next.next.val);
        assertEquals(4, result.next.next.next.next.val);
        assertEquals(4, result.next.next.next.next.next.val);
        assertNull(result.next.next.next.next.next.next);
    }

    @Test
    public void testMergeTwoSortedListsBothEmpty() {
        // Test case: [] and [] -> []
        ListNode result = MergeTwoSortedLists.mergeTwoLists(null, null);
        assertNull(result, "Both lists empty should return null");
    }

    @Test
    public void testMergeTwoSortedListsOneEmpty() {
        // Test case: [] and [0] -> [0]
        ListNode list2 = new ListNode(0);
        ListNode result = MergeTwoSortedLists.mergeTwoLists(null, list2);

        assertEquals(0, result.val);
        assertNull(result.next);
    }

    @Test
    public void testMergeTwoSortedListsDifferentLengths() {
        // Test case: [1,3,5,7] and [2,4] -> [1,2,3,4,5,7]
        ListNode list1 = new ListNode(1);
        list1.next = new ListNode(3);
        list1.next.next = new ListNode(5);
        list1.next.next.next = new ListNode(7);

        ListNode list2 = new ListNode(2);
        list2.next = new ListNode(4);

        ListNode result = MergeTwoSortedLists.mergeTwoLists(list1, list2);

        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertEquals(3, result.next.next.val);
        assertEquals(4, result.next.next.next.val);
        assertEquals(5, result.next.next.next.next.val);
        assertEquals(7, result.next.next.next.next.next.val);
        assertNull(result.next.next.next.next.next.next);
    }

    @Test
    public void testMergeTwoSortedListsAllSameValues() {
        // Test case: [1,1,1] and [1,1,1] -> [1,1,1,1,1,1]
        ListNode list1 = new ListNode(1);
        list1.next = new ListNode(1);
        list1.next.next = new ListNode(1);

        ListNode list2 = new ListNode(1);
        list2.next = new ListNode(1);
        list2.next.next = new ListNode(1);

        ListNode result = MergeTwoSortedLists.mergeTwoLists(list1, list2);

        for (int i = 0; i < 6; i++) {
            assertEquals(1, result.val, "All values should be 1");
            if (i < 5) {
                result = result.next;
            }
        }
        assertNull(result.next);
    }

    @Test
    public void testMergeTwoSortedListsNegativeNumbers() {
        // Test case: [-3,-1,0] and [-2,1] -> [-3,-2,-1,0,1]
        ListNode list1 = new ListNode(-3);
        list1.next = new ListNode(-1);
        list1.next.next = new ListNode(0);

        ListNode list2 = new ListNode(-2);
        list2.next = new ListNode(1);

        ListNode result = MergeTwoSortedLists.mergeTwoLists(list1, list2);

        assertEquals(-3, result.val);
        assertEquals(-2, result.next.val);
        assertEquals(-1, result.next.next.val);
        assertEquals(0, result.next.next.next.val);
        assertEquals(1, result.next.next.next.next.val);
        assertNull(result.next.next.next.next.next);
    }
}
