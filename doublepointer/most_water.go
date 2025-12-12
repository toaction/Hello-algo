package doublepointer

/* 
11. 盛最多水的容器：https://leetcode-cn.com/problems/container-with-most-water/

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

示例：
	输入：[1,8,6,2,5,4,8,3,7]
	输出：49 
*/


// 双指针，移动高度小的边
func MaxArea(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	for left < right {
		h := min(height[left], height[right])
		ans = max(ans, h*(right-left))
		// 移动短的一边
		if height[left] < height[right] {
			for left < right && height[left] <= h {
				left++
			}
		} else {
			for left < right && height[right] <= h {
				right--
			}
		}
	}
	return ans
}
