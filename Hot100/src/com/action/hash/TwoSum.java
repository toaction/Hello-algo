package com.action.hash;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class TwoSum {

    /*
    source: https://leetcode.cn/problems/two-sum
    desc:
        给定一个整数数组 nums 和一个整数目标值 target，
        请你在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标。
    solution:
        hashmap
    * */

    public static int[] twoSum(int[] nums, int target) {
        // num, index
        Map<Integer, Integer> map = new HashMap<>();
        for(int i=0; i<nums.length; i++){
            if(map.containsKey(target - nums[i])){
                int index = map.get(target - nums[i]);
                return new int[]{index, i};
            }
            map.put(nums[i], i);
        }
        return new int[]{0};
    }

    public static void main(String[] args) {
        int[] nums = new int[]{2,7,11,15};
        int[] result = twoSum(nums, 9);
        System.out.println(Arrays.toString(result));
    }

}
