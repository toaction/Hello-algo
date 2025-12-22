package skill

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic example",
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			name: "duplicate at the end",
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
		{
			name: "duplicate at the beginning",
			nums: []int{1, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicate(tt.nums); got != tt.want {
				t.Errorf("FindDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic example",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "majority at the end",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement(tt.nums); got != tt.want {
				t.Errorf("MajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic example",
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "needs reversal",
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "complex case",
			nums: []int{1, 5, 1},
			want: []int{5, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)
			NextPermutation(input)
			if !equalSlice(input, tt.want) {
				t.Errorf("NextPermutation(%v) = %v, want %v", tt.nums, input, tt.want)
			}
		})
	}
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic example",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "larger array",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.nums); got != tt.want {
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic example",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "already sorted",
			nums: []int{0, 0, 1, 1, 2, 2},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "reverse order",
			nums: []int{2, 2, 1, 1, 0, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)
			SortColors(input)
			if !equalSlice(input, tt.want) {
				t.Errorf("SortColors(%v) = %v, want %v", tt.nums, input, tt.want)
			}
		})
	}
}

// equalSlice compares two slices for equality
func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
