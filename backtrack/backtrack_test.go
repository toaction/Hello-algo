package backtrack

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "standard leetcode example",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "single element",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "two elements",
			nums: []int{1, 2},
			want: [][]int{{1, 2}, {2, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permute(tt.nums)
			if !equalMatrix(got, tt.want) {
				t.Errorf("Permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalMatrix(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]bool)
	for _, row := range a {
		key := fmt.Sprintf("%v", row)
		m[key] = true
	}
	for _, row := range b {
		key := fmt.Sprintf("%v", row)
		if !m[key] {
			return false
		}
	}
	return true
}

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "standard leetcode example",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		},
		{
			name: "empty array",
			nums: []int{},
			want: [][]int{{}},
		},
		{
			name: "single element",
			nums: []int{1},
			want: [][]int{{}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subsets(tt.nums)
			if !equalMatrix(got, tt.want) {
				t.Errorf("Subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "standard leetcode example",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "empty string",
			digits: "",
			want:   []string{},
		},
		{
			name:   "single digit",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LetterCombinations(tt.digits)
			if !equalStrings(got, tt.want) {
				t.Errorf("LetterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]bool)
	for _, s := range a {
		m[s] = true
	}
	for _, s := range b {
		if !m[s] {
			return false
		}
	}
	return true
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "standard leetcode example",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "with duplicates in candidates",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "single candidate multiple times",
			candidates: []int{2},
			target:     6,
			want:       [][]int{{2, 2, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombinationSum(tt.candidates, tt.target)
			if !equalMatrix(got, tt.want) {
				t.Errorf("CombinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "standard leetcode example",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "n = 1",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "n = 2",
			n:    2,
			want: []string{"(())", "()()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateParenthesis(tt.n)
			if !equalStrings(got, tt.want) {
				t.Errorf("GenerateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name: "standard leetcode example - found",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
		{
			name: "standard leetcode example - not found",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEEK",
			want: false,
		},
		{
			name: "single cell board",
			board: [][]byte{
				{'A'},
			},
			word: "A",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordExist(tt.board, tt.word)
			if got != tt.want {
				t.Errorf("WordExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{
			name: "standard leetcode example",
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "single character",
			s:    "a",
			want: [][]string{{"a"}},
		},
		{
			name: "all same characters",
			s:    "aaa",
			want: [][]string{{"a", "a", "a"}, {"a", "aa"}, {"aa", "a"}, {"aaa"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Partition(tt.s)
			if !equalStringMatrix(got, tt.want) {
				t.Errorf("Partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalStringMatrix(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]bool)
	for _, row := range a {
		key := fmt.Sprintf("%v", row)
		m[key] = true
	}
	for _, row := range b {
		key := fmt.Sprintf("%v", row)
		if !m[key] {
			return false
		}
	}
	return true
}

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		{
			name: "standard leetcode example - n=4",
			n:    4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name: "n=1",
			n:    1,
			want: [][]string{{"Q"}},
		},
		{
			name: "n=3 - no solution",
			n:    3,
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveNQueens(tt.n)
			if !equalStringMatrix(got, tt.want) {
				t.Errorf("SolveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
