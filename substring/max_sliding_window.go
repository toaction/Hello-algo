package substring

// 239. 滑动窗口最大值
//
// https://leetcode-cn.com/problems/sliding-window-maximum/
//
// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
// 返回滑动窗口中的最大值。
//
// 示例:
//
//	输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//	输出：[3,3,5,5,6,7]
//	解释：
//	滑动窗口的位置                最大值
//	---------------               -----
//	[1  3  -1] -3  5  3  6  7      3
//	1 [3  -1  -3] 5  3  6  7       3
//	1  3 [-1  -3  5] 3  6  7       5
//	1  3  -1 [-3  5  3] 6  7       5
//	1  3  -1  -3 [5  3  6] 7       6
//	1  3  -1  -3  5 [3  6  7]      7
//
// 解题思路：单调栈
func MaxSlidingWindow(nums []int, k int) []int {
	// 单调栈，存放窗口内大小递减的元素下标
	stack := make([]int, 0)
	ans := make([]int, 0)
	// 初始化栈和窗口
	for i := 0; i < k; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans = append(ans, nums[stack[0]])
	left, right := 1, k
	for right < len(nums) {
		// 淘汰能力不足的
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[right] {
			stack = stack[:len(stack)-1]
		}
		// 入职新员工
		stack = append(stack, right)
		// 优化年龄大的
		if stack[0] < left {
			stack = stack[1:]
		}
		// 记录窗口最大值
		ans = append(ans, nums[stack[0]])
		left++
		right++
	}
	return ans
}
