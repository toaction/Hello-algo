package matrix

import (
	"reflect"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			5,
			true,
		},
		{
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			20,
			false,
		},
		{
			[][]int{{1, 4}, {2, 5}},
			5,
			true,
		},
	}

	for _, tt := range tests {
		result := SearchMatrix(tt.matrix, tt.target)
		if result != tt.want {
			t.Errorf("SearchMatrix(%v, %v) = %v, expected %v", tt.matrix, tt.target, result, tt.want)
		}
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		{
			[][]int{{1}},
			[][]int{{1}},
		},
	}

	for _, tt := range tests {
		matrix := make([][]int, len(tt.matrix))
		for i := range tt.matrix {
			matrix[i] = make([]int, len(tt.matrix[i]))
			copy(matrix[i], tt.matrix[i])
		}
		Rotate(matrix)
		if !reflect.DeepEqual(matrix, tt.want) {
			t.Errorf("Rotate(%v) = %v, expected %v", tt.matrix, matrix, tt.want)
		}
	}
}

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			[][]int{{1, 2}, {3, 4}},
			[]int{1, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		result := SpiralOrder(tt.matrix)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("SpiralOrder(%v) = %v, expected %v", tt.matrix, result, tt.want)
		}
	}
}

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			[][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}},
			[][]int{{1, 0, 3}, {0, 0, 0}, {7, 0, 9}},
		},
	}

	for _, tt := range tests {
		matrix := make([][]int, len(tt.matrix))
		for i := range tt.matrix {
			matrix[i] = make([]int, len(tt.matrix[i]))
			copy(matrix[i], tt.matrix[i])
		}
		SetZeroes(matrix)
		if !reflect.DeepEqual(matrix, tt.want) {
			t.Errorf("SetZeroes(%v) = %v, expected %v", tt.matrix, matrix, tt.want)
		}
	}
}
