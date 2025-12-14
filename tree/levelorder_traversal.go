package tree

/*
102. 二叉树的层序遍历：https://leetcode.cn/problems/binary-tree-level-order-traversal/

给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
*/

func LevelOrder(root *TreeNode) [][]int {
	ans := &[][]int{}
	levelOrderDFS(root, 1, ans)
	return *ans
}

func levelOrderDFS(node *TreeNode, level int, vals *[][]int) {
	if node == nil {
		return
	}
	if len(*vals) < level {
		*vals = append(*vals, []int{})
	}
	(*vals)[level-1] = append((*vals)[level-1], node.Val)
	levelOrderDFS(node.Left, level+1, vals)
	levelOrderDFS(node.Right, level+1, vals)
}

// 迭代，使用队列存放每一层的节点
func LevelOrderBaseInteration(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		// 本层节点数量
		count := len(queue)
		vals := []int{}
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, vals)
	}
	return ans
}
