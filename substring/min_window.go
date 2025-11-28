package substring

// 76. 最小覆盖子串
//
// https://leetcode-cn.com/problems/minimum-window-substring/
//
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 示例 1：
//
//	输入：s = "ADOBECODEBANC", t = "ABC"
//	输出："BANC"
//
// 解题思路：滑动窗口
func MinWindow(s string, t string) string {
	// 记录所需 t 中字符个数
	need := make([]int, 128)
	// 总个数
	diff := len(t)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	minLeft, minRight := 0, len(s)
	left, right := 0, 0
	for right < len(s) {
		// t 中字符
		if need[s[right]] > 0 {
			diff--
		}
		need[s[right]]--
		if diff == 0 {
			// 尝试缩小窗口
			for need[s[left]] < 0 {
				need[s[left]]++
				left++
			}
			// 记录窗口大小
			if right-left < minRight-minLeft {
				minLeft, minRight = left, right
			}
			// 移动窗口
			need[s[left]]++
			left++
			diff++
		}
		right++
	}
	if minRight == len(s) {
		return ""
	}
	return s[minLeft : minRight+1]
}
