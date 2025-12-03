package graph

/*
207. 课程表：https://leetcode.cn/problems/course-schedule/

你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
	输入：numCourses = 2, prerequisites = [[1,0]]
	输出：true
	解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
*/

// 构造课程之间的关系：先修课程 -> 学习课程
// 使用一个邻接表记录课程之间的关系
// 维护一个队列存储入度为 0 的课程，将该课程作为先修课程减少要学习课程的入度
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// 每个课程的入度
	inDegree := make([]int, numCourses)
	// 先修课程 -> 学习课程
	preToStudy := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		study, pre := prerequisites[i][0], prerequisites[i][1]
		preToStudy[pre] = append(preToStudy[pre], study)
		inDegree[study]++
	}
	// 能学习的课程数量
	count := 0
	canStudy := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			count++
			canStudy = append(canStudy, i)
		}
	}
	// 学习课程，将该课程作为先修课程，减少相对应课程的入度
	for len(canStudy) > 0 {
		study := canStudy[0]
		canStudy = canStudy[1:]
		for _, v := range preToStudy[study] {
			inDegree[v]--
			if inDegree[v] == 0 {
				count++
				canStudy = append(canStudy, v)
			}
		}
	}
	return count == numCourses
}
