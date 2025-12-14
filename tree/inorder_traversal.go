package tree

/*
94. 二叉树的中序遍历：https://leetcode.cn/problems/binary-tree-inorder-traversal/

给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
*/

func InorderTraversal(root *TreeNode) []int {
	ans := &[]int{}
	inorderDFS(root, ans)
	return *ans
}

func inorderDFS(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	inorderDFS(node.Left, vals)
	*vals = append(*vals, node.Val)
	inorderDFS(node.Right, vals)
}

// 显式维护一个栈
// 遍历左子树，将节点入栈，直到遍历节点为 null
// 弹出栈顶节点，记录节点值
// 将节点变更为该节点的右子树根节点，重复以上过程直到栈为空
func InorderTraversalBaseIteration(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		ans = append(ans, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return ans
}
