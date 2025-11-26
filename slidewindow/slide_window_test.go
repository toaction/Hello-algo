package slidewindow

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{
		{
			name: "standard example from LeetCode",
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},
		{
			name: "multiple overlapping anagrams",
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
		{
			name: "no anagrams found",
			s:    "abcdefg",
			p:    "xyz",
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAnagrams(tt.s, tt.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "standard example from LeetCode",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "all unique characters",
			s:    "abcdefg",
			want: 7,
		},
		{
			name: "all same characters",
			s:    "aaaaa",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
