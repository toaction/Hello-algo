package matrix

/*
48. 旋转图像: https://leetcode.cn/problems/rotate-image/

给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

示例 1：
	输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
	输出：[[7,4,1],[8,5,2],[9,6,3]]
*/

// 先按照对角线交换元素，然后左右交换
func Rotate(matrix [][]int) {
	n := len(matrix)
	// 对角线交换
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 左右交换
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}
