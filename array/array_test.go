package array

import (
	"testing"
)

func TestMaxSubArrayBasePreSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "standard leetcode example",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "all positive numbers",
			nums: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "all negative numbers",
			nums: []int{-3, -2, -1, -4, -5},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSubArrayBasePreSum(tt.nums)
			if got != tt.want {
				t.Errorf("MaxSubArrayBasePreSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSubArrayBaseDP(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "standard leetcode example",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "single element positive",
			nums: []int{5},
			want: 5,
		},
		{
			name: "mixed with zero",
			nums: []int{-1, 0, -2, 3, -1},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSubArrayBaseDP(tt.nums)
			if got != tt.want {
				t.Errorf("MaxSubArrayBaseDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "standard example",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "k larger than length",
			nums: []int{1, 2, 3, 4},
			k:    6,
			want: []int{3, 4, 1, 2},
		},
		{
			name: "k is zero",
			nums: []int{1, 2, 3, 4},
			k:    0,
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)
			Rotate(input, tt.k)
			if !equalSlices(input, tt.want) {
				t.Errorf("Rotate() = %v, want %v", input, tt.want)
			}
		})
	}
}

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "standard example",
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			name: "contains 1 and 3",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			name: "all negative",
			nums: []int{-1, -2, -3},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)
			got := FirstMissingPositive(input)
			if got != tt.want {
				t.Errorf("FirstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "standard leetcode example",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "contains zero",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "single element",
			nums: []int{5},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProductExceptSelf(tt.nums)
			if !equalSlices(got, tt.want) {
				t.Errorf("ProductExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name: "standard leetcode example",
			intervals: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "all overlapping intervals",
			intervals: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
			want: [][]int{
				{1, 6},
			},
		},
		{
			name: "no overlapping intervals",
			intervals: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			want: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.intervals)
			if !equalIntervals(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
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

func equalIntervals(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalSlices(a[i], b[i]) {
			return false
		}
	}
	return true
}
