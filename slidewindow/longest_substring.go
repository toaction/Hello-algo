package slidewindow

/* 
3. 无重复字符的最长子串：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
	输入: "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。 
*/


func LengthOfLongestSubstring(s string) int {
    ans := 0
    // 统计窗口内各个字符的个数
    count := [128]int{}
    left, right := 0, 0
    for right < len(s) {
        // 移动右指针，直到遇到重复字符
        for right < len(s) && count[s[right]] == 0 {
            count[s[right]]++
            right++
        }
        ans = max(ans, right-left)
        // 缩小窗口，移除重复的字符
        for right < len(s) && count[s[right]] > 0 {
            count[s[left]]--
            left++
        }
    }
    return ans
}
