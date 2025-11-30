package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "basic example",
			arr:  []int{5, 2, 9, 1, 5, 6},
			want: []int{1, 2, 5, 5, 6, 9},
		},
		{
			name: "already sorted",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "single element",
			arr:  []int{42},
			want: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制数组避免修改原数组
			got := make([]int, len(tt.arr))
			copy(got, tt.arr)
			BubbleSort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "basic example",
			arr:  []int{8, 3, 1, 7, 0, 10, 2},
			want: []int{0, 1, 2, 3, 7, 8, 10},
		},
		{
			name: "reverse sorted",
			arr:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "negative numbers",
			arr:  []int{-3, 1, -4, 0, 2},
			want: []int{-4, -3, 0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.arr))
			copy(got, tt.arr)
			InsertionSort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "basic example",
			arr:  []int{29, 10, 14, 37, 13},
			want: []int{10, 13, 14, 29, 37},
		},
		{
			name: "with duplicates",
			arr:  []int{3, 1, 4, 1, 5, 9, 2, 6},
			want: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name: "all same elements",
			arr:  []int{7, 7, 7, 7},
			want: []int{7, 7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.arr))
			copy(got, tt.arr)
			SelectSort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "basic example",
			arr:  []int{3, 7, 8, 5, 2, 1, 9, 6, 4},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "already sorted",
			arr:  []int{1, 2, 3, 4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "with negative numbers",
			arr:  []int{0, -1, 2, -3, 4},
			want: []int{-3, -1, 0, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.arr))
			copy(got, tt.arr)
			if len(got) > 0 {
				QuickSort(got, 0, len(got)-1)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "basic example",
			arr:  []int{38, 27, 43, 3, 9, 82, 10},
			want: []int{3, 9, 10, 27, 38, 43, 82},
		},
		{
			name: "two elements",
			arr:  []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "mixed positive and negative",
			arr:  []int{5, -2, 9, -1, 0, 3},
			want: []int{-2, -1, 0, 3, 5, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.arr))
			copy(got, tt.arr)
			if len(got) > 0 {
				MergeSort(got, 0, len(got)-1)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "basic example",
			arr:  []int{4, 10, 3, 5, 1, 2},
			want: []int{1, 2, 3, 4, 5, 10},
		},
		{
			name: "already heap",
			arr:  []int{9, 5, 6, 2, 3},
			want: []int{2, 3, 5, 6, 9},
		},
		{
			name: "with duplicates",
			arr:  []int{4, 2, 2, 8, 3, 3, 1},
			want: []int{1, 2, 2, 3, 3, 4, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.arr))
			copy(got, tt.arr)
			HeapSort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
