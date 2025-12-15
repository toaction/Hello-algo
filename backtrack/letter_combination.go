package backtrack

/*
17. 电话号码的字母组合：https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

示例 1:
	输入：digits = "23"
	输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/

func LetterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	digitToWord := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	ans := []string{}
	path := []byte{}
	backTrackLetter(digits, digitToWord, 0, &ans, &path)
	return ans
}

func backTrackLetter(digits string, digitToWord map[byte]string, idx int, ans *[]string, path *[]byte) {
	if idx >= len(digits) {
		*ans = append(*ans, string(*path))
		return
	}
	word := digitToWord[digits[idx]]
	for i := 0; i < len(word); i++ {
		*path = append(*path, word[i])
		backTrackLetter(digits, digitToWord, idx+1, ans, path)
		*path = (*path)[:len(*path)-1]
	}
}
