package slideWindow

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"tmmzuxt", 5},
	}

	for _, tt := range tests {
		result := LengthOfLongestSubstring(tt.s)
		if result != tt.want {
			t.Errorf("LengthOfLongestSubstring(%q) = %v, expected %v", tt.s, result, tt.want)
		}
	}
}

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
		{"baa", "aa", []int{1}},
	}

	for _, tt := range tests {
		result := FindAnagrams(tt.s, tt.p)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("FindAnagrams(%q, %q) = %v, expected %v", tt.s, tt.p, result, tt.want)
		}
	}
}
