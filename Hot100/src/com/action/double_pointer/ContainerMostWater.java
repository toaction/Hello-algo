package com.action.double_pointer;

public class ContainerMostWater {

    /*
    source: https://leetcode.cn/problems/container-with-most-water/
    desc:
        给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
        找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
        返回容器可以储存的最大水量。
    * */



    /*
    初始左指针left指向0，右指针right指向n-1
    由于移动指针会导致长度变小，因此想要容纳更多的水，必须移动短板
    移动短板才有可能使面积变大, 尝试找到比短板更长的板
    */
    public static int maxArea(int[] height) {
        int ans = 0;
        int n = height.length;
        int left=0, right=n-1;
        while (left<right){
            ans = Math.max(ans, Math.min(height[left], height[right])*(right-left));
            if(height[left] < height[right]){
                int pre = height[left];
                // 移动左指针，找到比原来高的
                while (left<right && height[left]<=pre) left++;
            }else {
                int pre = height[right];
                // 移动右指针，找到比原来高的
                while (left<right && height[right]<=pre) right--;
            }
        }
        return ans;
    }


    public static void main(String[] args) {
        int[] height = {1,8,6,2,5,4,8,3,7};
        System.out.println(maxArea(height));
    }
}
