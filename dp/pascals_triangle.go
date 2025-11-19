package dp

// 118. 杨辉三角
//
// https://leetcode.cn/problems/pascals-triangle/
//
// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方数之和。
//
// 示例 1:
//
//	输入: numRows = 5
//	输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
func Generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ans[i][j] = 1
			} else {
				ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
			}
		}
	}
	return ans
}
