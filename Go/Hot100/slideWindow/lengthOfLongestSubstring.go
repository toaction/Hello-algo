package slideWindow

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
// 示例 1:
//
// 输入: s = "abcabcbb"
//
// 输出: 3
//
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
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
		// 缩小窗口，使窗口不存在重复字符
		for right < len(bytes) && cur[bytes[right]] == 1 {
			cur[bytes[left]]--
			left++
		}
	}
	return ans
}
