package doublepointer

// 42. 接雨水
//
// https://leetcode-cn.com/problems/trapping-rain-water/
//
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 示例 1：
//
//	输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//	输出：6
//
// Trap 计算柱状图中能接的雨水总量
//
// 解题思路：
// 1. 双指针：left从左往右，right从右往左
// 2. 维护两个变量：lm(左边最大高度), rm(右边最大高度)
// 3. 核心原理：
//   - 在任意位置，能接的雨水量取决于该位置左右两边最高柱子的较小值
//   - 如果左边最大高度 < 右边最大高度，当前位置的雨水量由左边决定
//   - 反之，由右边决定
//
// 4. 移动策略：
//   - lm < rm时：左边位置雨水量确定，移动left指针
//   - lm >= rm时：右边位置雨水量确定，移动right指针
//
// 5. 计算方式：当前位置雨水量 = min(lm, rm) - 当前高度
//
// 时间复杂度：O(n) - 只需遍历数组一次
// 空间复杂度：O(1) - 只使用常数额外空间
func Trap(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	lm, rm := 0, 0
	for left < right {
		lm = max(lm, height[left])
		rm = max(rm, height[right])
		if lm < rm {
			// 左边高度小，记录左边能接的雨水并移动指针
			ans += lm - height[left]
			left++
		} else {
			// 右边
			ans += rm - height[right]
			right--
		}
		// 移动短的一边，最后长的一边一定是最高的，无需 left = right
	}
	return ans
}
