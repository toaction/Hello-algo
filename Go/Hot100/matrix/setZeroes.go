package matrix

// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//
// 示例 1：
//
// 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：[[1,0,1],[0,0,0],[1,0,1]]
func SetZeroes(matrix [][]int) {
	// 使用额外的两个数记录 第一行 和 第一列 是否存在 0
	r, c := false, false
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			c = true
			break
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			r = true
			break
		}
	}
	// 使用第一行和第一列记录 第j列 和 第i行 是否存在 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				// 第 j 列为 0
				matrix[0][j] = 0
				// 第 i 行为 0
				matrix[i][0] = 0
			}
		}
	}
	// 置 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if r {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if c {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
