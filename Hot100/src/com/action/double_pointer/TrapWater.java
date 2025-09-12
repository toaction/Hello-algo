package com.action.double_pointer;

public class TrapWater {

    /*
    source: https://leetcode.cn/problems/trapping-rain-water/
    desc:
        给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
    * */


    /*
    每个柱子能接的雨水等于 Math.min(左边最高的柱子，右边最高的柱子)-柱子高度
    1. 两次遍历
        第一次遍历，找到最高的柱子
        从左遍历到最高的柱子，每个柱子能接到的雨水为其(左边最高的柱子-柱子高度)
        从右遍历到最高的柱子，每个柱子能接到的雨水为其(右边最高的柱子-柱子高度)
    2. 一次遍历
        双指针
        使用两个变量记录左边最高的柱子和右边最高的柱子
        若柱子小于两者之中的最大值，则一定小于全部柱子的最大值
        移动柱子低的一端的指针，记录能接的雨水
    */

    public static int trap(int[] height) {
        int left=0, leftHeight=0, right=height.length-1, rightHeight=0;
        int ans = 0;
        while(left<right){
            leftHeight = Math.max(height[left], leftHeight);
            rightHeight = Math.max(height[right], rightHeight);
            if(leftHeight < rightHeight){
                ans += leftHeight - height[left];
                left++;
            }else{
                ans += rightHeight - height[right];
                right--;
            }
        }
        return ans;
    }

    public static void main(String[] args) {
        int[] height = {0,1,0,2,1,0,1,3,2,1,2,1};
        System.out.println(trap(height));
    }
}
