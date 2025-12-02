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
