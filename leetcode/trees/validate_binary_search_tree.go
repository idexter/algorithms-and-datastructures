package trees

import "math"

// isValidBST checks is BST is valid.
// See: https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/625/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	left := recursiveCheck(root.Left, math.MinInt32-1, root.Val)
	right := recursiveCheck(root.Right, root.Val, math.MaxInt32+1)

	return left && right
}

func recursiveCheck(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if !(node.Val > min && node.Val < max) {
		return false
	}

	left := recursiveCheck(node.Left, min, node.Val)
	right := recursiveCheck(node.Right, node.Val, max)

	return left && right
}
