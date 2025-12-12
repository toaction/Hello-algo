package matrix

/*
73. 矩阵置零：https://leetcode.cn/problems/set-matrix-zeroes/

给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

示例 1:
	输入: matrix = [[1,1,1],[1,0,1],[1,1,1]]
	输出: [[1,0,1],[0,0,0],[1,0,1]]
*/

// 使用第一行保存 存在元素为0的列
// 使用第一列保存 存在元素为0的行
func SetZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	// 第一行、第一列是否存在 0
	row, col := false, false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col = true
			break
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			row = true
			break
		}
	}
	// 记录存在 0 的行和列
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	// 置零
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if col {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
