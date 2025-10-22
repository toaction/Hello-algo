package com.toaction.hot100.linkedList;

import java.util.HashMap;
import java.util.Map;

public class LRUCache {

    /*
    * https://leetcode.cn/problems/lru-cache/
    * */


    /**
     双向链表存放具体值，Map快速定位元素
     */
    class ListNode{
        int key;
        int val;
        ListNode next;
        ListNode pre;
        public ListNode(){};
        public ListNode(int key, int val){
            this.key = key;
            this.val = val;
            this.next = null;
            this.pre = null;
        }
    }

    private int capacity;
    private int size;
    private Map<Integer, ListNode> cache = new HashMap<>();
    private ListNode head;
    private ListNode tail;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.size = 0;
        this.head = new ListNode(0, 0);
        this.tail = new ListNode(0, 0);
        this.head.next = this.tail;
        this.tail.pre = this.head;
    }

    public int get(int key) {
        ListNode node = cache.getOrDefault(key, null);
        if(node == null){
            return -1;
        }
        // 移动到链表头部
        moveTohead(node);
        return node.val;
    }

    public void put(int key, int value) {
        // key已存在，变更数据值
        ListNode node = cache.getOrDefault(key, null);
        if(node != null){
            node.val = value;
            moveTohead(node);
            return;
        }
        // 已满，删除一个节点
        if(size == capacity){
            ListNode temp = tail.pre;
            deleteNode(temp);
            cache.remove(temp.key);
            size--;
        }
        // 插入新节点
        node = new ListNode(key, value);
        cache.put(key, node);
        insertHead(node);
        size++;
    }

    // 先删除节点，然后插入到头部
    private void moveTohead(ListNode node){
        deleteNode(node);
        insertHead(node);
    }

    private void deleteNode(ListNode node){
        ListNode temp = node.pre;
        temp.next = node.next;
        node.next.pre = temp;
    }

    private void insertHead(ListNode node){
        ListNode temp = head.next;
        head.next = node;
        node.next = temp;
        temp.pre = node;
        node.pre = head;
    }
}



