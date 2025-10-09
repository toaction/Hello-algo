package matrix

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
// 示例 1：
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
//
// 模拟旋转
func SpiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	m := len(matrix)
	n := len(matrix[0])
	top, bottom, left, right := 0, m-1, 0, n-1
	for top <= bottom && left <= right {
		// 从左到右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++
		// 从上到下
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		// 从右到左，需判断条件是否符合
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			bottom--
		}
		// 从下到上，判断条件是否符合
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}
	return ans
}
