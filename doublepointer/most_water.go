package doublepointer

// 11. 盛最多水的容器
//
// https://leetcode-cn.com/problems/container-with-most-water/
//
// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 示例：
//
//	输入：[1,8,6,2,5,4,8,3,7]
//	输出：49
//
// 解题思路：
// 1. 双指针初始位置：left=0, right=n-1，分别指向数组两端
// 2. 容器面积 = min(左边高度, 右边高度) * (右指针-左指针)
// 3. 关键思想：每次移动较短的那一边
//   - 移动较长边：宽度减少，高度不变或减小，面积一定减小
//   - 移动较短边：宽度减少，但高度可能增加，面积可能增大
//
// 4. 优化：跳过高度≤当前高度的指针位置，避免无效计算
//
// 时间复杂度：O(n) - 每个元素最多访问一次
// 空间复杂度：O(1) - 只使用常数额外空间
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
