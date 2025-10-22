package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.junit.jupiter.api.Assertions.assertFalse;

public class PalindromeListTest {

    @Test
    public void testIsPalindromeTrue() {
        // Create a palindrome linked list: 1 -> 2 -> 2 -> 1
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(2);
        head.next.next.next = new ListNode(1);

        boolean result = PalindromeList.isPalindrome(head);
        assertTrue(result, "Should return true for palindrome list [1,2,2,1]");
    }

    @Test
    public void testIsPalindromeFalse() {
        // Create a non-palindrome linked list: 1 -> 2 -> 3
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);

        boolean result = PalindromeList.isPalindrome(head);
        assertFalse(result, "Should return false for non-palindrome list [1,2,3]");
    }
}
