package graph

/*
200. 岛屿数量：https://leetcode.cn/problems/number-of-islands/

给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
	输入：grid = [
	['1','1','1','1','0'],
	['1','1','0','1','0'],
	['1','1','0','0','0'],
	['0','0','0','0','0']
	]
	输出：1
*/

// 深度优先搜索
// 若网格值为1，则将其置为2，代表已扫描
func NumIslands(grid [][]byte) int {
	count := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i int, j int) {
	// 越界
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	// 已遍历过或不属于陆地
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
