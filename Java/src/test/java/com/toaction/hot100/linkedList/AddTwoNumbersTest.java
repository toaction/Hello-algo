package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class AddTwoNumbersTest {

    @Test
    public void testAddTwoNumbersNormal() {
        // Test case: 342 + 465 = 807 (represented as 2->4->3 + 5->6->4 = 7->0->8)
        ListNode l1 = new ListNode(2);
        l1.next = new ListNode(4);
        l1.next.next = new ListNode(3);

        ListNode l2 = new ListNode(5);
        l2.next = new ListNode(6);
        l2.next.next = new ListNode(4);

        ListNode result = AddTwoNumbers.addTwoNumbers(l1, l2);

        assertEquals(7, result.val, "First digit should be 7");
        assertEquals(0, result.next.val, "Second digit should be 0");
        assertEquals(8, result.next.next.val, "Third digit should be 8");
        assertNull(result.next.next.next, "Should have no more digits");
    }

    @Test
    public void testAddTwoNumbersWithCarry() {
        // Test case: 9999999 + 9999 = 10009998 (represented as linked lists)
        ListNode l1 = new ListNode(9);
        l1.next = new ListNode(9);
        l1.next.next = new ListNode(9);
        l1.next.next.next = new ListNode(9);
        l1.next.next.next.next = new ListNode(9);
        l1.next.next.next.next.next = new ListNode(9);
        l1.next.next.next.next.next.next = new ListNode(9);

        ListNode l2 = new ListNode(9);
        l2.next = new ListNode(9);
        l2.next.next = new ListNode(9);
        l2.next.next.next = new ListNode(9);

        ListNode result = AddTwoNumbers.addTwoNumbers(l1, l2);

        assertEquals(8, result.val);
        assertEquals(9, result.next.val);
        assertEquals(9, result.next.next.val);
        assertEquals(9, result.next.next.next.val);
        assertEquals(0, result.next.next.next.next.val);
        assertEquals(0, result.next.next.next.next.next.val);
        assertEquals(0, result.next.next.next.next.next.next.val);
        assertEquals(1, result.next.next.next.next.next.next.next.val);
    }

    @Test
    public void testAddTwoNumbersZero() {
        // Test case: 0 + 0 = 0
        ListNode l1 = new ListNode(0);
        ListNode l2 = new ListNode(0);

        ListNode result = AddTwoNumbers.addTwoNumbers(l1, l2);

        assertEquals(0, result.val, "Result should be 0");
        assertNull(result.next, "Should have only one digit");
    }

    @Test
    public void testAddTwoNumbersDifferentLengths() {
        // Test case: 99 + 1 = 100 (represented as 9->9 + 1 = 0->0->1)
        ListNode l1 = new ListNode(9);
        l1.next = new ListNode(9);

        ListNode l2 = new ListNode(1);

        ListNode result = AddTwoNumbers.addTwoNumbers(l1, l2);

        assertEquals(0, result.val, "First digit should be 0");
        assertEquals(0, result.next.val, "Second digit should be 0");
        assertEquals(1, result.next.next.val, "Third digit should be 1");
        assertNull(result.next.next.next, "Should have no more digits");
    }
}
