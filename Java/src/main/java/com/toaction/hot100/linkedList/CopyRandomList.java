package com.toaction.hot100.linkedList;

import java.util.HashMap;
import java.util.Map;

public class CopyRandomList {

    /*
    https://leetcode.cn/problems/copy-list-with-random-pointer/
    给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
    请返回这个链表的深拷贝。
    * */

    public static RandomNode copyRandomList(RandomNode head) {
        // 原节点、新节点
        Map<RandomNode, RandomNode> map = new HashMap<>();
        RandomNode cur = head;
        while(cur != null) {
            map.put(cur, new RandomNode(cur.val));
            cur = cur.next;
        }
        cur = head;
        while(cur != null) {
            RandomNode node = map.get(cur);
            node.next = map.getOrDefault(cur.next, null);
            node.random = map.getOrDefault(cur.random, null);
            cur = cur.next;
        }
        return map.get(head);
    }
}
