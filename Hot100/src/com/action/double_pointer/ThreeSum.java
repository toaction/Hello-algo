package com.action.double_pointer;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class ThreeSum {

    /*
    source: https://leetcode.cn/problems/3sum
    desc:
        给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
        满足 i != j、i != k 且 j != k ，
        同时还满足 nums[i] + nums[j] + nums[k] == 0 。
        请你返回所有和为 0 且不重复的三元组。
    * */


    /*
    首先将数组排序
    外层循环指向第一个数，使用双指针指向第二个数和第三个数
    若三数之和小于0，则移动左指针，使结果增大；若大于0，移动右指针，使结果减小
    */
    public static List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> ans = new ArrayList<>();
        Arrays.sort(nums);
        // 固定最左边元素
        for (int i=0; i<nums.length-2; i++) {
            // 去除重复
            if (i>0 && nums[i]==nums[i-1]) continue;
            int left=i+1, right=nums.length-1;
            while(left < right){
                int sum = nums[i] + nums[left] + nums[right];
                if (sum < 0) {
                    left++;
                }else if (sum > 0) {
                    right--;
                }else {
                    ans.add(Arrays.asList(nums[i], nums[left], nums[right]));
                    // 移动，避免重复
                    int value = nums[left];
                    while(left<right && nums[left]==value) left++;
                }
            }
        }
        return ans;
    }


    public static void main(String[] args) {
        int[] nums = {-1, 0, 1, 2, -1, -4};
        List<List<Integer>> ans = threeSum(nums);
        for (List<Integer> list : ans) {
            for (Integer integer : list) {
                System.out.print(integer + " ");
            }
            System.out.println();
        }
    }
}
