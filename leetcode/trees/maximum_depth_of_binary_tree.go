package trees

// maxDepth calculates max tree depth.
// See: https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/555/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return visit(root, 0)
}

// visit visit node and return current depth.
func visit(node *TreeNode, currentDepth int) int {

	if node == nil {
		return currentDepth
	}

	currentDepth++
	leftDepth := visit(node.Left, currentDepth)
	rightDepth := visit(node.Right, currentDepth)

	if leftDepth >= rightDepth {
		return leftDepth
	}
	return rightDepth
}
