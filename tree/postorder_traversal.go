package tree

/*
145. 二叉树的后序遍历：https://leetcode.cn/problems/binary-tree-postorder-traversal/

给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
*/

func PostorderTraversal(root *TreeNode) []int {
	ans := &[]int{}
	postorderDFS(root, ans)
	return *ans
}

func postorderDFS(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	postorderDFS(node.Left, vals)
	postorderDFS(node.Right, vals)
	*vals = append(*vals, node.Val)
}

// 迭代
func PostorderTraversalBaseInteration(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	prev := new(TreeNode)
	for root != nil || len(stack) > 0 {
		// 遍历左子树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		// 右子树为空或右子树已经遍历过了，记录节点的值
		if root.Right == nil || root.Right == prev {
			ans = append(ans, root.Val)
			stack = stack[:len(stack)-1]
			prev = root
			root = nil
		} else {
			//  遍历右子树
			root = root.Right
		}
	}
	return ans
}
