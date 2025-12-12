package doublepointer

import "sort"

/* 
15. 三数之和：https://leetcode.cn/problems/3sum/

给你一个整数数组 nums ，
判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。
请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：
	输入：nums = [-1,0,1,2,-1,-4]
	输出：[[-1,-1,2],[-1,0,1]]
	解释：
	nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
	nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
	nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
	不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。 
*/


// 先将数组进行排序，然后使用三个指针 i, left, right 分别指向三个数
// 保持 i 不变，若总和小于 0，移动 left，大于 0，移动 right
func ThreeSum(nums []int) [][]int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    ans := make([][]int, 0)
    for i:=0; i<len(nums); i++ {
        // 外层去重
        if i>0 && nums[i]==nums[i-1] {
            continue
        }
        left, right := i+1, len(nums)-1
        for (left < right) {
            sum := nums[i] + nums[left] + nums[right]
            if sum < 0 {    // 移动左指针，总和增大
                left++
            }else if sum > 0 {
                right--     // 移动右指针，总和减小
            }else {
                ans = append(ans, []int{nums[i], nums[left], nums[right]})
                base := nums[left]
                for left < right && nums[left]==base {
                    left++  // 内层去重
                }
            }
        }
    }
    return ans
}
