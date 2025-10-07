package subString

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
// 示例 1：
//
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
// 滑动窗口
func MinWindow(s string, t string) string {
	as := []byte(s)
	at := []byte(t)
	// 还需字符的具体个数
	words := [128]int{}
	for _, v := range at {
		words[v]++
	}
	// 还需字符的总个数
	need := len(at)
	left, right := 0, 0
	ansLeft, ansRight := 0, len(as)
	for right < len(as) {
		if words[as[right]] > 0 {
			need--
		}
		words[as[right]]--
		if need == 0 {
			// 尽可能缩小窗口
			for words[as[left]] < 0 {
				words[as[left]]++
				left++
			}
			// 更新最小窗口
			if right-left < ansRight-ansLeft {
				ansLeft = left
				ansRight = right
			}
			// 移动窗口
			words[as[left]]++
			left++
			need++
		}
		right++
	}
	if ansRight == len(as) {
		return ""
	}
	return string(as[ansLeft : ansRight+1])
}
