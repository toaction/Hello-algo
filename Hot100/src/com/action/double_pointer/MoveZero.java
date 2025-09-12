package com.action.double_pointer;

import java.util.Arrays;

public class MoveZero {

    /*
    source: https://leetcode.cn/problems/move-zeroes/
    desc:
        给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
    require:
        空间复杂度： O(1)
    * */

    public static void moveZeroes(int[] nums) {
        int left = 0;
        for(int i=0; i<nums.length; i++){
            if(nums[i] == 0) continue;
            // 交换元素
            int temp = nums[left];
            nums[left] = nums[i];
            nums[i] = temp;
            left++;
        }
    }

    public static void main(String[] args) {
        int[] nums = new int[]{0,1,0,3,12};
        moveZeroes(nums);
        System.out.println(Arrays.toString(nums));
    }
}
