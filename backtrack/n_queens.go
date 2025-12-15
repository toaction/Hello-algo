package backtrack

/*
51. N 皇后：https://leetcode.cn/problems/n-queens/

按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例 1：
	输入：n = 4
	输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
*/

// 每行放一个皇后/每列放一个皇后，本质上是枚举所处列号的全排列，找到符合条件的排列
// 不符合条件：皇后之间可以攻击，左对角线：行号-列号相同；右对角线：行号+列号相同
func SolveNQueens(n int) [][]string {
	left, right := make([]bool, 2*n), make([]bool, 2*n)
	visited := make([]bool, n)
	path := []int{}
	ans := [][]string{}
	backTrackQueen(n, visited, left, right, &path, &ans)
	return ans
}

func backTrackQueen(n int, visited []bool, left []bool, right []bool, path *[]int, ans *[][]string) {
	if len(*path) == n {
		locations := make([]string, 0)
		for _, v := range *path {
			location := make([]byte, n)
			for i := 0; i < n; i++ {
				location[i] = '.'
			}
			location[v] = 'Q'
			locations = append(locations, string(location))
		}
		*ans = append(*ans, locations)
		return
	}
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		// 放置在第 k 行，第 i 列
		k := len(*path)
		if left[k-i+n] || right[k+i] {
			continue
		}
		left[k-i+n] = true
		right[k+i] = true
		visited[i] = true
		*path = append(*path, i)
		backTrackQueen(n, visited, left, right, path, ans)
		*path = (*path)[:len(*path)-1]
		visited[i] = false
		right[k+i] = false
		left[k-i+n] = false
	}
}
