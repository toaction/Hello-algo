package binarysearch

import (
	"reflect"
	"testing"
)

// TestSearchInsert tests the SearchInsert function
func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "target exists in array",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "target not exists, should insert at position 2",
			nums:   []int{1, 3, 5, 6},
			target: 4,
			want:   2,
		},
		{
			name:   "target greater than all elements",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchInsert(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSearchMatrix tests the SearchMatrix function
func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "target exists in matrix",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			name:   "target not exists in matrix",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
		{
			name:   "single element matrix, target exists",
			matrix: [][]int{{1}},
			target: 1,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchMatrix(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("SearchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSearchRange tests the SearchRange function
func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "target exists multiple times",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "target not exists",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "target exists once",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   []int{2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestFindMedianSortedArrays tests the FindMedianSortedArrays function
func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "odd total length",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "even total length",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			name:  "one array empty",
			nums1: []int{},
			nums2: []int{1},
			want:  1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMedianSortedArrays(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("FindMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSearchRotate tests the SearchRotate function
func TestSearchRotate(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "target exists in rotated array",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "target not exists in rotated array",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "no rotation, target exists",
			nums:   []int{1, 2, 3, 4, 5, 6},
			target: 4,
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRotate(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("SearchRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestFindMinInRotate tests the FindMinInRotate function
func TestFindMinInRotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "rotated array with min in middle",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "rotated array with min at end",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "no rotation, min at beginning",
			nums: []int{1, 2, 3, 4, 5},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMinInRotate(tt.nums)
			if got != tt.want {
				t.Errorf("FindMinInRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
