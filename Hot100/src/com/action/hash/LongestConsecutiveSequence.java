package com.action.hash;

import java.util.HashMap;
import java.util.Map;

public class LongestConsecutiveSequence {

    /*
    source: https://leetcode.cn/problems/longest-consecutive-sequence/
    desc:
        给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
    require:
        时间复杂度：O(n)
    * */


    public static int longestConsecutive(int[] nums) {
        // 将所有整数存放到 map 中，方便判断该整数是否存在前位数
        Map<Integer, Integer> map = new HashMap<>();
        for(int i:nums){
            map.put(i, 0);
        }
        int ans = 0;
        // 遍历 map，若不存在前位数，代表该元素为起点，计算连续序列的长度
        for(Integer key:map.keySet()){
            if(map.containsKey(key-1)) continue;
            int k = key;
            int count = 0;
            while(map.containsKey(k)){
                k++;
                count++;
            }
            ans = Math.max(ans, count);
        }
        return ans;
    }

    public static void main(String[] args) {
        int[] nums = {100,4,200,1,3,2};
        System.out.println(longestConsecutive(nums));
    }
}
