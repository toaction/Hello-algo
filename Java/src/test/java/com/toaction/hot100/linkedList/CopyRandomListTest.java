package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

public class CopyRandomListTest {

    @Test
    public void testCopyRandomListNormal() {
        // Create list: [[7,null],[13,0],[11,4],[10,2],[1,0]]
        RandomNode node1 = new RandomNode(7);
        RandomNode node2 = new RandomNode(13);
        RandomNode node3 = new RandomNode(11);
        RandomNode node4 = new RandomNode(10);
        RandomNode node5 = new RandomNode(1);

        node1.next = node2;
        node2.next = node3;
        node3.next = node4;
        node4.next = node5;

        node1.random = null;
        node2.random = node1;
        node3.random = node5;
        node4.random = node3;
        node5.random = node1;

        RandomNode result = CopyRandomList.copyRandomList(node1);

        // Verify deep copy
        assertNotSame(node1, result, "Should be a different object");

        // Verify values
        assertEquals(7, result.val);
        assertEquals(13, result.next.val);
        assertEquals(11, result.next.next.val);
        assertEquals(10, result.next.next.next.val);
        assertEquals(1, result.next.next.next.next.val);

        // Verify random pointers
        assertNull(result.random);
        assertEquals(7, result.next.random.val);
        assertEquals(1, result.next.next.random.val);
        assertEquals(11, result.next.next.next.random.val);
        assertEquals(7, result.next.next.next.next.random.val);

        // Verify original list is not modified
        assertEquals(7, node1.val);
        assertNull(node1.random);
    }

    @Test
    public void testCopyRandomListSingleNode() {
        // Create list: [[1,null]]
        RandomNode node = new RandomNode(1);
        node.next = null;
        node.random = null;

        RandomNode result = CopyRandomList.copyRandomList(node);

        assertNotSame(node, result, "Should be a different object");
        assertEquals(1, result.val);
        assertNull(result.next);
        assertNull(result.random);
    }

    @Test
    public void testCopyRandomListWithSelfPointing() {
        // Create list where node points to itself: [[1,0],[2,0]]
        RandomNode node1 = new RandomNode(1);
        RandomNode node2 = new RandomNode(2);

        node1.next = node2;
        node1.random = node1;
        node2.random = node1;

        RandomNode result = CopyRandomList.copyRandomList(node1);

        assertNotSame(node1, result);
        assertEquals(1, result.val);
        assertEquals(2, result.next.val);
        assertSame(result, result.random, "First node should point to itself");
        assertSame(result, result.next.random, "Second node should point to first");
    }

    @Test
    public void testCopyRandomListEmpty() {
        // Test with null input
        RandomNode result = CopyRandomList.copyRandomList(null);
        assertNull(result);
    }
}
