package normalArray

import (
	"reflect"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		result := FirstMissingPositive(nums)
		if result != tt.want {
			t.Errorf("FirstMissingPositive(%v) = %v, expected %v", tt.nums, result, tt.want)
		}
	}
}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for _, tt := range tests {
		result := MaxSubArray(tt.nums)
		if result != tt.want {
			t.Errorf("MaxSubArray(%v) = %v, expected %v", tt.nums, result, tt.want)
		}
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			[][]int{{1, 4}, {0, 4}},
			[][]int{{0, 4}},
		},
	}

	for _, tt := range tests {
		intervals := make([][]int, len(tt.intervals))
		for i, interval := range tt.intervals {
			intervals[i] = make([]int, len(interval))
			copy(intervals[i], interval)
		}
		result := Merge(intervals)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("Merge(%v) = %v, expected %v", tt.intervals, result, tt.want)
		}
	}
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{2, 3, 4, 5}, []int{60, 40, 30, 24}},
	}

	for _, tt := range tests {
		result := ProductExceptSelf(tt.nums)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("ProductExceptSelf(%v) = %v, expected %v", tt.nums, result, tt.want)
		}
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2}, 3, []int{2, 1}},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		Rotate(nums, tt.k)
		if !reflect.DeepEqual(nums, tt.want) {
			t.Errorf("Rotate(%v, %v) = %v, expected %v", tt.nums, tt.k, nums, tt.want)
		}
	}
}
