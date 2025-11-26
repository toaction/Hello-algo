package slidewindow

// 3. 无重复字符的最长子串
//
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
//	输入: "abcabcbb"
//	输出: 3
//	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// 解题思路：
// 1. 滑动窗口：维护一个无重复字符的窗口
// 2. 双指针：left指向窗口开始，right指向窗口结束
// 3. 字符计数：用cur数组记录窗口中每个字符的出现次数
// 4. 窗口操作：
//   - 扩展：右指针向右移动，直到遇到重复字符
//   - 收缩：当遇到重复时，左指针向右移动直到移除重复字符
//   - 记录：每次扩展后更新最大窗口长度
//
// 5. 核心逻辑：
//   - cur[bytes[right]] == 0：字符未重复，可以加入窗口
//   - cur[bytes[right]] == 1：字符重复，需要收缩窗口
//   - ans = max(ans, right-left)：记录当前最大长度
//
// 时间复杂度：O(n) - 每个字符最多被访问两次
// 空间复杂度：O(1) - 固定大小的字符数组
func LengthOfLongestSubstring(s string) int {
	bytes := []byte(s)
	ans := 0
	left, right := 0, 0
	cur := [128]int{}
	for right < len(bytes) {
		// 移动右指针，直到遇到重复字符
		for right < len(bytes) && cur[bytes[right]] == 0 {
			cur[bytes[right]]++
			right++
		}
		// 记录窗口大小
		ans = max(ans, right-left)
		// 缩小窗口，移除重复的字符
		for right < len(bytes) && cur[bytes[right]] == 1 {
			cur[bytes[left]]--
			left++
		}
	}
	return ans
}
