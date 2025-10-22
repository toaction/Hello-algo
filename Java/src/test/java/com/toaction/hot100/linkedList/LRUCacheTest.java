package com.toaction.hot100.linkedList;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class LRUCacheTest {

    @Test
    public void testLRUCacheBasicOperations() {
        LRUCache cache = new LRUCache(2);

        cache.put(1, 1);
        cache.put(2, 2);
        assertEquals(1, cache.get(1), "Should get value 1 for key 1");

        cache.put(3, 3); // Evicts key 2
        assertEquals(-1, cache.get(2), "Key 2 should be evicted");

        cache.put(4, 4); // Evicts key 1
        assertEquals(-1, cache.get(1), "Key 1 should be evicted");
        assertEquals(3, cache.get(3), "Should get value 3 for key 3");
        assertEquals(4, cache.get(4), "Should get value 4 for key 4");
    }

    @Test
    public void testLRUCacheUpdateValue() {
        LRUCache cache = new LRUCache(2);

        cache.put(1, 1);
        cache.put(2, 2);
        assertEquals(1, cache.get(1), "Should get value 1 for key 1");

        cache.put(1, 10); // Update key 1
        assertEquals(10, cache.get(1), "Should get updated value 10 for key 1");
        assertEquals(2, cache.get(2), "Should still get value 2 for key 2");
    }

    @Test
    public void testLRUCacheSingleCapacity() {
        LRUCache cache = new LRUCache(1);

        cache.put(1, 1);
        assertEquals(1, cache.get(1), "Should get value 1 for key 1");

        cache.put(2, 2); // Evicts key 1
        assertEquals(-1, cache.get(1), "Key 1 should be evicted");
        assertEquals(2, cache.get(2), "Should get value 2 for key 2");
    }

    @Test
    public void testLRUCacheGetNonExistent() {
        LRUCache cache = new LRUCache(2);

        assertEquals(-1, cache.get(1), "Should return -1 for non-existent key");

        cache.put(1, 1);
        assertEquals(1, cache.get(1), "Should get value 1 for key 1");
        assertEquals(-1, cache.get(2), "Should return -1 for non-existent key 2");
    }

    @Test
    public void testLRUCacheAccessOrder() {
        LRUCache cache = new LRUCache(2);

        cache.put(1, 1);
        cache.put(2, 2);
        assertEquals(1, cache.get(1), "Access key 1, making it most recently used");

        cache.put(3, 3); // Should evict key 2 (least recently used)
        assertEquals(-1, cache.get(2), "Key 2 should be evicted");
        assertEquals(1, cache.get(1), "Key 1 should still be in cache");
        assertEquals(3, cache.get(3), "Key 3 should be in cache");
    }

    @Test
    public void testLRUCacheLargerCapacity() {
        LRUCache cache = new LRUCache(3);

        cache.put(1, 1);
        cache.put(2, 2);
        cache.put(3, 3);
        cache.put(4, 4); // Evicts key 1

        assertEquals(-1, cache.get(1), "Key 1 should be evicted");
        assertEquals(2, cache.get(2), "Key 2 should be in cache");
        assertEquals(3, cache.get(3), "Key 3 should be in cache");
        assertEquals(4, cache.get(4), "Key 4 should be in cache");
    }

    @Test
    public void testLRUCachePutAfterGet() {
        LRUCache cache = new LRUCache(2);

        cache.put(2, 1);
        cache.put(1, 1);
        cache.put(2, 3); // Update key 2
        cache.put(4, 1); // Evicts key 1

        assertEquals(-1, cache.get(1), "Key 1 should be evicted");
        assertEquals(3, cache.get(2), "Key 2 should have updated value 3");
    }
}
