package tree

/*
144. 二叉树的前序遍历：https://leetcode.cn/problems/binary-tree-preorder-traversal/

给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
*/

func PreorderTraversal(root *TreeNode) []int {
	ans := &[]int{}
	preorderDFS(root, ans)
	return *ans
}

func preorderDFS(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	*vals = append(*vals, node.Val)
	preorderDFS(node.Left, vals)
	preorderDFS(node.Right, vals)
}

// 迭代
// 先遍历左子树节点，同时记录节点值以及将节点入栈
// 遍历到 null 时，弹出栈顶节点，将节点变更为该节点的右子树根节点，重复以上遍历
func PreorderTraversalBaseInteration(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			ans = append(ans, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return ans
}
