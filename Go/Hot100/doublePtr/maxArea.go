package doublePtr

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 思路：
//
// 移动指针，宽度减少，高度就要增大
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
