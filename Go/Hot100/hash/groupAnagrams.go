package hash

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 示例 1:
//
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 解释：
//
// 在 strs 中没有字符串可以通过重新排列来形成 "bat"。
// 字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
// 字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。
func GroupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for _, c := range str {
			key[c-'a']++
		}
		hashMap[key] = append(hashMap[key], str)
	}
	result := make([][]string, 0, len(hashMap))
	for _, v := range hashMap {
		result = append(result, v)
	}
	return result
}
