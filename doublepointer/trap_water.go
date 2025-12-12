package doublepointer

/* 
42. 接雨水：https://leetcode-cn.com/problems/trapping-rain-water/

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
	输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
	输出：6 
*/


// 每个柱子能接的雨水量 = min(左边最高的柱子，右边最高的柱子) - 柱子高度
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
