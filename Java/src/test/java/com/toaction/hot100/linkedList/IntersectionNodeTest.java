package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class IntersectionNodeTest {

    @Test
    public void testGetIntersectionNode1WithIntersection() {
        // Create two intersecting linked lists
        // List A: 4 -> 1 -> 8 -> 4 -> 5
        // List B: 5 -> 6 -> 1 -> 8 -> 4 -> 5
        // Intersection at node with value 8
        ListNode intersect = new ListNode(8);
        intersect.next = new ListNode(4);
        intersect.next.next = new ListNode(5);

        ListNode headA = new ListNode(4);
        headA.next = new ListNode(1);
        headA.next.next = intersect;

        ListNode headB = new ListNode(5);
        headB.next = new ListNode(6);
        headB.next.next = new ListNode(1);
        headB.next.next.next = intersect;

        ListNode result = IntersectionNode.getIntersectionNode1(headA, headB);
        assertEquals(intersect, result, "Should return the intersection node");
    }

    @Test
    public void testGetIntersectionNode1NoIntersection() {
        // Create two non-intersecting linked lists
        // List A: 2 -> 6 -> 4
        // List B: 1 -> 5
        ListNode headA = new ListNode(2);
        headA.next = new ListNode(6);
        headA.next.next = new ListNode(4);

        ListNode headB = new ListNode(1);
        headB.next = new ListNode(5);

        ListNode result = IntersectionNode.getIntersectionNode1(headA, headB);
        assertNull(result, "Should return null when there is no intersection");
    }

    @Test
    public void testGetIntersectionNode2WithIntersection() {
        // Create two intersecting linked lists
        // List A: 1 -> 9 -> 1 -> 2 -> 4
        // List B: 3 -> 2 -> 4
        // Intersection at node with value 2
        ListNode intersect = new ListNode(2);
        intersect.next = new ListNode(4);

        ListNode headA = new ListNode(1);
        headA.next = new ListNode(9);
        headA.next.next = new ListNode(1);
        headA.next.next.next = intersect;

        ListNode headB = new ListNode(3);
        headB.next = intersect;

        ListNode result = IntersectionNode.getIntersectionNode2(headA, headB);
        assertEquals(intersect, result, "Should return the intersection node");
    }

    @Test
    public void testGetIntersectionNode2NoIntersection() {
        // Create two non-intersecting linked lists
        // List A: 1 -> 2 -> 3
        // List B: 4 -> 5 -> 6
        ListNode headA = new ListNode(1);
        headA.next = new ListNode(2);
        headA.next.next = new ListNode(3);

        ListNode headB = new ListNode(4);
        headB.next = new ListNode(5);
        headB.next.next = new ListNode(6);

        ListNode result = IntersectionNode.getIntersectionNode2(headA, headB);
        assertNull(result, "Should return null when there is no intersection");
    }
}
