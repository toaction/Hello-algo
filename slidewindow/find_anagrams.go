package slidewindow

/*
438. 找到字符串中所有字母异位词：https://leetcode.cn/problems/find-all-anagrams-in-a-string/

给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

示例 1:
	输入: s = "cbaebabacd", p = "abc"
	输出: [0,6]
	解释:
	起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
	起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
*/

func FindAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	// 所需字符的总个数
	need := len(p)
	// 相比于 p，窗口内缺失字符的具体个数
	count := [26]int{}
	for _, v := range p {
		count[v-'a']++
	}
	left, right := 0, 0
	for right < len(s) {
		// 所需字符
		if count[s[right]-'a'] > 0 {
			need--
		}
		count[s[right]-'a']--
		if need == 0 {
			// 尝试缩小窗口
			for left <= right && count[s[left]-'a'] < 0 {
				count[s[left]-'a']++
				left++
			}
			// 记录字母异位词其实索引
			if (right - left + 1) == len(p) {
				ans = append(ans, left)
			}
			// 移动窗口
			count[s[left]-'a']++
			left++
			need++
		}
		right++
	}
	return ans
}
