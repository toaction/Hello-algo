package subString

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
// 返回 滑动窗口中的最大值 。
//
// 示例 1：
//
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
//  [1  3  -1] -3  5  3  6  7      3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
//
// 单调栈
func MaxSlidingWindow(nums []int, k int) []int {
	stack := make([]int, 0)
	ans := make([]int, 0)
	for i, num := range nums {
		// 不要能力差的
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		// 移除年纪大的
		if i >= k && stack[0] <= i-k {
			stack = stack[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[stack[0]])
		}
	}
	return ans
}
